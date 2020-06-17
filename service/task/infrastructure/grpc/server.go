package grpc

import (
	"context"
	"sync"
	"time"

	cli "github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	reg "github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/server"

	"github.com/kzmake/micro-kit/pkg/constant"
	logWrapper "github.com/kzmake/micro-kit/pkg/wrapper/logger"

	"github.com/kzmake/micro-kit/service/task/pkg/config"
)

var (
	service = constant.Service.Task
	version = "v0.1.0"
)

var wg = new(sync.WaitGroup)

func waitgroup(waitGroup *sync.WaitGroup) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			waitGroup.Add(1)
			defer waitGroup.Done()
			return h(ctx, req, rsp)
		}
	}
}

// New はサーバーを生成します。
func New(conf *config.Config) micro.Service {
	s := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Address(conf.Endpoint),

		micro.RegisterTTL(30*time.Second),      // nolint:gomnd
		micro.RegisterInterval(10*time.Second), // nolint:gomnd
		micro.Registry(etcd.NewRegistry(
			reg.Addrs(conf.ServiceDiscovery.Endpoint),
		)),

		micro.WrapHandler(
			waitgroup(wg),
			logWrapper.NewHandlerWrapper(),
		),
		micro.WrapSubscriber(
			logWrapper.NewSubscriberWrapper(),
		),
		micro.WrapClient(
			logWrapper.NewClientWrapper(),
		),

		micro.BeforeStop(func() error {
			wg.Wait()
			return nil
		}),
	)

	s.Init(
		micro.Action(func(c *cli.Context) error {
			return nil
		}),
	)

	return s
}
