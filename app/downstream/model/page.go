package model

import "mulberry/host/common/model"

type PageInfo struct {
	model.BaseModel
	PageBase
}

type PageBase struct {
	Name            string            `binding:"required" label:"名称" gorm:"column:name;size:300;comment:名称;" json:"name"`                     // 名称
	Title           string            `binding:"required" label:"标题" gorm:"column:title;size:300;comment:标题;" json:"title"`                   // 标题
	Version         string            `label:"版本" gorm:"column:version;size:50;comment:版本;" json:"version"`                                   // 版本
	Port            uint              `binding:"required" label:"端口号" gorm:"column:port;comment:端口信息编码;" json:"port"`                         // 端口信息编码
	ProductCode     string            `binding:"required" label:"产品代码" gorm:"column:product_code;size:200;comment:产品代码;" json:"product_code"` // 产品代码
	UseVersionRoute int               `label:"使用版本路由" gorm:"column:use_version_route;type:int;comment:使用版本路由;" json:"use_version_route"`      // 使用版本路由
	Status          bool              `label:"状态" gorm:"column:status;type:int;comment:状态 0 停用 1 启用;" json:"status"`                          // 状态
	Remark          string            `label:"备注" gorm:"column:remark;size:-1;comment:备注;" json:"remark"`                                     // 备注
	PageVersion     []PageVersionInfo `gorm:"foreignKey:PageId" json:"page_version"`
}

func (PageInfo) TableName() string { return "ds_page_info" }

type PageVersionInfo struct {
	model.BaseModel
	PageVersionBase
}

type PageVersionBase struct {
	PageId   uint   `binding:"required" label:"页面ID" gorm:"column:page_id;type:int;comment:页面ID;" json:"page_id"` // 页面ID
	Version  string `binding:"required" label:"版本" gorm:"column:version;size:50;comment:版本;" json:"version"`      // 版本
	PagePath string `label:"页面路径" gorm:"column:page_path;size:2000;comment:页面路径;" json:"page_path"`               // 页面路径
	Mark     string `label:"备注" gorm:"column:mark;size:2000;comment:备注;" json:"mark"`                             // 备注
}

func (PageVersionInfo) TableName() string { return "ds_page_version_info" }
