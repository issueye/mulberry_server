package pages

import "mulberry/host/global"

func WriteLog(log string) {
	if global.Logger != nil {
		global.Logger.Info(log)
	}

	global.MsgChannel <- log
}
