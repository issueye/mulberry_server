package v1

import (
	"mulberry/host/app/downstream/logic"
	"mulberry/host/app/downstream/requests"
	"mulberry/host/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTarget doc
//
//	@tags			目标服务管理
//	@Summary		添加目标服务信息
//	@Description	添加目标服务信息
//	@Produce		json
//	@Param			body	body		requests.CreateTarget	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/target [post]
//	@Security		ApiKeyAuth
func CreateTarget(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateTarget()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateTarget(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateTarget doc
//
//	@tags			目标服务管理
//	@Summary		修改目标服务信息
//	@Description	修改目标服务信息
//	@Produce		json
//	@Param			body	body		requests.UpdateTarget	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/target [put]
//	@Security		ApiKeyAuth
func UpdateTarget(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateTarget()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateTarget(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteTarget doc
//
//	@tags			目标服务管理
//	@Summary		删除目标服务信息
//	@Description	删除目标服务信息
//	@Produce		json
//	@Param			id		path	int	true	"目标服务id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/target/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteTarget(c *gin.Context) {
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

	err = logic.DeleteTarget(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// TargetList doc
//
//	@tags			目标服务管理
//	@Summary		目标服务列表
//	@Description	目标服务列表
//	@Produce		json
//	@Param			body	body		requests.QueryTarget	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/target/list [post]
//	@Security		ApiKeyAuth
func TargetList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryTarget()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.TargetList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetTarget doc
//
//	@tags			目标服务管理
//	@Summary		目标服务详情
//	@Description	目标服务详情
//	@Produce		json
//	@Param			id		path	int	true	"目标服务id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/target/{id} [get]
//	@Security		ApiKeyAuth
func GetTarget(c *gin.Context) {
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

	info, err := logic.GetTarget(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}

// ModifyStatusTarget doc
//
//	@tags			目标服务管理
//	@Summary		修改目标服务状态
//	@Description	修改目标服务状态
//	@Produce		json
//	@Param			id		path	int	true	"目标服务id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/target/status/{id} [put]
//	@Security		ApiKeyAuth
func ModifyStatusTarget(c *gin.Context) {
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

	err = logic.ModifyStatusTarget(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
