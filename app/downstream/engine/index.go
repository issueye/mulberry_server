package engine

import (
	"context"
	"fmt"
	"mulberry/host/app/downstream/model"
	"mulberry/host/app/downstream/requests"
	"mulberry/host/app/downstream/service"
	"mulberry/host/global"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type ActionType int

const (
	AT_START  ActionType = iota // 启动
	AT_STOP                     // 停用
	AT_RELOAD                   // 重载
)

type Port struct {
	model.PortInfo
	Action ActionType
}

// 存放 http server 对象
var (
	Servers  = new(sync.Map)
	PortChan = make(chan *Port, 10)
)

func Start(ctx context.Context) {
	go func(c context.Context) {
		for {
			select {
			case p := <-PortChan: // 接收端口号信息
				{
					switch p.Action {
					case AT_START:
						RunServer(p.PortInfo)
					case AT_STOP:
						StopServer(p.PortInfo)
					case AT_RELOAD:
						ReloadServer(p.PortInfo)
					}
				}

			case <-c.Done(): // 主程序退出
				{
					return
				}
			}
		}
	}(ctx)
}

type muxHandler = func(http.ResponseWriter, *http.Request)

type GrapeEngine struct {
	PortId      uint         // 端口号编码
	Port        uint         // 端口号
	UseGzip     bool         // 使用GZIP
	Engine      *gin.Engine  // gin
	*mux.Router              // mux 对象
	Rules       []*RouteRule // 节点规则列表
}

type RouteRule struct {
	Name      string              `json:"name"`    // 匹配规则
	Target    string              `json:"target"`  // 目标地址
	Route     string              `json:"route"`   // 路由
	Page      string              `json:"Page"`    // 节点
	Method    string              `json:"method"`  // 方法
	HttpProxy *CustomReverseProxy `json:"proxy"`   // HTTP代理转发
	WsProxy   *WebsocketProxy     `json:"wsProxy"` // WS代理转发
	Handler   muxHandler          // 方法
}

func NewGrapeEngine(port model.PortInfo) *GrapeEngine {
	en := &GrapeEngine{
		PortId:  port.ID,
		Port:    port.Port,
		UseGzip: port.UseGzip,
		Rules:   make([]*RouteRule, 0),
	}

	en.Router = mux.NewRouter()
	// 处理 未匹配到的路由

	return en
}

func (grape *GrapeEngine) Init() error {

	if grape.UseGzip {
		// err := grape.GinGzipFilter()
		// if err != nil {
		// 	return err
		// }
	}

	// 自定义路由
	err := grape.CustomRoutes()
	if err != nil {
		return err
	}
	// 静态页面
	return grape.Pages()
}

type Page struct {
	RoutePath  string // 路径
	StaticPath string // 静态资源路径
}

// GinPages
// 加载页面
func (grape *GrapeEngine) Pages() error {
	// 处理页面
	pageSrv := service.NewPage()
	pageList, err := pageSrv.GetDatasByMap(map[string]any{
		"port":   grape.Port,
		"status": 1,
	})
	if err != nil {
		return err
	}

	vSrv := service.NewPageVersion()
	for _, page := range pageList {
		versionInfo, err := vSrv.GetByMap(map[string]any{
			"page_id": page.ID,
			"version": page.Version,
		})
		if err != nil {
			global.Logger.Sugar().Errorf("页面[%s]未找到激活版本[%s] %s", page.Title, page.Version, err.Error())
			continue
		}
		// 在使用版本路由
		path := ""
		if page.UseVersionRoute == 1 {
			path = fmt.Sprintf("/%s/%s/", page.Name, page.Version)
		} else {
			path = fmt.Sprintf("/%s/", page.Name)
		}
		grape.PathPrefix(path).Handler(http.StripPrefix(path, http.FileServer(http.Dir(versionInfo.PagePath))))
		// grape.Engine.StaticFS(path, http.Dir(versionInfo.PagePath))
		global.Logger.Sugar().Infof("页面[%s]加载成功 路径：%s", page.Title, path)
	}

	return nil
}

func (rule *RouteRule) replace(key, value string) string {
	v := value
	v = strings.TrimPrefix(v, "/")
	return strings.ReplaceAll(rule.Route, fmt.Sprintf("{%s}", key), v)
}

func (grape *GrapeEngine) CustomRoutes() error {
	ruleSrv := service.NewRule()

	ruleList, err := ruleSrv.GetDatas(&requests.QueryRuleDatas{
		Port:   grape.Port,
		Status: true,
	})
	if err != nil {
		return err
	}

	targetSrv := service.NewTarget()
	for _, rule := range ruleList {
		tInfo, err := targetSrv.GetByField("id", rule.TargetId)
		if err != nil {
			global.Logger.Sugar().Errorf("规则[%s]未找到目标[%s] %s", rule.Name, rule.TargetId, err.Error())
			continue
		}

		custom := &RouteRule{
			Name:   rule.Name,
			Target: tInfo.Name,
			Route:  rule.TargetRoute,
			Method: rule.Method,
		}

		// websocket 代码
		global.Logger.Sugar().Infof("[%s] 是否是WS: %v", rule.Name, rule.IsWs)
		if rule.IsWs {
			//TODO url 需要重新处理一下
			addr := fmt.Sprintf("ws://%s%s", custom.Target, rule.TargetRoute)
			wsProxy, err := NewProxy(addr, nil)
			if err != nil {
				global.Logger.Sugar().Errorf("创建WS代理失败 %s", err.Error())
				continue
			}

			custom.WsProxy = wsProxy
			custom.Handler = func(w http.ResponseWriter, r *http.Request) {
				custom.WsProxy.ServeHTTP(w, r)
			}
		} else {
			// http 代理
			addr := fmt.Sprintf("http://%s", custom.Target)
			httpProxy, err := NewReverseProxy(addr)
			if err != nil {
				global.Logger.Sugar().Error("创建HTTP代理失败 %s", err.Error())
				continue
			}

			custom.HttpProxy = httpProxy
			custom.Handler = func(w http.ResponseWriter, r *http.Request) {
				if strings.HasSuffix(custom.Route, "/*path") {
					targetName := strings.TrimSuffix(custom.Route, "/*path")
					r.URL.Path = targetName + r.URL.Path
				} else {
					vars := mux.Vars(r)

					route := ""
					if len(vars) == 0 {
						route = custom.Route
					} else {
						for key, val := range vars {
							route = custom.replace(key, val)
						}
					}

					r.URL.Path = route
				}

				custom.HttpProxy.ServeHTTP(w, r)
			}
		}

		grape.Rules = append(grape.Rules, custom)
	}

	for _, custom := range grape.Rules {
		var r *mux.Route

		global.Logger.Sugar().Infof("路由:%s", custom.Name)

		// 如果最后是 * 则表示未前缀匹配
		if strings.HasSuffix(custom.Name, "/*path") {
			name := strings.TrimSuffix(custom.Name, "/*path")
			r = grape.PathPrefix(name).Handler(http.StripPrefix(name, http.HandlerFunc(custom.Handler)))
		} else {
			r = grape.HandleFunc(custom.Name, custom.Handler)
		}

		switch strings.ToUpper(custom.Method) {
		case "POST":
			r.Methods("POST")
		case "GET":
			r.Methods("GET")
		case "PUT":
			r.Methods("PUT")
		case "PATCH":
			r.Methods("PATCH")
		case "DELETE":
			r.Methods("DELETE")
		}
	}

	return nil
}

func (grape *GrapeEngine) Run() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", grape.Port),
		Handler: grape,
	}

	// 存放到 map 中
	Servers.Store(grape.PortId, server)

	err := server.ListenAndServe()
	if err != nil {
		global.Logger.Sugar().Errorf("启动服务失败 %s", err.Error())
		return err
	}

	return nil
}

func ReloadServer(port model.PortInfo) {
	global.Logger.Sugar().Infof("[%d]端口号开始重启...", port.Port)

	StopServer(port)
	RunServer(port)
}

func StopServer(port model.PortInfo) {
	global.Logger.Sugar().Infof("[%d]端口号停用服务...", port.Port)

	value, ok := Servers.Load(port.ID)
	if ok {
		server := value.(*http.Server)
		server.Shutdown(context.Background())

		// 删除对象
		Servers.Delete(port.ID)
	}
}

func runServer(port model.PortInfo) {
	grape := NewGrapeEngine(port)
	err := grape.Init()
	if err != nil {
		global.Logger.Sugar().Errorf("初始化失败 %s", err.Error())
		return
	}

	err = grape.Run()
	if err != nil {
		global.Logger.Sugar().Errorf("启动失败 %s", err.Error())
		return
	}
}

func RunServer(port model.PortInfo) {
	global.Logger.Sugar().Infof("[%d]端口号启用服务...", port.Port)
	go runServer(port)
}
