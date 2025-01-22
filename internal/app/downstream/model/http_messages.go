package model

import (
	"encoding/json"
	"time"
)

// PortStatistics 端口转发流量统计
type PortStatistics struct {
	TotalRequests int64        `json:"total_requests"`  // 总请求数
	TotalInBytes  int64        `json:"total_in_bytes"`  // 总输入字节数
	TotalOutBytes int64        `json:"total_out_bytes"` // 总输出字节数
	Ports         []int        `json:"ports"`           // 端口列表
	Data          []*PortStats `json:"data"`            // 按端口统计
}

func (ps *PortStatistics) GetPort(port int) *PortStats {
	for index, p := range ps.Ports {
		if p == port {
			return ps.Data[index]
		}
	}

	ps.Ports = append(ps.Ports, port)
	rtn := &PortStats{
		Requests: 0,
		InBytes:  0,
		OutBytes: 0,
	}
	ps.Data = append(ps.Data, rtn)

	return rtn
}

// PortStats 单个端口的流量统计
type PortStats struct {
	Requests int64 `json:"requests"`  // 请求数
	InBytes  int64 `json:"in_bytes"`  // 输入字节数
	OutBytes int64 `json:"out_bytes"` // 输出字节数
}

type TrafficStatistics struct {
	ID       string        `json:"id"`       // 数据编码
	Port     int           `json:"port"`     // 端口
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
