package model

import "time"

// HourlyTrafficStatistics represents traffic statistics for an hour
type HourlyTrafficStatistics struct {
	Hour          int
	Minute        int
	TotalRequests int
	TotalInBytes  int
	TotalOutBytes int
}

// Statistics 表示端口流量统计
type Statistics struct {
	TotalRequests int       `json:"total_requests"`
	TotalInBytes  int       `json:"total_in_bytes"`
	TotalOutBytes int       `json:"total_out_bytes"`
	CreateTime    time.Time `json:"create_time"`
	LastUpdated   time.Time `json:"last_updated"`
}
