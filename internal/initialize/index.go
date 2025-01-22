package initialize

import (
	"context"
	"mulberry/internal/global"
)

func RunServer(ctx context.Context) {
	InitRuntime()
	InitConfig()
	InitLogger()
	InitDB()
	InitEngine(ctx)
	InitStore()
	InitHttpServer(ctx)
}

func StopServer() {
	// http 服务
	FreeHttpServer()
	// 释放数据库连接
	FreeDB()

	// 处理日志
	if global.Logger != nil {
		global.WriteLog("日志关闭")
		global.Logger.Close()
		global.WriteLog("日志关闭成功")
	}
}
