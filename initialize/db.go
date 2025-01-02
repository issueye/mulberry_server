package initialize

import (
	"mulberry/app/admin/logic"
	adminModel "mulberry/app/admin/model"
	dsLogic "mulberry/app/downstream/logic"
	downstreamModel "mulberry/app/downstream/model"
	"mulberry/global"
	"mulberry/pkg/db"
	"path/filepath"

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

	// admin
	logic.InitRoles()
	logic.InitRoleMenus()
	logic.InitUserRole()
	logic.InitAdminUser()
	logic.InitMenus()

	// downstream
	dsLogic.InitPort()
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
