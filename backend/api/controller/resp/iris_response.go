package resp

import (
	"github.com/Qoiuy/go-base/backend/constant"
	"github.com/Qoiuy/go-base/backend/flog"
	"github.com/kataras/iris/v12"
)

// ErrorResponse 失败时回包
func ErrorResponse(ctx iris.Context, errDefine constant.Error, errData interface{}) {

	resp := iris.Map{
		"code":    errDefine.Code,
		"message": errDefine.Msg,
	}

	var errMsg string
	switch data := errData.(type) {
	case error:
		errMsg = data.Error()
	case string:
		errMsg = data
	default:
		errMsg = ""
	}

	resp["data"] = errMsg

	i, err := ctx.JSON(resp)
	if err != nil {
		flog.Errorf("ErrorResponse ctx.JSON err:%s, i:%d, errDefine:%+v, errData:%+v", err, i, errDefine, errData)
	}
}

// SuccessResponse 成功时回包
func SuccessResponse(ctx iris.Context, data interface{}) {

	if data == nil {
		data = map[string]interface{}{}
	}

	i, err := ctx.JSON(iris.Map{
		"code":    0,
		"message": "",
		"data":    data,
	})
	if err != nil {
		flog.Errorf("SuccessResponse ctx.JSON err:%s, i:%d, data:%+v", err, i, data)
	}
}

func SuccessView(ctx iris.Context, filename string, data interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	err := ctx.View(filename, data)
	if err != nil {
		flog.Errorf("SuccessResponse ctx.JSON err:%s,  data:%+v", err, data)
	}
}
