package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"mulberry/internal/common/model"
)

type RuleInfo struct {
	model.BaseModel
	RuleBase
}

type KV struct {
	K string `json:"k"`
	V string `json:"v"`
}

type Headers []KV

// Value 将 Headers 转换为 JSON 格式，用于存储到数据库
func (h Headers) Value() (driver.Value, error) {
	if len(h) == 0 {
		return nil, nil
	}
	return json.Marshal(h)
}

// Scan 从数据库中读取 JSON 数据并解析为 Headers
func (h *Headers) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return errors.New("不支持的数据类型")
	}

	if len(data) == 0 || string(data) == "{}" {
		return nil
	}

	return json.Unmarshal(data, h)
}

type RuleBase struct {
	Name        string     `gorm:"column:name;size:300;comment:匹配路由名称;" json:"name"`                                      // 匹配路由名称
	TargetId    uint       `gorm:"column:target_id;int;comment:目标路由;" json:"target_id"`                                   // 目标地址编码
	TargetRoute string     `gorm:"column:target_route;size:300;comment:目标路由;" json:"target_route"`                        // 目标路由
	Port        uint       `gorm:"column:port;type:int;comment:端口号;" json:"port"`                                         // 端口号                                      // 端口信息编码
	Method      string     `gorm:"column:method;size:100;comment:请求方法;" json:"method"`                                    // 请求方法
	Status      bool       `gorm:"column:status;type:int;comment:状态 0 停用 1 启用;" json:"status"`                            // 状态
	MatchType   uint       `gorm:"column:match_type;type:int;comment:匹配类型 (0: 精确匹配, 1: 前缀匹配, 2: 正则匹配)" json:"match_type"` // 匹配类型 (0: 精确匹配, 1: 前缀匹配, 2: 正则匹配)
	Remark      string     `gorm:"column:remark;size:-1;comment:备注;" json:"remark"`                                       // 备注
	Target      TargetInfo `gorm:"foreignKey:TargetId" json:"target"`                                                     // 目标地址
	Order       uint       `gorm:"column:order;type:int;comment:排序;" json:"order"`                                        // 排序
	IsWs        bool       `gorm:"column:is_ws;type:int;comment:是否是ws;" json:"is_ws"`                                     // 是否是ws
	Headers     *Headers   `gorm:"column:headers;type:varchar(500);comment:Header 匹配规则;" json:"headers"`                  // Header 匹配规则
}

func (RuleInfo) TableName() string { return "dz_rule_info" }
