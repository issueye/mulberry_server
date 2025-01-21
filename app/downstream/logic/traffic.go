package logic

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"mulberry/app/downstream/model"
	"mulberry/app/downstream/requests"
	commonModel "mulberry/common/model"
	"mulberry/global"
	"mulberry/pkg/utils"
	"slices"
)

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

// PortForwardingTraffic 统计端口转发流量
func PortForwardingTraffic() (*model.PortForwardingStatistics, error) {
	nowDateStr := time.Now().Format("2006-01-02")
	saveKey := fmt.Sprintf("TRAFFIC:%s:", nowDateStr)

	stats := &model.PortForwardingStatistics{
		TotalRequests:  0,
		TotalInBytes:   0,
		TotalOutBytes:  0,
		PortStatistics: make(map[int]*model.PortStats),
	}

	global.STORE.ForEach(saveKey, func(key string, value []byte) error {
		statistics, err := model.TrafficStatistics{}.FromJson(value)
		if err != nil {
			global.Logger.Sugar().Errorf("解析流量数据失败: %s", err.Error())
			return nil
		}

		// 判断是否为端口转发请求
		stats.TotalRequests++
		stats.TotalInBytes += statistics.Request.InHeaderBytes + statistics.Request.InBodyBytes
		stats.TotalOutBytes += statistics.Response.OutHeaderBytes + statistics.Response.OutBodyBytes

		// 按端口统计
		// port := extractPortFromPath(statistics.Request.Path)
		port := statistics.Port
		if port > 0 {
			if _, exists := stats.PortStatistics[port]; !exists {
				stats.PortStatistics[port] = &model.PortStats{
					Requests: 0,
					InBytes:  0,
					OutBytes: 0,
				}
			}
			stats.PortStatistics[port].Requests++
			stats.PortStatistics[port].InBytes += statistics.Request.InHeaderBytes + statistics.Request.InBodyBytes
			stats.PortStatistics[port].OutBytes += statistics.Response.OutHeaderBytes + statistics.Response.OutBodyBytes
		}
		return nil
	})

	return stats, nil
}

// extractPortFromPath 从路径中提取端口号
func extractPortFromPath(path string) int {
	parts := strings.Split(path, "/")
	if len(parts) > 2 {
		port, err := strconv.Atoi(parts[2])
		if err == nil {
			return port
		}
	}
	return 0
}
