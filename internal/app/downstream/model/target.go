package model

import "mulberry/internal/common/model"

type TargetInfo struct {
	model.BaseModel
	TargetBase
}

type TargetBase struct {
	Title  string `gorm:"column:title;size:300;comment:目标名称;" json:"title"`           // 目标名称
	Name   string `gorm:"column:name;size:300;comment:目标地址;" json:"name"`             // 目标地址
	Status bool   `gorm:"column:status;type:int;comment:状态 0 停用 1 启用;" json:"status"` // 状态
	Remark string `gorm:"column:remark;size:-1;comment:备注;" json:"remark"`            // 备注
}

func (TargetInfo) TableName() string { return "dz_target_info" }
