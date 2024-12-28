package v1

import (
	"mulberry/host/app/downstream/logic"
	"mulberry/host/app/downstream/requests"
	"mulberry/host/common/controller"

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
