package v1

import (
	"mulberry/internal/app/downstream/logic"
	"mulberry/internal/app/downstream/requests"
	"mulberry/internal/common/controller"

	"github.com/gin-gonic/gin"
)

// TrafficMessages doc
//
//	@tags			查询
//	@Summary		查询代理流量信息
//	@Description	查询代理流量信息
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/traffic/status [post]
//	@Security		ApiKeyAuth
func TrafficMessages(c *gin.Context) {
	ctl := controller.New(c)

	// 解析参数
	req := requests.NewQueryTraffic()
	err := c.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.TrafficMessages(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// PortStatistics doc
//
//	@tags			查询
//	@Summary		查询端口转发流量统计（按端口分组）
//	@Description	查询端口转发流量统计信息，按端口分组返回
//	@Produce		json
//	@Success		200	{object}	controller.Response{data=model.PortForwardingStatistics}	"成功返回带端口统计的数据结构"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/traffic/port_statistics [get]
//	@Security		ApiKeyAuth
func PortStatistics(c *gin.Context) {
	ctl := controller.New(c)

	stats, err := logic.PortForwardingTraffic()
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(stats)
}
