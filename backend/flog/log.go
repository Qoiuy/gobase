package flog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kataras/iris/v12/context"
	"github.com/mattn/go-colorable"
	"github.com/valyala/fasttemplate"
)

type (
	Logger struct {
		prefix     string
		traceId    string
		level      Lvl
		skip       int
		output     io.Writer
		template   *fasttemplate.Template
		levels     []string
		bufferPool sync.Pool
		mutex      sync.Mutex
	}

	Lvl uint8

	JSON map[string]interface{}
)

const (
	DEBUG Lvl = iota + 1
	INFO
	WARN
	ERROR
	OFF
	panicLevel
	fatalLevel
)

var (
	ctx       context.Context
	globalLog = NewLog("-")
	adsImport = NewLog("adsImport")
	//defaultHeader = `{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}",` +
	//	`"file":"${long_file}","line":"${line}","traceId":"${traceId}"}` // kafka使用
	//defaultHeader = `{"time":"${time_rfc3339_nano}","level":"${level}",` +
	//	`"file":"${long_file}","line":"${line}"}`
	defaultHeader = `time:${time_rfc3339_nano}, level:${level}, file:${long_file}, line:${line}`
)

func SetContext(context context.Context) {
	ctx = context
}

func init() {
	globalLog.skip = 3
}

func NewLog(prefix string) (l *Logger) {
	l = &Logger{
		level:  INFO,
		skip:   2,
		prefix: prefix,
		bufferPool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 256))
			},
		},
		template: fasttemplate.New(defaultHeader, "${", "}"),
	}
	l.initLevels()
	l.SetOutput(colorable.NewColorableStdout())
	return
}

func (l *Logger) initLevels() {
	l.levels = []string{
		"-",
		"30",
		"31",
		"32",
		"33",
		"34",
		"35",
		"36",
		"37",
		"90",
	}

}

func (l *Logger) DisableColor() {
	l.initLevels()
}

func (l *Logger) EnableColor() {
	l.initLevels()
}

func (l *Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) SetPrefix(p string) {
	l.prefix = p
}

func (l *Logger) Level() Lvl {
	return l.level
}

func (l *Logger) SetLevel(v Lvl) {
	l.level = v
}

func (l *Logger) Output() io.Writer {
	return l.output
}

func (l *Logger) SetOutput(w io.Writer) {
	l.output = w
}

func (l *Logger) Print(i ...interface{}) {
	l.log(0, "", i...)
	// fmt.Fprintln(l.output, i...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.log(0, format, args...)
}

func (l *Logger) Printj(j JSON) {
	l.log(0, "json", j)
}

func (l *Logger) Debug(i ...interface{}) {
	l.log(DEBUG, "", i...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

func (l *Logger) Debugj(j JSON) {
	l.log(DEBUG, "json", j)
}

func (l *Logger) Info(i ...interface{}) {
	l.log(INFO, "", i...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

func (l *Logger) Infoj(j JSON) {
	l.log(INFO, "json", j)
}

func (l *Logger) Warn(i ...interface{}) {
	l.log(WARN, "", i...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

func (l *Logger) Warnj(j JSON) {
	l.log(WARN, "json", j)
}

func (l *Logger) Error(i ...interface{}) {
	l.log(ERROR, "", i...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

func (l *Logger) Errorj(j JSON) {
	l.log(ERROR, "json", j)
}

func (l *Logger) Fatal(i ...interface{}) {
	l.log(fatalLevel, "", i...)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log(fatalLevel, format, args...)
	os.Exit(1)
}

func (l *Logger) Fatalj(j JSON) {
	l.log(fatalLevel, "json", j)
	os.Exit(1)
}

func (l *Logger) Panic(i ...interface{}) {
	l.log(panicLevel, "", i...)
	panic(any(fmt.Sprint(i...)))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.log(panicLevel, format, args...)
	panic(any(fmt.Sprintf(format, args)))
}

func (l *Logger) Panicj(j JSON) {
	l.log(panicLevel, "json", j)
	panic(any(j))
}

func DisableColor() {
	globalLog.DisableColor()
}

func EnableColor() {
	globalLog.EnableColor()
}

func Prefix() string {
	return globalLog.Prefix()
}

func SetPrefix(p string) {
	globalLog.SetPrefix(p)
}

func Level() Lvl {
	return globalLog.Level()
}

func SetLevel(v Lvl) {
	globalLog.SetLevel(v)
}

func Output() io.Writer {
	return globalLog.Output()
}

func SetOutput(w io.Writer) {
	globalLog.SetOutput(w)
}

func Print(i ...interface{}) {
	globalLog.Print(i...)
}

func Printf(format string, args ...interface{}) {
	globalLog.Printf(format, args...)
}

func Printj(j JSON) {
	globalLog.Printj(j)
}

func Debug(i ...interface{}) {
	globalLog.Debug(i...)
}

func Debugf(format string, args ...interface{}) {
	globalLog.Debugf(format, args...)
}

func Debugj(j JSON) {
	globalLog.Debugj(j)
}

func Info(i ...interface{}) {
	globalLog.Info(i...)
}

func Infof(format string, args ...interface{}) {
	globalLog.Infof(format, args...)
}

func Infoj(j JSON) {
	globalLog.Infoj(j)
}

func Warn(i ...interface{}) {
	globalLog.Warn(i...)
}

func Warnf(format string, args ...interface{}) {
	globalLog.Warnf(format, args...)
}

func Warnj(j JSON) {
	globalLog.Warnj(j)
}

func Error(i ...interface{}) {
	globalLog.Error(i...)
}

func Errorf(format string, args ...interface{}) {
	globalLog.Errorf(format, args...)
}

func Errorj(j JSON) {
	globalLog.Errorj(j)
}

func Fatal(i ...interface{}) {
	globalLog.Fatal(i...)
}

func Fatalf(format string, args ...interface{}) {
	globalLog.Fatalf(format, args...)
}

func Fatalj(j JSON) {
	globalLog.Fatalj(j)
}

func Panic(i ...interface{}) {
	globalLog.Panic(i...)
}

func Panicf(format string, args ...interface{}) {
	globalLog.Panicf(format, args...)
}

func Panicj(j JSON) {
	globalLog.Panicj(j)
}

func (l *Logger) log(v Lvl, format string, args ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	buf := l.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer l.bufferPool.Put(buf)
	_, file, line, _ := runtime.Caller(l.skip)

	if v >= l.level || v == 0 {
		message := ""
		if format == "" {
			message = fmt.Sprint(args...)
		} else if format == "json" {
			b, err := json.Marshal(args[0])
			if err != nil {
				panic(any(err))
			}
			message = string(b)
		} else {
			message = fmt.Sprintf(format, args...)
		}

		_, err := l.template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
			switch tag {
			case "time_rfc3339":
				return w.Write([]byte(time.Now().Format(time.RFC3339)))
			case "time_rfc3339_nano":
				return w.Write([]byte(time.Now().Format(time.RFC3339Nano)))
			case "level":
				return w.Write([]byte(l.levels[v]))
			case "prefix":
				return w.Write([]byte(l.prefix))
			case "long_file":
				return w.Write([]byte(file))
			case "short_file":
				return w.Write([]byte(path.Base(file)))
			case "line":
				return w.Write([]byte(strconv.Itoa(line)))
			}
			return 0, nil
		})

		if err == nil {
			s := buf.String()
			i := buf.Len() - 1
			if s[i] == '}' {
				// JSON header
				buf.Truncate(i)
				buf.WriteByte(',')
				if format == "json" {
					buf.WriteString(message[1:])
				} else {
					buf.WriteString(` message:`)
					buf.WriteString(strconv.Quote(message))
					//buf.WriteString(fmtMsg(message))
					buf.WriteString(` `)
				}
			} else {
				// Text header
				buf.WriteByte(' ')
				buf.WriteString(message)
			}
			buf.WriteByte('\n')
			l.output.Write(buf.Bytes())
		}
	}
}

func fmtMsg(msg string) string {
	var length int
	arr := strings.Split(msg, "\n")
	if len(arr) > 20 {
		length = 20
	} else {
		length = len(arr)
	}
	return strings.Join(arr[:length], "\n")
}

func IInfof(format string, args ...interface{}) {
	adsImport.Infof(format, args...)
}

func IWarnf(format string, args ...interface{}) {
	adsImport.Warnf(format, args...)
}

func IErrorf(format string, args ...interface{}) {
	adsImport.Errorf(format, args...)
}
