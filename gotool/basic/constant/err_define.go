package constant

var ErrBasic = Err{Code: 001, Msg: "gotool.basic 包执行异常", Level: Error}
var Err7z = Err{Code: 010, Msg: "gotool.7z 包执行异常", Level: Error}
var ErrFlow = Err{Code: 011, Msg: "gotool.flow 包执行异常", Level: Error}
var ErrLrv = Err{Code: 011, Msg: "gotool.lrv 包执行异常", Level: Error}
var ErrWeb = Err{Code: 012, Msg: "gotool.web 包执行异常", Level: Error}

type ErrorLevel = string

const (
	Debug ErrorLevel = "Debug"
	Info  ErrorLevel = "Info"
	Warn  ErrorLevel = "Warn"
	Error ErrorLevel = "Error"
	Fatal ErrorLevel = "Fatal"
)

type Err struct {
	Code  int
	Msg   string
	Level string
}

func (e Err) Error() string {
	return e.Msg
}

func (e *Err) SetErrorMsg(msg string) {
	e.Msg = msg
}

func (e *Err) SetErrorLevel(level ErrorLevel) {
	e.Level = level
}
