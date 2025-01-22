package route

import (
	adminRoute "mulberry/internal/app/admin/route"
	downstreamRoute "mulberry/internal/app/downstream/route"
	"mulberry/internal/common/controller"
	"mulberry/internal/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctl := controller.New(ctx)
			ctl.SuccessData(map[string]any{"msg": "pong"})
		})

		// 注册管理路由
		adminRoute.Register(v1)
		// 转发服务理由
		downstreamRoute.Register(v1)
	}

	r.NoRoute(func(ctx *gin.Context) {
		global.Logger.Logger.Error("404", zap.String("path", ctx.Request.URL.Path), zap.String("method", ctx.Request.Method))
		ctl := controller.New(ctx)
		ctl.FailWithCode(http.StatusNotFound, "not found")
	})
}
