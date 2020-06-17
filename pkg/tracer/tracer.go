package tracer

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// New は tracer を生成します。
func New(serviceName string) (opentracing.Tracer, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	cfg.ServiceName = serviceName

	tracer, _, err := cfg.NewTracer(
		config.Logger(log.NullLogger),
		config.Metrics(metrics.NullFactory),
	)
	if err != nil {
		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)

	return tracer, nil
}
