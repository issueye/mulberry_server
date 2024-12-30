package route

import (
	v1 "mulberry/app/downstream/controller/v1"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	port := r.Group("port")
	{
		port.POST("", v1.CreatePort)
		port.PUT("", v1.UpdatePort)
		port.DELETE(":id", v1.DeletePort)
		port.POST("list", v1.PortList)
		port.GET(":id", v1.GetPort)
		port.PUT("reload/:port", v1.Reload)
		port.PUT("start/:port", v1.Start)
		port.PUT("stop/:port", v1.Stop)
		port.PUT("use_gz/:port", v1.ModifyUseGZ)
	}

	rule := r.Group("rule")
	{
		rule.POST("", v1.CreateRule)
		rule.PUT("", v1.UpdateRule)
		rule.PUT("status/:id", v1.UpdateRuleStatus)
		rule.DELETE(":id", v1.DeleteRule)
		rule.POST("list", v1.RuleList)
		rule.GET(":id", v1.GetRule)
	}

	page := r.Group("page")
	{
		page.POST("", v1.CreatePage)
		page.PUT("", v1.UpdatePage)
		page.PUT("status/:id", v1.UpdatePageStatus)
		page.DELETE(":id", v1.DeletePage)
		page.POST("list", v1.PageList)
		page.GET(":id", v1.GetPage)
		page.POST("save_version", v1.SaveVersionPage)
		page.GET("version/:id", v1.VersionPageList)
	}

	gzip_filter := r.Group("gzip_filter")
	{
		gzip_filter.POST("", v1.CreateGzipFilter)
		gzip_filter.PUT("", v1.UpdateGzipFilter)
		gzip_filter.DELETE(":id", v1.DeleteGzipFilter)
		gzip_filter.POST("list", v1.GzipFilterList)
		gzip_filter.GET(":id", v1.GetGzipFilter)
		gzip_filter.PUT("status/:id", v1.ModifyStatusGzipFilter)
	}

	target := r.Group("target")
	{
		target.POST("", v1.CreateTarget)
		target.PUT("", v1.UpdateTarget)
		target.DELETE(":id", v1.DeleteTarget)
		target.POST("list", v1.TargetList)
		target.GET(":id", v1.GetTarget)
		target.PUT("status/:id", v1.ModifyStatusTarget)
	}

	proxy := r.Group("proxy")
	{
		proxy.POST("traffic_messages", v1.TrafficMessages)
	}
}
