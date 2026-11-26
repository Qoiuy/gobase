package constant

var ErrServerDown = Err{10, "服务进程异常", EError}
var ErrSignalDone = Err{11, "系统关闭服务进程", Warn}
var ErrUnknown = Err{11, "未知异常", Warn}

type ErrorLevel = string

const (
	Debug  ErrorLevel = "Debug"
	Info   ErrorLevel = "Info"
	Warn   ErrorLevel = "Warn"
	EError ErrorLevel = "Error"
	Fatal  ErrorLevel = "Fatal"
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
