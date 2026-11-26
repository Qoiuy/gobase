# golang 实现 web 服务的脚手架

基于 Iris v12 的 Go Web 服务脚手架，包含配置加载、MySQL / Redis 接入、统一响应、Prometheus 指标、日志封装和常用工具包。

> 当前模块路径为 `github.com/Qoiuy/go-base`。

## 环境要求

- Go 1.18 及以上（`go.mod` 声明 `go 1.18`）
- 可连接的 MySQL
- 可连接的 Redis

## 启动

1. 复制配置模板：

   ```bash
   cp config.example.yaml config.yaml
   ```

2. 填写 `mysql_*`、`redis_*` 等配置，配置项见下文。

3. 在仓库根目录启动：

   ```bash
   go run ./backend/api
   ```

默认读取工作目录下的 `config.yaml`，也可以通过环境变量 `CONFIG_FILE` 指定其它路径：

```bash
CONFIG_FILE=/etc/go-base/config.yaml go run ./backend/api
```

启动时 `handler.NewApp` 会先初始化 MySQL，再初始化 Redis；任一连接失败都会直接退出，不支持只启动 HTTP 服务。`ui/html`、`data` 等静态资源目录按相对路径注册，建议始终在仓库根目录启动。

## 配置

`config.yaml` 已被 `.gitignore` 忽略，仓库只提交 `config.example.yaml`。配置项：

| 配置项 | 示例值 | 说明 |
| --- | --- | --- |
| `debug` | `false` | `true` 时日志级别为 debug，并打开 GORM 的 SQL 日志；Prometheus 的 `env` 标签也会从 `prod` 变为 `dev` |
| `port` | `8080` | HTTP 监听端口，建议显式配置 |
| `base_url` | `""` | API 前缀。为空时为 `/api/v1`；填 `xxx` 时为 `/xxx/api/v1`。按示例填写，不要带前导 `/` |
| `redis_address` | `127.0.0.1:6379` | Redis 地址，`host:port` |
| `redis_password` | `""` | Redis 密码，无密码留空 |
| `redis_db` | `0` | Redis database 编号 |
| `mysql_address` | `127.0.0.1:3306` | MySQL 地址，`host:port` |
| `mysql_user` | `""` | MySQL 用户名 |
| `mysql_password` | `""` | MySQL 密码 |
| `mysql_database` | `""` | MySQL 数据库名 |
| `logger_address` | `""` | 预留字段，当前未接入远程日志；必须与 `logger_namespace` 同时为空或同时填写，否则启动时报错退出 |
| `logger_namespace` | `""` | 同上 |
| `logger_password` | `""` | 预留字段，当前未使用 |
| `logger_prefix` | `go-base-daemon` | flog 输出前缀；留空时使用默认值 |

> `backend/api/config.Options` 里还有 `max_connections`、`conn_max_life_time`、`max_idle_conns`、`enable_logging` 等 MySQL 连接池字段，但 `handler.InitDB` 目前没有传入这些值，实际使用 `backend/mysql` 的代码默认值。

## 访问入口

服务默认监听 `:8080`（以 `port` 配置为准）：

| 路径 | 说明 |
| --- | --- |
| `GET /api/v1/hello` | 示例接口，返回 `{"code":0,"data":{"message":"hello"},"message":""}` |
| `GET /hello.html` | 静态示例页，页面里有指向接口和指标页的链接 |
| `GET /metrics` | 应用端口上的 Prometheus 指标 |
| `0.0.0.0:8399/metrics` | `handler.NewPrometheus` 额外启动的指标服务，端口写死为 `8399` |
| `GET /` | 302 重定向到 `/search.html`；仓库里目前没有 `search.html`，调试请直接访问 `/hello.html` |

`ui/html` 注册为模板目录（后缀 `.tmpl`）；`/data`、`/ui/icon`、`/ui/js` 作为静态目录暴露。

## 目录结构

```text
backend/api/            程序入口
  config/               viper 读取 config.yaml
  controller/           MVC controller 与统一响应
  handler/              Iris App、MySQL/Redis 初始化、Prometheus
  router/               路由注册
backend/mysql/           GORM Repository 与连接池配置
backend/redis/           redigo 连接池与常用命令封装
backend/flog/            API 当前使用的 stdout 日志
backend/zlog/            zap + lumberjack 文件日志（独立实现，未被 API 引用）
backend/constant/        错误码与错误级别
backend/errorrr/         错误处理与日志分发
backend/pacp/           预留模块（当前为空）
backend/mistakeminer/   预留模块（当前为空）
backend/pdf/            预留模块（当前只有 output_images 目录）
backend/tools/          预留工具目录（当前为空）
gotool/basic/           独立 Go module：bstring、btime、filetree、uuid、constant
ui/html/                静态页面与 .tmpl 模板
ui/css/ ui/dist/ ui/icon/ ui/js/ ui/src/   前端静态目录（当前基本为空）
data/ log/              运行时目录（已在 .gitignore 中忽略）
repositories/sql/       SQL 目录（当前为空）
```

`gotool/basic` 有自己的 `go.mod`，不会包含在根模块的 `go test ./...` 中，需要在它自己的目录里执行测试。

## 编译与测试

```bash
# 根模块编译检查
go build ./...

# 通用工具包测试（独立 module）
cd gotool/basic && go test ./...
```

根模块当前没有 `*_test.go`。`go test ./...` 会在 `backend/flog` 上触发 `go vet` 的 printf 检查失败（例如 `Panicf` 没有把 `args...` 传给 `Sprintf`，以及内部 `log` 调用使用空 format），代码本身可以正常编译和运行。若只想跑测试，可用：

```bash
go test -vet=off ./...
```

## 日志

- `backend/flog` 是 API 实际使用的日志：带级别、前缀，输出到标准输出；`debug: true` 时级别为 DEBUG。
- `config.yaml` 里的 `logger_address` / `logger_namespace` 目前只做“要么都填、要么都不填”的校验，并没有建立远程日志连接；`logger_password` 未使用。
- `backend/zlog` 是另一套 zap + lumberjack 文件日志，导入时初始化，日志目标为 `./logs/go-base.log`，但 `backend/api` 没有引用它。详见 [backend/zlog/README.md](backend/zlog/README.md)。
