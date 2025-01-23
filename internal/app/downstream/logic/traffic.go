package logic

import (
	"fmt"
	"sync"
	"time"

	"mulberry/internal/app/downstream/model"
	"mulberry/internal/app/downstream/requests"
	commonModel "mulberry/internal/common/model"
	"mulberry/internal/global"
	"mulberry/pkg/utils"
	"slices"
)

var (
	portTrafficStats   = make(map[int]*model.Statistics)
	portTrafficStatsMu sync.RWMutex
)

func init() {
	go startTrafficStatsTask()
}

func startTrafficStatsTask() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			updatePortTrafficStats()
		}
	}
}

func updatePortTrafficStats() {
	now := time.Now()
	for port, stats := range portTrafficStats {
		fmt.Printf("Port: %d, TotalRequests: %d, InBytes: %d, OutBytes: %d\n",
			port, stats.TotalRequests, stats.TotalInBytes, stats.TotalOutBytes)
	}

	// 清理超过五小时的数据
	for port, stats := range portTrafficStats {
		if now.Sub(stats.LastUpdated) > 5*time.Hour {
			delete(portTrafficStats, port)
		}
	}
}

func TrafficMessages(condition *commonModel.PageQuery[*requests.QueryTraffic]) (*commonModel.ResPage[model.TrafficStatistics], error) {
	list := make([]*model.TrafficStatistics, 0)
	nowDateStr := time.Now().Format("2006-01-02")
	saveKey := fmt.Sprintf("TRAFFIC:%s:", nowDateStr)
	global.STORE.ForEach(saveKey, func(key string, value []byte) error {
		statistics, err := model.TrafficStatistics{}.FromJson(value)
		if err != nil {
			global.Logger.Sugar().Errorf("解析流量数据失败: %s", err.Error())
			return nil
		}

		list = append(list, statistics)
		return nil
	})

	// 排序
	slices.SortFunc(list, func(a *model.TrafficStatistics, b *model.TrafficStatistics) int {
		// 按照时间倒序
		return int(b.Request.Time.Unix() - a.Request.Time.Unix())
	})

	condition.Total = len(list)
	start, end := utils.SlicePage(condition.PageNum, condition.PageSize, condition.Total)
	res := commonModel.NewResPage(condition.PageNum, condition.PageSize, condition.Total, list[start:end])
	return res, nil
}

// PortTraffic 统计端口转发流量
func Traffic() (*model.Statistics, error) {
	nowDateStr := time.Now().Format("2006-01-02")
	saveKey := fmt.Sprintf("TRAFFIC:%s:", nowDateStr)

	stats := &model.Statistics{
		TotalRequests: 0,
		TotalOutBytes: 0,
		TotalInBytes:  0,
	}

	global.STORE.ForEach(saveKey, func(key string, value []byte) error {
		statistics, err := model.TrafficStatistics{}.FromJson(value)
		if err != nil {
			global.Logger.Sugar().Errorf("解析流量数据失败: %s", err.Error())
			return nil
		}

		// 判断是否为端口转发请求
		stats.TotalRequests++
		stats.TotalInBytes += int(statistics.Request.InHeaderBytes + statistics.Request.InBodyBytes)
		stats.TotalOutBytes += int(statistics.Response.OutHeaderBytes + statistics.Response.OutBodyBytes)

		return nil
	})

	return stats, nil
}

func GetPortTrafficStats() map[int]*model.Statistics {
	// 数据至少保证5小时

	return portTrafficStats
}

// HourlyTraffic 统计每小时的流量
func HourlyTraffic() ([]*model.HourlyTrafficStatistics, error) {
	nowDateStr := time.Now().Format("2006-01-02")
	saveKey := fmt.Sprintf("TRAFFIC:%s:", nowDateStr)
	hourlyStats := make(map[string]*model.HourlyTrafficStatistics)

	global.STORE.ForEach(saveKey, func(key string, value []byte) error {
		statistics, err := model.TrafficStatistics{}.FromJson(value)
		if err != nil {
			global.Logger.Sugar().Errorf("解析流量数据失败: %s", err.Error())
			return nil
		}

		hour := statistics.Request.Time.Hour()
		minute := statistics.Request.Time.Minute() / 30 * 30 // Round to nearest half-hour
		key = fmt.Sprintf("%02d:%02d", hour, minute)
		if _, exists := hourlyStats[key]; !exists {
			hourlyStats[key] = &model.HourlyTrafficStatistics{
				Hour:          hour,
				Minute:        minute,
				TotalRequests: 0,
				TotalInBytes:  0,
				TotalOutBytes: 0,
			}
		}

		hourlyStats[key].TotalRequests++
		hourlyStats[key].TotalInBytes += int(statistics.Request.InHeaderBytes + statistics.Request.InBodyBytes)
		hourlyStats[key].TotalOutBytes += int(statistics.Response.OutHeaderBytes + statistics.Response.OutBodyBytes)

		return nil
	})

	// Ensure data for the last 10 half-hour intervals
	currentTime := time.Now()
	result := make([]*model.HourlyTrafficStatistics, 0, 10)
	for i := 0; i < 10; i++ {
		intervalTime := currentTime.Add(-time.Duration(i*30) * time.Minute)
		hour := intervalTime.Hour()
		minute := intervalTime.Minute() / 30 * 30 // Round to nearest half-hour
		key := fmt.Sprintf("%02d:%02d", hour, minute)
		if _, exists := hourlyStats[key]; !exists {
			hourlyStats[key] = &model.HourlyTrafficStatistics{
				Hour:          hour,
				Minute:        minute,
				TotalRequests: 0,
				TotalInBytes:  0,
				TotalOutBytes: 0,
			}
		}

		result = append(result, hourlyStats[key])
	}

	// 按小时排序
	slices.SortFunc(result, func(a, b *model.HourlyTrafficStatistics) int {
		if a.Hour == b.Hour {
			return a.Minute - b.Minute
		}
		return a.Hour - b.Hour
	})

	return result, nil
}
