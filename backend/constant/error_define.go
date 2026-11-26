package constant

// Error 定义response结构体	返回错误信息
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// FixError 定义response结构体	返回错误信息
type FixError struct {
	Error
	FixId string `json:"fixId"`
}

var ErrNil = Error{0, "没有这个错误"}
var ErrParamFailed = Error{-4001, "参数错误相关"}
var ErrMysqlDBFailed = Error{-2001, "数据库出现问题"}
var ErrRedisFailed = Error{-2002, "数据库出现问题"}
var ErrGolangThirdLibFailed = Error{-3001, "Golang 三方库出现问题"}
var ErrGoLangLibFailed = Error{-3002, "Golang库调用时出现问题"}
var ErrPdfServiceFailed = Error{-8001, "PDF服务(业务代码)出现问题"}
var ErrChangeSelectFailed = Error{-8002, "错误点击后果自负"} // 前端显示的参数错误-可以增加参数
var ErrPythonShellFailed = Error{-8003, "Python 脚本出现问题"}
