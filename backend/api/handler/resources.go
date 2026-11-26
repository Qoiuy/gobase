package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Qoiuy/go-base/backend/api/config"
	"github.com/Qoiuy/go-base/backend/api/router"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"html/template"
)

// Here you can initialize and manage system resources like Redis, MySQL, etc.

// NewApp 新建App
func NewApp(ctx context.Context) error {
	app := iris.New()
	conf := iris.Configuration{
		// 解析客户端真实ip
		RemoteAddrHeaders: map[string]bool{
			"X-Real-Ip":        true,
			"X-Forwarded-For":  true,
			"CF-Connecting-IP": true,
		}}
	app.Configure(iris.WithConfiguration(conf))

	// ========初始化资源=============
	db := InitDB()
	redisConn := InitRedisCache()

	// 全局 middleware
	app.Use(recover.New())

	// 打印日志
	//app.Use(irislog2.IrisLog())
	app.Use(HttpInterceptor)

	// health check TODO?
	//healthURL := "healthcheck"

	// 监控界面
	app.Get("/metrics", iris.FromStd(promhttp.Handler()))
	// v1
	v1URL := "api/v1"
	if len(config.Opts.BaseURL) > 0 {
		v1URL = config.Opts.BaseURL + "/" + v1URL
	}

	// Allow Cors
	crs := cors.AllowAll()

	// ui 页面入口
	uiParty := app.Party("/", crs).AllowMethods(iris.MethodOptions)
	uiRouter := router.NewUIRouter(redisConn, db, config.Opts)
	mvc.Configure(uiParty.Party(uiRouter.GetParty()), uiRouter.Router())
	html := iris.HTML("ui/html", ".tmpl")
	html.AddFunc("toJSON", toJSON) // 添加html解析模板
	app.RegisterView(html)
	app.HandleDir("/data", "./data")
	app.HandleDir("/ui/icon", "./ui/icon")
	app.HandleDir("/ui/js", "./ui/js")

	// 设置主页重定向
	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect("/search.html")
	})

	// hello 页面示例
	app.Get("/hello.html", func(ctx iris.Context) {
		if err := ctx.ServeFile("./ui/html/hello.html", false); err != nil {
			ctx.StatusCode(iris.StatusNotFound)
			_, _ = ctx.WriteString("hello.html 不存在")
		}
	})

	v1 := app.Party("/"+v1URL, crs).AllowMethods(iris.MethodOptions)

	// 路由入口
	apiRouter := router.NewApiRouter(redisConn, db, config.Opts.Debug)
	apiRouter.EntryRouter(v1)

	go func() {
		select {
		case <-ctx.Done():
			_ = app.Shutdown(ctx)
		}
	}()

	port := config.Opts.Port
	if port <= 0 {
		port = 8080
	}
	return app.Run(iris.Addr(fmt.Sprintf(":%v", config.Opts.Port)), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

func toJSON(data interface{}) (template.JS, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return template.JS(jsonData), nil
}
