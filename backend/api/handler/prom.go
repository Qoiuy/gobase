package handler

import (
	"context"
	"github.com/Qoiuy/go-base/backend/api/config"
	irisContext "github.com/kataras/iris/v12/context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
	"time"
)

func NewPrometheus(ctx context.Context) error {
	Env := "dev"
	if !config.Opts.Debug {
		Env = "prod"
	}

	InitPrometheus("go-base-api", Env)

	server := http.Server{Addr: "0.0.0.0:8399"}
	go func() {
		select {
		case <-ctx.Done():
			_ = server.Shutdown(ctx)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	return server.ListenAndServe()
}

// HttpServerCounterVec 接口访问量统计
var HttpServerCounterVec *prometheus.CounterVec

// HttpServerTimerVec 接口相应时间统计
var HttpServerTimerVec *prometheus.HistogramVec

// InitPrometheus 初始化prometheus（main函数调用此函数）
func InitPrometheus(ServiceName string, Env string) {
	constLabels := map[string]string{"service": ServiceName, "env": Env, "host": "0.0.0.0"}

	HttpServerTimerVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        "http_request_duration_seconds",
			Help:        "How long it took to process the HTTP request, partitioned by status code, method and HTTP path.",
			ConstLabels: constLabels,
			Buckets:     []float64{0.3, 1.2, 5.0},
		},
		[]string{"code", "method", "path"},
	)

	HttpServerCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "http_requests_total",
			Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
			ConstLabels: constLabels,
		},
		[]string{"code", "method", "path"},
	)

	prometheus.MustRegister(HttpServerTimerVec, HttpServerCounterVec)
}

// iris中间件metrics统计
func HttpInterceptor(ctx irisContext.Context) {

	start := time.Now()
	ctx.Next()
	r := ctx.Request()
	statusCode := strconv.Itoa(ctx.GetStatusCode())
	duration := float64(time.Since(start).Nanoseconds()) / 1000000000
	labels := []string{statusCode, r.Method, r.URL.Path}

	HttpServerCounterVec.WithLabelValues(labels...).Inc()
	HttpServerTimerVec.WithLabelValues(labels...).Observe(duration)

}
