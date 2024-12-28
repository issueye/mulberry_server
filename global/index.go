package global

import (
	"io/fs"
	"mulberry/common/logger"
	"mulberry/host/common/store"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nalgeon/redka"
	"gorm.io/gorm"
)

var (
	MsgChannel = make(chan string, 10)
	Logger     *logger.LoggerWrapper
	HttpEngine *gin.Engine
	HttpServer *http.Server
	DB         *gorm.DB
	R_DB       *redka.DB
	S_WEB      fs.FS
	STORE      store.Store
)

const (
	TOPIC_CONSOLE_LOG = "TOPIC_CONSOLE_LOG"
	ROOT_PATH         = "root"
	DEFAULT_PWD       = "123456"
	DB_Key            = "data_base:info"
)

func WriteLog(msg string) {
	MsgChannel <- msg
}

var (
	APP_NAME = "桑葚小助手"
	VERSION  = "v1.0.0.1"
)