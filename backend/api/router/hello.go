package router

import (
	"github.com/Qoiuy/go-base/backend/api/config"
	"github.com/Qoiuy/go-base/backend/api/controller/hello"
	"github.com/Qoiuy/go-base/backend/mysql"
	"github.com/Qoiuy/go-base/backend/redis"
	"github.com/kataras/iris/v12/mvc"
)

// HelloRouter hello 模块的路由
type HelloRouter struct {
	RedisConn *redis.RedisServer
	Db        *mysql.Repository
	Opts      *config.Options
}

// NewHelloRouter 新建 hello 路由
func NewHelloRouter(redisConn *redis.RedisServer, repository *mysql.Repository, opts *config.Options) *HelloRouter {
	return &HelloRouter{
		RedisConn: redisConn,
		Db:        repository,
		Opts:      opts,
	}
}

// GetParty 返回路由分组，最终访问路径为 /api/v1/hello
func (r *HelloRouter) GetParty() string {
	return "hello"
}

// Router 注册 hello 的 controller
func (r *HelloRouter) Router() func(app *mvc.Application) {
	return func(app *mvc.Application) {
		app.Handle(new(hello.Controller))
	}
}
