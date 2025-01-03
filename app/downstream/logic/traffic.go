package logic

import (
	"fmt"
	"mulberry/app/downstream/model"
	"mulberry/app/downstream/requests"
	commonModel "mulberry/common/model"
	"mulberry/global"
	"mulberry/pkg/utils"
	"slices"
	"time"
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
