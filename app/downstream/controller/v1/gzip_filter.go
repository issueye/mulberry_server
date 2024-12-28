package v1

import (
	"mulberry/host/app/downstream/logic"
	"mulberry/host/app/downstream/requests"
	"mulberry/host/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateGzipFilter doc
//
//	@tags			GZIP过滤管理
//	@Summary		添加GZIP过滤信息
//	@Description	添加GZIP过滤信息
//	@Produce		json
//	@Param			body	body		requests.CreateGzipFilter	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/gzip_filter [post]
//	@Security		ApiKeyAuth
func CreateGzipFilter(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateGzipFilter()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateGzipFilter(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateGzipFilter doc
//
//	@tags			GZIP过滤管理
//	@Summary		修改GZIP过滤信息
//	@Description	修改GZIP过滤信息
//	@Produce		json
//	@Param			body	body		requests.UpdateGzipFilter	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/gzip_filter [put]
//	@Security		ApiKeyAuth
func UpdateGzipFilter(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateGzipFilter()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateGzipFilter(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteGzipFilter doc
//
//	@tags			GZIP过滤管理
//	@Summary		删除GZIP过滤信息
//	@Description	删除GZIP过滤信息
//	@Produce		json
//	@Param			id		path	int	true	"GZIP过滤id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/gzip_filter/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteGzipFilter(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.DeleteGzipFilter(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// GzipFilterList doc
//
//	@tags			GZIP过滤管理
//	@Summary		GZIP过滤列表
//	@Description	GZIP过滤列表
//	@Produce		json
//	@Param			body	body		requests.QueryGzipFilter	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/gzip_filter/list [post]
//	@Security		ApiKeyAuth
func GzipFilterList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryGzipFilter()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.GzipFilterList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetGzipFilter doc
//
//	@tags			GZIP过滤管理
//	@Summary		GZIP过滤详情
//	@Description	GZIP过滤详情
//	@Produce		json
//	@Param			id		path	int	true	"GZIP过滤id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/gzip_filter/{id} [get]
//	@Security		ApiKeyAuth
func GetGzipFilter(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	info, err := logic.GetGzipFilter(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}

// ModifyStatusGzipFilter doc
//
//	@tags			目标服务管理
//	@Summary		修改目标服务状态
//	@Description	修改目标服务状态
//	@Produce		json
//	@Param			id		path	int	true	"目标服务id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/gzip_filter/status/{id} [put]
//	@Security		ApiKeyAuth
func ModifyStatusGzipFilter(c *gin.Context) {
	ctl := controller.New(c)

	id := c.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.ModifyStatusGzipFilter(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
