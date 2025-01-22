package v1

import (
	"mulberry/internal/app/downstream/logic"
	"mulberry/internal/app/downstream/requests"
	"mulberry/internal/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateRule doc
//
//	@tags			规则管理
//	@Summary		添加规则信息
//	@Description	添加规则信息
//	@Produce		json
//	@Param			body	body		requests.CreateRule	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/rule [post]
//	@Security		ApiKeyAuth
func CreateRule(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreateRule()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateRule(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateRule doc
//
//	@tags			规则管理
//	@Summary		修改规则信息
//	@Description	修改规则信息
//	@Produce		json
//	@Param			body	body		requests.UpdateRule	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/rule [put]
//	@Security		ApiKeyAuth
func UpdateRule(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdateRule()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateRule(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateRuleStatus doc
//
//	@tags			规则管理
//	@Summary		修改规则信息
//	@Description	修改规则信息
//	@Produce		json
//	@Param			id		path	int	true	"规则id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/rule/status/{id} [put]
//	@Security		ApiKeyAuth
func UpdateRuleStatus(c *gin.Context) {
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

	err = logic.UpdateRuleStatus(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeleteRule doc
//
//	@tags			规则管理
//	@Summary		删除规则信息
//	@Description	删除规则信息
//	@Produce		json
//	@Param			id		path	int	true	"规则id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/rule/{id} [delete]
//	@Security		ApiKeyAuth
func DeleteRule(c *gin.Context) {
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

	err = logic.DeleteRule(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// RuleList doc
//
//	@tags			规则管理
//	@Summary		规则列表
//	@Description	规则列表
//	@Produce		json
//	@Param			body	body		requests.QueryRule	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/rule/list [post]
//	@Security		ApiKeyAuth
func RuleList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryRule()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.RuleList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetRule doc
//
//	@tags			规则管理
//	@Summary		规则详情
//	@Description	规则详情
//	@Produce		json
//	@Param			id		path	int	true	"规则id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/rule/{id} [get]
//	@Security		ApiKeyAuth
func GetRule(c *gin.Context) {
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

	info, err := logic.GetRule(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}
