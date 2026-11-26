package router

import (
	"strings"

	"github.com/Qoiuy/go-base/backend/api/config"
	"github.com/Qoiuy/go-base/backend/mysql"
	"github.com/Qoiuy/go-base/backend/redis"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type ApiRouter interface {
	Router() func(app *mvc.Application)
	GetParty() string
}

type ApiVPRouter struct {
	apiRouters []ApiRouter
	debug      bool
	RedisConn  *redis.RedisServer
	Db         *mysql.Repository
}

func NewApiRouter(
	redisConn *redis.RedisServer,
	repository *mysql.Repository,
	debug bool) *ApiVPRouter {

	return &ApiVPRouter{
		RedisConn: redisConn,
		debug:     debug,
		Db:        repository,
	}
}

func (r *ApiVPRouter) load() {
	r.apiRouters = make([]ApiRouter, 0)
	hR := NewHelloRouter(r.RedisConn, r.Db, config.Opts)
	r.apiRouters = append(r.apiRouters,
		hR,
	)

}

func (r *ApiVPRouter) EntryRouter(v1 router.Party) {
	r.load()
	for _, apiRouter := range r.apiRouters {
		mvc.Configure(v1.Party(partyPath(apiRouter.GetParty())), apiRouter.Router())
	}
}

// partyPath iris v12.1.8 的 Party 是把父路径和子路径直接拼接的，
// 子分组如果不带前导 "/"，会被拼成 /api/v1hello 这种路径。
func partyPath(relative string) string {
	if strings.HasPrefix(relative, "/") {
		return relative
	}
	return "/" + relative
}
