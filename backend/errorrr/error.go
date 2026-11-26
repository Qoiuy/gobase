package errorrr

import (
	"fmt"
	"github.com/Qoiuy/go-base/backend/constant"
	"github.com/Qoiuy/go-base/backend/flog"
	"github.com/pkg/errors"
	"strings"
)

func fmtError(err error) string {
	var length int
	arr := strings.Split(fmt.Sprintf("%+v", err), "\n")
	if len(arr) > 15 {
		length = 15
	} else {
		length = len(arr)
	}
	return strings.Join(arr[:length], "\n")
}
func HandleError(err error) constant.Err {
	// 获取error中的string
	final := fmtError(err)

	// 获取根因
	root := errors.Cause(err)
	custom, ok := root.(constant.Err)
	if !ok {
		custom = constant.ErrUnknown
	}

	// 处理日志
	switch custom.Level {
	case constant.Debug:
		flog.Debug(final)
	case constant.Info:
		flog.Info(final)
	case constant.Warn:
		flog.Warn(final)
	case constant.EError:
		flog.Error(final)
	case constant.Fatal:
		flog.Fatal(final)
	}
	return custom
}
