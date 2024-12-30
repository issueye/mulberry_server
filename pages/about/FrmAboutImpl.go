package about

import (
	"mulberry/global"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// ::private::
type TFrmAboutFields struct {
}

func (f *TFrmAbout) OnFormCreate(sender vcl.IObject) {
	f.TForm.SetPosition(types.PoScreenCenter)
	f.SetCaption(global.APP_NAME + " " + global.VERSION)
	f.Lbl_server_name.SetCaption("服务名称：" + global.APP_NAME)
	f.Lbl_server_version.SetCaption("服务版本：" + global.VERSION)
}
