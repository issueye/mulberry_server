# 项目名称

## 项目简介
该项目是一个用Go语言编写的多模块应用程序，包含了多个子目录和文件。项目结构清晰，分为admin和downstream两个主要部分，每个部分又包含controller、logic、model、requests、response、route、service等子目录。此外，还有common、global、initialize、pages、root、web等目录。

## 目录结构
```
.
├── bin/
├── docs/
├── internal/
│   ├── app/
│   │   ├── admin/
│   │   │   ├── controller/
│   │   │   ├── logic/
│   │   │   ├── model/
│   │   │   ├── requests/
│   │   │   ├── response/
│   │   │   ├── route/
│   │   │   └── service/
│   │   ├── downstream/
│   │   │   ├── controller/
│   │   │   ├── engine/
│   │   │   ├── logic/
│   │   │   ├── model/
│   │   │   ├── requests/
│   │   │   ├── route/
│   │   │   └── service/
│   ├── common/
│   ├── global/
│   ├── initialize/
│   └── pages/
├── root/
├── web/
└── winappres/
```

## 安装与使用
1. 克隆项目到本地：
   ```sh
   git clone <repository-url>
   ```
2. 进入项目目录：
   ```sh
   cd <project-directory>
   ```
3. 安装依赖：
   ```sh
   go mod tidy
   ```
4. 运行项目：
   ```sh
   go run main.go
   ```

## 详细说明
### Admin 部分
- **controller**: 处理HTTP请求的控制器。
  - `auth.go`: 处理认证相关的请求。
  - `common.go`: 处理通用请求。
  - `dict.go`: 处理字典相关的请求。
  - `index.go`: 处理首页相关的请求。
  - `menu.go`: 处理菜单相关的请求。
  - `role.go`: 处理角色相关的请求。
  - `settings.go`: 处理设置相关的请求。
  - `user.go`: 处理用户相关的请求。
- **logic**: 业务逻辑层，处理具体的业务需求。
  - `dict.go`: 字典业务逻辑。
  - `index.go`: 首页业务逻辑。
  - `menu.go`: 菜单业务逻辑。
  - `role.go`: 角色业务逻辑。
  - `settings.go`: 设置业务逻辑。
  - `user.go`: 用户业务逻辑。
- **model**: 数据模型，定义数据库表结构。
  - `dict.go`: 字典数据模型。
  - `menu.go`: 菜单数据模型。
  - `user.go`: 用户数据模型。
- **requests**: 请求参数的定义和验证。
  - `dict.go`: 字典请求参数。
  - `index.go`: 首页请求参数。
  - `menu.go`: 菜单请求参数。
- **response**: 响应数据的定义。
  - `index.go`: 首页响应数据。
- **route**: 路由定义，配置URL路径和对应的控制器。
  - `index.go`: 首页路由配置。
- **service**: 服务层，封装具体的业务服务。
  - `dict.go`: 字典服务。
  - `menu.go`: 菜单服务。
  - `role.go`: 角色服务。
  - `user.go`: 用户服务。

### Downstream 部分
- **controller**: 处理HTTP请求的控制器。
  - `gzip_filter.go`: 处理Gzip过滤器相关的请求。
  - `page.go`: 处理页面相关的请求。
  - `port.go`: 处理端口相关的请求。
  - `rule.go`: 处理规则相关的请求。
  - `target.go`: 处理目标相关的请求。
  - `traffic.go`: 处理流量相关的请求。
- **engine**: 核心引擎，处理底层逻辑。
  - `byte_slice_pool.go`: 字节切片池。
  - `http_proxy.go`: HTTP代理。
  - `index.go`: 引擎入口。
  - `tcp_proxy.go`: TCP代理。
  - `tls_config.go`: TLS配置。
  - `ws_proxy.go`: WebSocket代理。
- **logic**: 业务逻辑层，处理具体的业务需求。
  - `gzip_filter.go`: Gzip过滤器业务逻辑。
  - `page.go`: 页面业务逻辑。
  - `port.go`: 端口业务逻辑。
  - `rule.go`: 规则业务逻辑。
  - `target.go`: 目标业务逻辑。
  - `traffic.go`: 流量业务逻辑。
- **model**: 数据模型，定义数据库表结构。
  - `cert.go`: 证书数据模型。
  - `gzip_filter.go`: Gzip过滤器数据模型。
  - `http_messages.go`: HTTP消息数据模型。
  - `page.go`: 页面数据模型。
  - `port.go`: 端口数据模型。
  - `rule.go`: 规则数据模型。
  - `target.go`: 目标数据模型。
- **requests**: 请求参数的定义和验证。
  - `cert.go`: 证书请求参数。
  - `gzip_filter.go`: Gzip过滤器请求参数。
  - `page_version.go`: 页面版本请求参数。
  - `page.go`: 页面请求参数。
  - `port.go`: 端口请求参数。
  - `rule.go`: 规则请求参数。
  - `target.go`: 目标请求参数。
  - `traffic.go`: 流量请求参数。
- **route**: 路由定义，配置URL路径和对应的控制器。
  - `index.go`: 路由配置。
- **service**: 服务层，封装具体的业务服务。
  - `cert.go`: 证书服务。
  - `gzip_filter.go`: Gzip过滤器服务。
  - `page_version.go`: 页面版本服务。
  - `page.go`: 页面服务。
  - `port.go`: 端口服务。
  - `rule.go`: 规则服务。
  - `target.go`: 目标服务。

### Common 部分
- **config**: 配置文件和参数的定义。
  - `arr.go`: 数组配置。
  - `config.go`: 配置文件。
  - `index.go`: 配置入口。
  - `param.go`: 参数配置。
- **controller**: 公共控制器。
  - `index.go`: 公共控制器入口。
- **middleware**: 中间件，处理请求的预处理和后处理。
  - `cros.go`: 跨域中间件。
  - `jwt_auth.go`: JWT认证中间件。
  - `limit_handler.go`: 限流中间件。
  - `logger.go`: 日志中间件。
- **model**: 公共数据模型。
  - `index.go`: 公共数据模型入口。
- **route**: 公共路由定义。
  - `index.go`: 公共路由配置。
- **service**: 公共服务层。
  - `index.go`: 公共服务入口。

### Initialize 部分
- **config**: 初始化配置。
  - `config.go`: 配置初始化。
- **db**: 数据库初始化。
  - `db.go`: 数据库初始化逻辑。
- **http_server**: HTTP服务器初始化。
  - `http_server.go`: HTTP服务器初始化逻辑。
- **logger**: 日志初始化。
  - `logger.go`: 日志初始化逻辑。
- **proxy**: 代理初始化。
  - `proxy.go`: 代理初始化逻辑。
- **runtime**: 运行时初始化。
  - `runtime.go`: 运行时初始化逻辑。
- **store**: 存储初始化。
  - `store.go`: 存储初始化逻辑。

### Pages 部分
- **home**: 首页相关文件。
  - `FrmHome.gfm`: 首页表单。
  - `FrmHome.go`: 首页逻辑。
  - `FrmHomeImpl.go`: 首页实现。
  - `service.go`: 首页服务。

## 引用的三方库
- `github.com/TelenLiu/knife4j_go`: Knife4j Go 版本，用于生成API文档。
- `github.com/bwmarrin/snowflake`: 用于生成分布式唯一ID。
- `github.com/cockroachdb/pebble`: Pebble 是一个键值存储库。
- `github.com/dgrijalva/jwt-go`: 用于生成和验证JWT。
- `github.com/didip/tollbooth`: 用于限流。
- `github.com/gin-contrib/cors`: Gin框架的CORS中间件。
- `github.com/gin-contrib/zap`: Gin框架的Zap日志中间件。
- `github.com/gin-gonic/gin`: Gin Web框架。
- `github.com/glebarez/sqlite`: SQLite数据库驱动。
- `github.com/godoes/gorm-oracle`: GORM的Oracle数据库驱动。
- `github.com/gorilla/mux`: HTTP请求路由器和分发器。
- `github.com/ncruces/go-sqlite3`: SQLite数据库驱动。
- `github.com/satori/go.uuid`: 用于生成UUID。
- `github.com/spf13/cast`: 用于类型转换。
- `github.com/swaggo/swag`: 用于生成Swagger文档。
- `github.com/ying32/govcl`: 用于创建GUI应用程序。
- `go.uber.org/zap`: 高性能日志库。
- `golang.org/x/crypto`: Go的扩展加密包。
- `golang.org/x/sys`: Go的系统调用包。
- `gopkg.in/natefinch/lumberjack.v2`: 日志轮转库。
- `gorm.io/driver/mysql`: GORM的MySQL数据库驱动。
- `gorm.io/driver/sqlserver`: GORM的SQL Server数据库驱动。
- `gorm.io/gorm`: ORM库。

## 贡献
欢迎提交问题（Issues）和合并请求（Pull Requests）来贡献代码。

## 许可证
该项目使用MIT许可证。