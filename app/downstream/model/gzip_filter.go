package model

import "mulberry/common/model"

type GzipFilterInfo struct {
	model.BaseModel
	GzipFilterBase
}

type GzipFilterBase struct {
	Port         uint   `binding:"required" label:"端口号" gorm:"column:port;comment:端口信息编码;" json:"port"` // 端口信息编码
	MatchType    uint   `gorm:"column:match_type;type:int;comment:匹配模式;" json:"match_type"`             // 匹配模式
	MatchContent string `gorm:"column:match_content;size:-1;comment:匹配内容;" json:"match_content"`        // 匹配内容
	Status       bool   `gorm:"column:status;type:int;comment:状态 0 停用 1 启用;" json:"status"`             // 状态
	Remark       string `gorm:"column:remark;size:-1;comment:备注;" json:"remark"`                        // 备注
}

func (GzipFilterInfo) TableName() string { return "ds_gzip_filter_info" }
