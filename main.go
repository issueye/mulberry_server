package main

import (
	"embed"
	"io/fs"
	"mulberry/global"
	"mulberry/initialize"
	"mulberry/pages/home"
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
	initialize.InitRuntime()
	vcl.RunApp(&home.FrmHome)
}
