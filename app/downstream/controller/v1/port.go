package v1

import (
	"mulberry/host/app/downstream/logic"
	"mulberry/host/app/downstream/requests"
	"mulberry/host/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePort doc
//
//	@tags			端口号管理
//	@Summary		添加端口号信息
//	@Description	添加端口号信息
//	@Produce		json
//	@Param			body	body		requests.CreatePort	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/port [post]
//	@Security		ApiKeyAuth
func CreatePort(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewCreatePort()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreatePort(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdatePort doc
//
//	@tags			端口号管理
//	@Summary		修改端口号信息
//	@Description	修改端口号信息
//	@Produce		json
//	@Param			body	body		requests.UpdatePort	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/port [put]
//	@Security		ApiKeyAuth
func UpdatePort(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewUpdatePort()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdatePort(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// DeletePort doc
//
//	@tags			端口号管理
//	@Summary		删除端口号信息
//	@Description	删除端口号信息
//	@Produce		json
//	@Param			id		path	int	true	"端口号id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/port/{id} [delete]
//	@Security		ApiKeyAuth
func DeletePort(c *gin.Context) {
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

	err = logic.DeletePort(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// PortList doc
//
//	@tags			端口号管理
//	@Summary		端口号列表
//	@Description	端口号列表
//	@Produce		json
//	@Param			body	body		requests.QueryPort	true	"body"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/port/list [post]
//	@Security		ApiKeyAuth
func PortList(c *gin.Context) {
	ctl := controller.New(c)

	req := requests.NewQueryPort()

	err := ctl.Bind(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	list, err := logic.PortList(req)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(list)
}

// GetPort doc
//
//	@tags			端口号管理
//	@Summary		端口号详情
//	@Description	端口号详情
//	@Produce		json
//	@Param			id		path	int	true	"端口号id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/port/{id} [get]
//	@Security		ApiKeyAuth
func GetPort(c *gin.Context) {
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

	info, err := logic.GetPort(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(info)
}

// Reload doc
//
//	@tags			端口号管理
//	@Summary		重载
//	@Description	重载
//	@Produce		json
//	@Param			id		path	int	true	"端口号id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/port/reload/{id} [put]
//	@Security		ApiKeyAuth
func Reload(c *gin.Context) {
	ctl := controller.New(c)

	portStr := c.Param("port")
	if portStr == "" {
		ctl.Fail("port不能为空")
		return
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.Reload(uint(port))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// Start doc
//
//	@tags			端口号管理
//	@Summary		开启
//	@Description	开启
//	@Produce		json
//	@Param			id		path	int	true	"端口号id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/port/reload/{id} [put]
//	@Security		ApiKeyAuth
func Start(c *gin.Context) {
	ctl := controller.New(c)

	portStr := c.Param("port")
	if portStr == "" {
		ctl.Fail("port不能为空")
		return
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.Start(uint(port))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// Stop doc
//
//	@tags			端口号管理
//	@Summary		关闭
//	@Description	关闭
//	@Produce		json
//	@Param			id		path	int	true	"端口号id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/port/reload/{id} [put]
//	@Security		ApiKeyAuth
func Stop(c *gin.Context) {
	ctl := controller.New(c)

	portStr := c.Param("port")
	if portStr == "" {
		ctl.Fail("port不能为空")
		return
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.Stop(uint(port))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// ModifyUseGZ doc
//
//	@tags			端口号管理
//	@Summary		使用GZIP
//	@Description	使用GZIP
//	@Produce		json
//	@Param			id		path	int	true	"端口号id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/port/use_gz/{id} [put]
//	@Security		ApiKeyAuth
func ModifyUseGZ(c *gin.Context) {
	ctl := controller.New(c)

	portStr := c.Param("port")
	if portStr == "" {
		ctl.Fail("port不能为空")
		return
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.ModifyUseGZ(uint(port))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
