package logic

import (
	"fmt"
	"mulberry/common/utils"
	"mulberry/host/app/downstream/model"
	"mulberry/host/app/downstream/requests"
	commonModel "mulberry/host/common/model"
	"mulberry/host/global"
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

	condition.Total = len(list)
	start, end := utils.SlicePage(condition.PageNum, condition.PageSize, condition.Total)
	res := commonModel.NewResPage(condition.PageNum, condition.PageSize, condition.Total, list[start:end])
	return res, nil
}
