# zlog

`backend/zlog` 是基于 zap 和 lumberjack 的文件日志封装。

## 当前状态

`backend/api` 没有引用本包，API 实际使用的是 `backend/flog`（输出到标准输出）。本包目前是独立实现，导入时会由 `init()` 自动初始化。

## 初始化

- 包被导入时执行 `init()`，调用 `NewLog("./logs/go-base.log")`。
- 路径相对于进程工作目录；目录名是 `logs/`（复数），不是仓库根目录的 `log/`。
- lumberjack 会自动创建 `logs/` 目录，但父目录需要可写；日志文件在第一次写入时创建。
- `NewLog(filePath)` 可以替换全局 logger，但不返回 logger 实例；目前只暴露 `Infof` 一个记录方法。

## 日志格式与切割

- 使用 `zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())`，每行一条 JSON。
- 只创建了 core，没有启用 caller 或 stacktrace。
- 级别固定为 `Info`，没有暴露修改级别的接口。
- lumberjack 参数：

  | 参数 | 值 |
  | --- | --- |
  | `MaxSize` | 15 MB |
  | `MaxBackups` | 1024 |
  | `MaxAge` | 365 × 3 天 |
  | `Compress` | `true` |

## 用法

```go
import (
	"go.uber.org/zap"

	"github.com/Qoiuy/go-base/backend/zlog"
)

func main() {
	zlog.Infof("server started", zap.String("addr", ":8080"))
}
```

`Infof` 的第二个参数起是 `zap.Field`，不是 `fmt` 的格式化参数；第一个参数会作为 message 原样传给 `zap.Logger.Log`，不处理 `%s`、`%d` 等占位符。
