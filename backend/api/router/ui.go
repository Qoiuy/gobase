package router

import (
	"github.com/Qoiuy/go-base/backend/api/config"
	"github.com/Qoiuy/go-base/backend/mysql"
	"github.com/Qoiuy/go-base/backend/redis"
	"github.com/kataras/iris/v12/mvc"
)

// UIRouter 页面路由。ui/ 目录下目前还没有页面，这里先占位，
// 后续在 Router() 里注册对应的 controller 即可。
type UIRouter struct {
	RedisConn *redis.RedisServer
	Db        *mysql.Repository
	Opts      *config.Options
}

// NewUIRouter 新建页面路由
func NewUIRouter(redisConn *redis.RedisServer, repository *mysql.Repository, opts *config.Options) *UIRouter {
	return &UIRouter{
		RedisConn: redisConn,
		Db:        repository,
		Opts:      opts,
	}
}

// GetParty 返回路由分组
func (r *UIRouter) GetParty() string {
	return "ui"
}

// Router 注册页面 controller
func (r *UIRouter) Router() func(app *mvc.Application) {
	return func(app *mvc.Application) {}
}
