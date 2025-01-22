package main

import (
	"context"
	"embed"
	"io/fs"
	"mulberry/internal/global"
	"mulberry/internal/initialize"
	"mulberry/internal/pages/home"
	_ "mulberry/winappres"

	"github.com/ying32/govcl/vcl"
)

//	@title			桑葚小助手v1.0
//	@version		V1.1
//	@description	桑葚小助手v1.0

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//go:embed web/*
var webDir embed.FS

func main() {
	staticFp, _ := fs.Sub(webDir, "web")
	global.S_WEB = staticFp

	ctx := context.Background()
	initialize.RunServer(ctx)
	vcl.RunApp(&home.FrmHome)
	initialize.StopServer()
}
