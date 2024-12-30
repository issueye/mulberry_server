package v1

import (
	"mulberry/app/admin/logic"
	"mulberry/app/admin/requests"
	"mulberry/common/controller"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMenus doc
//
//	@tags			菜单
//	@Summary		获取菜单列表
//	@Description	获取菜单列表
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/list [get]
//	@Security		ApiKeyAuth
func GetMenus(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewQueryMenu()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	roles, err := logic.ListMenu(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(roles)
}

// GetCatalog doc
//
//	@tags			菜单
//	@Summary		获取一级菜单列表
//	@Description	获取一级菜单列表
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/catalog [get]
//	@Security		ApiKeyAuth
func GetCatalog(c *gin.Context) {
	ctl := controller.New(c)
	menus, err := logic.GetCatalog()
	if err != nil {
		ctl.FailWithError(err)
		return
	}
	ctl.SuccessData(menus)
}

// GetRoleMenus doc
//
//	@tags			菜单
//	@Summary		获取菜单列表
//	@Description	获取菜单列表
//	@Produce		json
//	@Param			code		path	string	true	"角色编码"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/roleMenus/{code} [get]
//	@Security		ApiKeyAuth
func GetRoleMenus(c *gin.Context) {
	ctl := controller.New(c)
	code := c.Param("code")
	if code == "" {
		ctl.Fail("code不能为空")
		return
	}

	menus, err := logic.GetRoleMenus(code)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.SuccessData(menus)
}

// SaveRoleMenus doc
//
//	@tags			菜单
//	@Summary		保存角色菜单
//	@Description	保存角色菜单
//	@Produce		json
//	@Param			code		path	string	true	"角色编码"
//	@Param			menus		body	[]string	true	"菜单id列表"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/saveRoleMenus/{code} [post]
//	@Security		ApiKeyAuth
func SaveRoleMenus(c *gin.Context) {
	ctl := controller.New(c)
	code := c.Param("code")
	if code == "" {
		ctl.Fail("code不能为空")
		return
	}

	menus := make([]string, 0)
	err := c.BindJSON(&menus)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.SaveRoleMenus(code, menus)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// CreateMenu doc
//
//	@tags			菜单
//	@Summary		创建菜单
//	@Description	创建菜单
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/create [post]
//	@Security		ApiKeyAuth
func CreateMenu(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewCreateMenu()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.CreateMenu(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}

// UpdateMenu doc
//
//	@tags			菜单
//	@Summary		修改菜单
//	@Description	修改菜单
//	@Produce		json
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/update [put]
//	@Security		ApiKeyAuth
func UpdateMenu(c *gin.Context) {
	ctl := controller.New(c)

	condition := requests.NewUpdateMenu()
	err := ctl.Bind(condition)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.UpdateMenu(condition)
	if err != nil {
		ctl.FailWithError(err)
	}

	ctl.Success()
}

// DeleteMenu doc
//
//	@tags			菜单
//	@Summary		删除菜单
//	@Description	删除菜单
//	@Produce		json
//	@Param			id		path	int	true	"菜单id"
//	@Success		200	{object}	controller.Response	"code: 200 成功"
//	@Failure		500	{object}	controller.Response						"错误返回内容"
//	@Router			/api/v1/menu/delete [delete]
//	@Security		ApiKeyAuth
func DeleteMenu(c *gin.Context) {
	ctl := controller.New(c)

	id := ctl.Param("id")
	if id == "" {
		ctl.Fail("id不能为空")
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	err = logic.DeleteMenu(uint(i))
	if err != nil {
		ctl.FailWithError(err)
		return
	}

	ctl.Success()
}
