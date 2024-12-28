package model

import (
	"encoding/json"
	"time"
)

type TrafficStatistics struct {
	ID       string        `json:"id"`       // 数据编码
	Request  *HttpRequest  `json:"request"`  // 响应信息
	Response *HttpResponse `json:"response"` // 响应信息
}

func NewTrafficStatistics() *TrafficStatistics {
	return &TrafficStatistics{
		Request: &HttpRequest{
			Header: make(map[string][]string),
		},
		Response: &HttpResponse{
			Header: make(map[string][]string),
		},
	}
}

func (TrafficStatistics) FromJson(value []byte) (*TrafficStatistics, error) {
	data := TrafficStatistics{}
	err := json.Unmarshal([]byte(value), &data)
	return &data, err
}

type HttpRequest struct {
	Time          time.Time           `json:"time"`            // 请求时间
	Method        string              `json:"method"`          // 方法
	Path          string              `json:"path"`            // 路由
	Query         string              `json:"query"`           // 请求参数
	Header        map[string][]string `json:"header"`          // 请求头
	Body          string              `json:"body"`            // 请求体
	InHeaderBytes int64               `json:"in_header_bytes"` // 请求头字节数
	InBodyBytes   int64               `json:"in_body_bytes"`   // 请求体字节数
}

type HttpResponse struct {
	Header         map[string][]string `json:"header"`           // 响应头
	Body           string              `json:"body"`             // 响应体
	StatusCode     int                 `json:"status_code"`      // 状态码
	OutHeaderBytes int64               `json:"out_header_bytes"` // 响应头字节数
	OutBodyBytes   int64               `json:"out_body_bytes"`   // 响应体字节数
}
