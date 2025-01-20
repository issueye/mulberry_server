package engine

import (
	"context"
	"fmt"
	"mulberry/app/downstream/model"
	"mulberry/app/downstream/requests"
	"mulberry/app/downstream/service"
	"mulberry/global"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/liushuochen/gotable"
)

// 协议类型
type ProtocolType string

const (
	HTTP  ProtocolType = "http://"
	HTTPS ProtocolType = "https://"
	WS    ProtocolType = "ws://"
	WSS   ProtocolType = "wss://"
)

type ActionType int

const (
	AT_START  ActionType = iota // 启动
	AT_STOP                     // 停用
	AT_RELOAD                   // 重载
)

type Port struct {
	model.PortInfo
	Action   ActionType
	Protocol string
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

type MatchType int

const (
	MT_EXACT    MatchType = 0 // 精确匹配
	MT_PREFIX             = 1 // 前缀匹配
	MT_REGEX              = 2 // 正则匹配
	MT_WILDCARD           = 3 // 通配符匹配
)

type RouteRule struct {
	Name      string              `json:"name"`      // 匹配规则
	Target    string              `json:"target"`    // 目标地址
	Route     string              `json:"route"`     // 路由
	Page      string              `json:"Page"`      // 节点
	Method    string              `json:"method"`    // 方法
	MatchType MatchType           `json:"matchType"` // 匹配类型
	Wildcard  bool                `json:"wildcard"`  // 是否使用通配符匹配
	Headers   map[string]string   `json:"headers"`   // Header 匹配规则
	HttpProxy *CustomReverseProxy `json:"proxy"`     // HTTP代理转发
	WsProxy   *WebsocketProxy     `json:"wsProxy"`   // WS代理转发
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
			Name:      rule.Name,
			Target:    tInfo.Name,
			Route:     rule.TargetRoute,
			Method:    rule.Method,
			Headers:   parseHeaders(rule.Headers), // 添加 Header 匹配规则
			MatchType: MatchType(rule.MatchType),
		}

		// 处理 WebSocket 路由
		if rule.IsWs {
			if err := grape.setupWebSocketProxy(custom, rule); err != nil {
				global.Logger.Sugar().Errorf("创建WS代理失败: %s", err.Error())
				continue
			}
		} else {
			if err := grape.setupHTTPProxy(custom); err != nil {
				global.Logger.Sugar().Errorf("创建HTTP代理失败: %s", err.Error())
				continue
			}
		}

		grape.Rules = append(grape.Rules, custom)
	}

	grape.registerRoutes()
	return nil
}

func parseHeaders(headers *model.Headers) map[string]string {
	result := make(map[string]string)
	if headers != nil {
		for _, v := range *headers {
			result[v.K] = v.V
		}
	}
	return result
}

func (grape *GrapeEngine) setupWebSocketProxy(custom *RouteRule, rule *model.RuleInfo) error {
	addr := fmt.Sprintf("%s%s%s", WS, custom.Target, rule.TargetRoute)
	wsProxy, err := NewProxy(addr, nil)
	if err != nil {
		return err
	}
	custom.WsProxy = wsProxy
	custom.Handler = wsProxy.ServeHTTP
	return nil
}

func (grape *GrapeEngine) setupHTTPProxy(custom *RouteRule) error {
	// 处理 HTTP 路由
	var tlsConfig *TLSConfig
	var err error

	// Check if HTTPS is needed
	if strings.HasPrefix(custom.Target, "https://") {
		// Load certificate from database
		certSrv := service.NewCert()
		certInfo, err := certSrv.GetByField("name", custom.Target)
		if err != nil {
			global.Logger.Sugar().Error("获取证书失败 %s", err.Error())
			return err
		}

		tlsConfig, err = NewTLSConfig(&model.CertInfo{
			CertBase: certInfo.CertBase,
		})
		if err != nil {
			global.Logger.Sugar().Error("创建TLS配置失败 %s", err.Error())
			return err
		}
	}

	addr := fmt.Sprintf("%s%s", HTTP, custom.Target)
	httpProxy, err := NewReverseProxy(addr, tlsConfig)
	if err != nil {
		global.Logger.Sugar().Error("创建HTTP代理失败 %s", err.Error())
		return err
	}

	custom.HttpProxy = httpProxy
	custom.Handler = func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		switch custom.MatchType {
		case MT_PREFIX:
			{
				// 替换匹配路由 Name 的前缀为目标路 Route 由的前缀
				path = strings.Replace(path, custom.Name, custom.Route, 1)
				r.URL.Path = path
			}
		case MT_REGEX:
			{
				// 将正则匹配的路由 替换为目标路由
				// 匹配正则表达式
				re := regexp.MustCompile(custom.Name)
				// 替换匹配的部分为目标路由
				path = re.ReplaceAllString(path, custom.Route)
				r.URL.Path = path
			}
		}

		custom.HttpProxy.ServeHTTP(w, r)
	}

	return nil
}

func (grape *GrapeEngine) registerRoutes() {
	table, err := gotable.Create("方法", "代理路由", "目标路由")
	if err != nil {
		return
	}

	values := []map[string]string{}

	for _, custom := range grape.Rules {
		values = append(values, map[string]string{
			"方法":   custom.Method,
			"代理路由": custom.Name,
			"目标路由": custom.Route,
		})
		grape.registerRoute(custom)
	}

	table.AddRows(values)
	tableStr := table.String()
	global.Logger.Sugar().Infof("\n%s", tableStr)
}

func (grape *GrapeEngine) registerRoute(custom *RouteRule) {
	var r *mux.Route
	switch custom.MatchType {
	case MT_EXACT:
		r = grape.HandleFunc(custom.Name, custom.Handler)
	case MT_PREFIX:
		r = grape.PathPrefix(custom.Name).Handler(http.HandlerFunc(custom.Handler))
	case MT_REGEX:
		r = grape.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
			matched, _ := regexp.MatchString(custom.Name, r.URL.Path)
			return matched
		}).Handler(http.HandlerFunc(custom.Handler))
	default:
		r = grape.HandleFunc(custom.Name, custom.Handler)
	}

	method := strings.ToUpper(custom.Method)
	if method != "ANY" {
		r.Methods(strings.ToUpper(custom.Method))
	}

	for k, v := range custom.Headers {
		r.Headers(k, v)
	}
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
