package home

import (
	"context"
	"fmt"
	"mulberry/host/global"
	"mulberry/host/pages/about"
	"time"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// ::private::
type TFrmHomeFields struct {
	IsServerRun  bool
	ShowLogCount int32
	IsTrueClose  bool
	ActiveItem   int64
	ctx          context.Context
	cancel       context.CancelFunc
	StartTime    time.Time
	runDate      int

	serverRunCtx    context.Context
	serverRunCancel context.CancelFunc
}

func (f *TFrmHome) OnFormCreate(sender vcl.IObject) {
	f.IsServerRun = false
	f.Btn_server_run.SetImageIndex(0)

	f.SetCaption(global.APP_NAME)
	f.TForm.SetPosition(types.PoScreenCenter)
	f.ShowHint()
	f.Tray_icon.SetVisible(true)
	f.Tray_icon.SetHint(global.APP_NAME)
	f.StatusBar.Panels().Items(0).SetWidth(180)
	f.ctx, f.cancel = context.WithCancel(context.Background())

	f.InitData()
	f.SetEvents()

	f.Monitor()
	f.EventMonitor(f.ctx)
}

func (f *TFrmHome) SetEvents() {
	f.TForm.SetOnClose(f.OnFormClose)
	f.TForm.SetOnDestroy(f.OnFormDestroy)

	f.Tray_icon.SetOnClick(f.OnTrayIconClick)

	f.Timer.SetOnTimer(f.OnTimer)
	f.Meu_item_about.SetOnClick(f.OnAboutClick)

	f.PM_close.SetOnClick(f.OnAppCloseClick)
	f.PM_show.SetOnClick(f.OnAppShowClick)
	f.Btn_server_run.SetOnClick(f.OnBtn_server_runClick)
	f.Btn_clear_log.SetOnClick(f.OnBtn_clear_logClick)

	go func() {
		timer := time.NewTimer(time.Second * 2)
		<-timer.C
		f.Btn_server_run.SetEnabled(false)
		f.OnBtn_server_runClick(nil)
	}()
}

func (f *TFrmHome) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	f.Hide()
	if !f.IsTrueClose {
		*action = types.CaHide
	}
}

func (f *TFrmHome) OnBtn_clear_logClick(sender vcl.IObject) {
	f.Mmo_run_log.Clear()
}

func (f *TFrmHome) OnBtn_server_runClick(sender vcl.IObject) {
	if f.IsServerRun {
		f.Btn_server_run.SetImageIndex(0)
		f.IsServerRun = false
		f.StopServer()
	} else {
		f.Btn_server_run.SetImageIndex(1)
		f.IsServerRun = true
		f.ShowRunInfo()
		f.StartServer()
	}
}

func (f *TFrmHome) OnFormDestroy(sender vcl.IObject) {
	f.cancel()
}

func (f *TFrmHome) OnAppCloseClick(sender vcl.IObject) {
	f.IsTrueClose = true
	f.Close()
}

func (f *TFrmHome) OnAppShowClick(sender vcl.IObject) {
	f.Show()
}

func (f *TFrmHome) OnAboutClick(sender vcl.IObject) {
	about := about.NewFrmAbout(f)
	about.SetPosition(types.PoOwnerFormCenter)
	about.ShowModal()
}

func (f *TFrmHome) OnTrayIconClick(sender vcl.IObject) {
	f.Show()
}

func (f *TFrmHome) OnTimer(sender vcl.IObject) {
	item := f.StatusBar.Panels().Items(0)
	now := time.Now().Format("当前时间：2006/01/02 15:04:05")
	item.SetText(now)

	item_run_date := f.StatusBar.Panels().Items(1)

	// 计算已经运行天数
	if f.IsServerRun {
		if f.runDate != GetRunDate(f.StartTime) {
			f.runDate = GetRunDate(f.StartTime)
			item_run_date.SetText(fmt.Sprintf("运行天数：%d", f.runDate))
		}
	}
}

func GetRunDate(start time.Time) int {
	diff := time.Since(start)
	return int(diff.Hours() / 24)
}
