package hello

import (
	"github.com/Qoiuy/go-base/backend/api/controller/resp"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// Controller hello 控制器
type Controller struct {
	Ctx iris.Context
}

// BeforeActivation 注册路由，对应 GET /api/v1/hello
func (c *Controller) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/", "Get")
}

// Get 处理 GET /api/v1/hello
func (c *Controller) Get() {
	resp.SuccessResponse(c.Ctx, iris.Map{
		"message": "hello",
	})
}
