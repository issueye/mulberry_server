package model

import "mulberry/common/model"

// 证书信息
type CertInfo struct {
	model.BaseModel
	CertBase
}

type CertBase struct {
	Name    string `binding:"required" label:"标题" gorm:"column:name;type:nvarchar(300);comment:名称;" json:"name"` // 名称
	Public  string `label:"公钥" gorm:"column:public;size:-1;comment:公有证书路径;" json:"public"`                       // 公有证书路径
	Private string `label:"私钥" gorm:"column:private;size:-1;comment:私有证书路径;" json:"private"`                     // 私有证书路径
	Remark  string `gorm:"column:remark;size:-1;comment:备注;" json:"remark"`                                      // 备注
}

func (CertInfo) TableName() string { return "ds_cert_info" }
