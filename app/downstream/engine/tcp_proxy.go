package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mulberry/app/downstream/model"
	"mulberry/global"
	"mulberry/pkg/utils"
	"net"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TCPProxy struct {
	TargetAddr string
}

func NewTCPProxy(targetAddr string) *TCPProxy {
	return &TCPProxy{
		TargetAddr: targetAddr,
	}
}

func (p *TCPProxy) Serve(conn net.Conn) {
	defer conn.Close()

	// Create traffic statistics
	traffic := model.NewTrafficStatistics()
	traffic.ID = uuid.NewV1().String()
	traffic.Request.Time = time.Now()
	traffic.Request.Method = "TCP"
	traffic.Request.Path = p.TargetAddr

	// Connect to target
	targetConn, err := net.Dial("tcp", p.TargetAddr)
	if err != nil {
		global.Logger.Sugar().Errorf("Failed to connect to target: %s", err.Error())
		return
	}
	defer targetConn.Close()

	// Setup context with traffic data
	ctx := context.Background()
	ctx = context.WithValue(ctx, "traffic", traffic)

	// Start proxying
	go p.copyData(ctx, conn, targetConn, traffic, true) // client -> target
	p.copyData(ctx, targetConn, conn, traffic, false)   // target -> client

	// Save traffic stats
	trafficInfo, _ := json.Marshal(traffic)
	nowDateStr := time.Now().Format("2006-01-02")
	saveKey := fmt.Sprintf("TRAFFIC:%s:%d", nowDateStr, utils.GenID())
	global.STORE.Set(saveKey, trafficInfo)
}

func (p *TCPProxy) copyData(ctx context.Context, src, dst net.Conn, traffic *model.TrafficStatistics, isRequest bool) {
	buf := make([]byte, 32*1024)
	for {
		n, err := src.Read(buf)
		if err != nil {
			if err != io.EOF {
				global.Logger.Sugar().Errorf("Connection read error: %s", err.Error())
			}
			break
		}

		if isRequest {
			traffic.Request.InBodyBytes += int64(n)
		} else {
			traffic.Response.OutBodyBytes += int64(n)
		}

		_, err = dst.Write(buf[:n])
		if err != nil {
			global.Logger.Sugar().Errorf("Connection write error: %s", err.Error())
			break
		}
	}
}
