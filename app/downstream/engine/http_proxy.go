package engine

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mulberry/common/utils"
	"mulberry/host/app/downstream/model"
	"mulberry/host/global"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	uuid "github.com/satori/go.uuid"
)

type CustomReverseProxy struct {
	*httputil.ReverseProxy
}

func NewReverseProxy(targetURL string) (*CustomReverseProxy, error) {
	// 解析目标URL
	target, err := url.Parse(targetURL)
	if err != nil {
		global.Logger.Sugar().Errorf("解析目标URL失败 %s", err.Error())
		return nil, err
	}

	// 创建反向代理实例
	proxy := &CustomReverseProxy{
		ReverseProxy: &httputil.ReverseProxy{
			Director:       director(target),
			ModifyResponse: modifyResponse,
			Transport:      &myRoundTripper{http.DefaultTransport},
		},
	}

	return proxy, nil
}

func (p *CustomReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// 调用父类的ServeHTTP方法，传递自定义的ResponseWriter
	p.ReverseProxy.ServeHTTP(rw, req)
}

type CustomResponseWriter struct {
	http.ResponseWriter
	Body []byte
}

func (crw *CustomResponseWriter) Write(b []byte) (int, error) {
	crw.Body = append(crw.Body, b...)
	return crw.ResponseWriter.Write(b)
}

type myRoundTripper struct {
	http.RoundTripper
}

func (m *myRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	traffic := model.NewTrafficStatistics()
	traffic.Request.Time = time.Now()
	traffic.ID = uuid.NewV1().String()

	headerSize := int64(0)
	url := req.URL.Path
	if req.URL.RawQuery != "" {
		traffic.Request.Query = req.URL.RawQuery
	}

	traffic.Request.Path = url
	traffic.Request.Method = req.Method

	for key, value := range req.Header {
		headerSize += int64(len(fmt.Sprintf("%s: %s\r\n", key, value)))
		traffic.Request.Header[key] = value
	}

	traffic.Request.InHeaderBytes = headerSize

	if req.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			global.Logger.Sugar().Errorf("读取请求内容失败 %s", err.Error())
			return m.RoundTripper.RoundTrip(req)
		}

		bodyBuf := bytes.NewBuffer(body)

		traffic.Request.InBodyBytes = int64(bodyBuf.Len())
		traffic.Request.Body = bodyBuf.String()

		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx, "traffic", traffic)
	req = req.WithContext(ctx)

	resp, err := m.RoundTripper.RoundTrip(req)
	if err != nil {
		global.Logger.Sugar().Errorf("转发失败 %s", err.Error())
	}

	return resp, err
}

func director(target *url.URL) func(*http.Request) {
	return func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
}

func modifyResponse(resp *http.Response) error {
	var (
		traffic *model.TrafficStatistics
		ok      bool
	)

	if resp != nil {
		// 读取请求主体和头部的大小
		ctx := resp.Request.Context()
		data := ctx.Value("traffic")
		if data != nil {
			traffic, ok = data.(*model.TrafficStatistics)
			if !ok {
				traffic = model.NewTrafficStatistics()
			}
		}

		defer func() {
			// 将数据通过管道的方式传入 global.Index
			// global.IndexDB <- traffic
			trafficInfo, _ := json.Marshal(traffic)
			nowDateStr := time.Now().Format("2006-01-02")
			saveKey := fmt.Sprintf("TRAFFIC:%s:%d", nowDateStr, utils.GenID())
			global.STORE.Set(saveKey, trafficInfo)
		}()

		for key, value := range resp.Header {
			traffic.Response.Header[key] = value
		}

		// 读取响应头部的大小
		respDump, err := httputil.DumpResponse(resp, false)
		if err != nil {
			global.Logger.Sugar().Errorf("读取返回头部失败 %s", err.Error())
			return err
		}
		traffic.Response.OutHeaderBytes = int64(len(respDump))

		if resp.Body != nil {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				global.Logger.Sugar().Errorf("读取返回内容 %s", err.Error())
				return err
			}
			bodyBuf := bytes.NewBuffer(body)
			traffic.Response.Body = bodyBuf.String()
			traffic.Response.OutBodyBytes = int64(bodyBuf.Len())
			traffic.Response.StatusCode = resp.StatusCode

			resp.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		return nil
	}

	return nil
}
