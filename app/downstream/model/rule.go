package model

import "mulberry/common/model"

type RuleInfo struct {
	model.BaseModel
	RuleBase
}

type RuleBase struct {
	Name        string     `gorm:"column:name;size:300;comment:匹配路由名称;" json:"name"`                       // 匹配路由名称
	TargetId    uint       `gorm:"column:target_id;int;comment:目标路由;" json:"target_id"`                    // 目标地址编码
	TargetRoute string     `gorm:"column:target_route;size:300;comment:目标路由;" json:"target_route"`         // 目标路由
	Port        uint       `gorm:"column:port;type:int;comment:端口号;" json:"port"`                          // 端口号                                      // 端口信息编码
	Method      string     `gorm:"column:method;size:100;comment:请求方法;" json:"method"`                     // 请求方法
	Status      bool       `gorm:"column:status;type:int;comment:状态 0 停用 1 启用;" json:"status"`             // 状态
	MatchType   uint       `gorm:"column:match_type;type:int;comment:0 GIN 匹配 1 MUX 匹配" json:"match_type"` // 0 GIN 匹配 1 MUX 匹配
	Remark      string     `gorm:"column:remark;size:-1;comment:备注;" json:"remark"`                        // 备注
	Target      TargetInfo `gorm:"foreignKey:TargetId" json:"target"`                                      // 目标地址
	Order       uint       `gorm:"column:order;type:int;comment:排序;" json:"order"`                         // 排序
	IsWs        bool       `gorm:"column:is_ws;type:int;comment:是否是ws;" json:"is_ws"`                      // 是否是ws
}

func (RuleInfo) TableName() string { return "dz_rule_info" }
