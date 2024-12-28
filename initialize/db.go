package initialize

import (
	"mulberry/common/db"
	"mulberry/host/app/admin/logic"
	adminModel "mulberry/host/app/admin/model"
	downstreamModel "mulberry/host/app/downstream/model"
	"mulberry/host/global"
	"path/filepath"

	"github.com/nalgeon/redka"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"gorm.io/gorm"
)

func InitDB() {
	path := filepath.Join(global.ROOT_PATH, "data", "data.db")
	global.DB = db.InitSqlite(path, global.Logger.Sugar())

	InitDATA(global.DB)
}

func InitDATA(db *gorm.DB) {
	db.AutoMigrate(&adminModel.User{})
	db.AutoMigrate(&adminModel.Role{})
	db.AutoMigrate(&adminModel.UserRole{})
	db.AutoMigrate(&adminModel.RoleMenu{})
	db.AutoMigrate(&adminModel.Menu{})

	db.AutoMigrate(&downstreamModel.PortInfo{})
	db.AutoMigrate(&downstreamModel.PageVersionInfo{})
	db.AutoMigrate(&downstreamModel.PageInfo{})
	db.AutoMigrate(&downstreamModel.GzipFilterInfo{})
	db.AutoMigrate(&downstreamModel.RuleInfo{})
	db.AutoMigrate(&downstreamModel.CertInfo{})
	db.AutoMigrate(&downstreamModel.TargetInfo{})

	logic.InitRoles()
	logic.InitRoleMenus()
	logic.InitUserRole()
	logic.InitAdminUser()
	logic.InitMenus()
}

func FreeDB() {
	sqldb, err := global.DB.DB()
	if err != nil {
		global.Logger.Sugar().Errorf("close db error: %v", err)
	}

	if err = sqldb.Close(); err != nil {
		global.Logger.Sugar().Errorf("close db error: %v", err)
	}
}

func InitRedkaDB() {
	path := filepath.Join(global.ROOT_PATH, "data", "rdb.db")

	rdb, err := redka.Open(path, nil)
	if err != nil {
		panic(err)
	}

	global.R_DB = rdb
}
