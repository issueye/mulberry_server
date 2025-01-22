package pages

import "mulberry/internal/global"

func WriteLog(log string) {
	if global.Logger != nil {
		global.Logger.Info(log)
	}

	global.MsgChannel <- log
}
