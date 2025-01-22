package home

import (
	"context"
	"fmt"
	"mulberry/internal/common/config"
	"mulberry/internal/global"
	"mulberry/internal/initialize"
	"mulberry/internal/pages"
	"os"
	"os/exec"
	"runtime"
)

// 初始化数据
func (f *TFrmHome) InitData() {
	f.ShowRunInfo()
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		return fmt.Errorf("unsupported operating system")
	}
	return cmd.Start()
}

func (f *TFrmHome) StartServer() {
	f.serverRunCtx, f.serverRunCancel = context.WithCancel(context.Background())
	initialize.RunServer(f.ctx)
	pages.WriteLog("启动服务")
}

func (f *TFrmHome) StopServer() {
	defer f.serverRunCancel()

	initialize.StopServer()
	pages.WriteLog("停止服务")
	// 如果服务已经停止，就强制 GC
	runtime.GC()
}

func (f *TFrmHome) ShowRunInfo() {
	f.Lbl_name.SetCaption("名称：" + global.APP_NAME)
	port := config.GetParam(config.SERVER, "http-port", 6678).Int()
	f.Lbl_port.SetCaption(fmt.Sprintf("端口：%d", port))
	f.Lbl_version.SetCaption(fmt.Sprintf("版本：%s", global.VERSION))

	item_pid := f.StatusBar.Panels().Items(0)
	item_pid.SetText(fmt.Sprintf("PID：%d", os.Getpid()))
}
