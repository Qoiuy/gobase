package main

import (
	"context"
	"fmt"
	"github.com/Qoiuy/go-base/backend/api/config"
	"github.com/Qoiuy/go-base/backend/api/handler"
	"github.com/Qoiuy/go-base/backend/constant"
	"github.com/Qoiuy/go-base/backend/errorrr"
	"github.com/Qoiuy/go-base/backend/flog"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.Setup()
	loggerConfig := flog.LoggerConfig{
		LoggerAddress:   config.Opts.LoggerAddress,
		LoggerNamespace: config.Opts.LoggerNamespace,
		LoggerPassword:  config.Opts.LoggerPassword,
		LoggerPrefix:    config.Opts.LoggerPrefix,
	}
	flog.InitLogger(loggerConfig, config.Opts.Debug)
}
func main() {

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			close(signalChannel)
			return nil
		case <-signalChannel:
			return errors.WithStack(constant.ErrSignalDone)
		}
	})
	group.Go(func() error {
		return errors.Wrap(constant.ErrServerDown, fmt.Sprintf("iris app发生错误, error : %s", handler.NewApp(ctx)))
	})
	group.Go(func() error {
		return errors.Wrap(constant.ErrServerDown, fmt.Sprintf("prometheus服务发生错误, error : %s", handler.NewPrometheus(ctx)))
	})

	if err := group.Wait(); err != nil {
		_ = errorrr.HandleError(err)
	}

	time.Sleep(5 * time.Second)

}
