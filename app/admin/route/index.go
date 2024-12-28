package route

import (
	v1 "mulberry/host/app/admin/controller/v1"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	auth := r.Group("auth")
	{
		auth.POST("login", v1.Login)
		auth.POST("logout", v1.Logout)
		auth.GET("refresh", v1.RefreshToken)
	}

	admin := r.Group("admin")
	{
		admin.GET("userInfo", v1.GetUserInfo)
		admin.POST("upload", v1.Upload)
		admin.POST("updateuserinfo", v1.UpdateUserinfo)
		admin.POST("updatepassword", v1.UpdatePassword)
		admin.GET("homeCount", v1.GetHomeCount)
	}

	user := r.Group("user")
	{
		user.POST("list", v1.GetUsers)
		user.PUT("update", v1.UpdateUser)
		user.DELETE("delete/:id", v1.DeleteUser)
		user.POST("add", v1.CreateUser)
	}

	role := r.Group("role")
	{
		role.POST("list", v1.GetRoles)
		role.PUT("update", v1.UpdateRole)
		role.DELETE("delete/:id", v1.DeleteRole)
		role.POST("add", v1.CreateRole)
	}

	menu := r.Group("menu")
	{
		menu.POST("list", v1.GetMenus)
		menu.GET("roleMenus/:code", v1.GetRoleMenus)
		menu.GET("catalog", v1.GetCatalog)
		menu.POST("saveRoleMenus/:code", v1.SaveRoleMenus)
		menu.PUT("update", v1.UpdateMenu)
		menu.DELETE("delete/:id", v1.DeleteMenu)
		menu.POST("add", v1.CreateMenu)
	}

	settings := r.Group("settings")
	{
		settings.GET("system", v1.GetSystemSettings)
		settings.PUT("system", v1.SetSystemSettings)
		settings.GET("logger", v1.GetLoggerSettings)
		settings.PUT("logger", v1.SetLoggerSettings)
	}
}
