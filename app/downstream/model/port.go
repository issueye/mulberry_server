package model

import "mulberry/host/common/model"

type PortInfo struct {
	model.BaseModel
	PortBase
}

type PortBase struct {
	Port            uint   `gorm:"column:port;type:int;comment:端口号;" json:"port"`                                // 端口号
	Status          bool   `gorm:"column:status;type:int;comment:状态 0 停用 1 启用;" json:"status"`                   // 状态
	IsTLS           bool   `gorm:"column:is_tls;type:int;comment:是否https;" json:"is_tls"`                        // 是否证书加密
	CerId           string `gorm:"column:cert_id;size:100;comment:编码;" json:"cert_id"`                           // 证书编码
	UseGzip         bool   `gorm:"column:use_gzip;type:int;comment:使用GZIP 0 停用 1 启用;" json:"use_gzip"`           // 使用GZIP 0 停用 1 启用
	PageCount       int    `gorm:"column:page_count;type:int;comment:页面总数;" json:"page_count"`                   // 页面总数
	RuleCount       int    `gorm:"column:rule_count;type:int;comment:规则总数;" json:"rule_count"`                   // 规则总数
	GzipFilterCount int    `gorm:"column:gzip_filter_count;type:int;comment:GZIP过滤总数;" json:"gzip_filter_count"` // GZIP过滤总数
	Remark          string `gorm:"column:remark;size:-1;comment:备注;" json:"remark"`                              // 备注
}

func (PortInfo) TableName() string { return "ds_port_info" }
