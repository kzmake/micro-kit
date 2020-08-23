package grpc

import (
	"context"
	"sync"
	"time"

	micro "github.com/micro/go-micro/v2"
	mserver "github.com/micro/go-micro/v2/server"

	plogger "github.com/kzmake/micro-kit/pkg/wrapper/logger"
)

func waitgroup(waitGroup *sync.WaitGroup) mserver.HandlerWrapper {
	return func(h mserver.HandlerFunc) mserver.HandlerFunc {
		return func(ctx context.Context, req mserver.Request, rsp interface{}) error {
			waitGroup.Add(1)
			defer waitGroup.Done()
			return h(ctx, req, rsp)
		}
	}
}

// New は gRPC Service を生成します。
func New(opts ...micro.Option) (micro.Service, error) {
	wg := new(sync.WaitGroup)
	o := append(opts, []micro.Option{
		micro.RegisterTTL(30 * time.Second),      // nolint:gomnd
		micro.RegisterInterval(10 * time.Second), // nolint:gomnd

		micro.WrapHandler(
			waitgroup(wg),
			plogger.NewHandlerWrapper(),
		),

		micro.BeforeStop(func() error {
			wg.Wait()
			return nil
		}),
	}...)

	return micro.NewService(o...), nil
}
