package grpc

import (
	"context"
	"sync"
	"time"

	cli "github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	mregistry "github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	mserver "github.com/micro/go-micro/v2/server"

	pconfig "github.com/kzmake/micro-kit/pkg/config"
	pconstant "github.com/kzmake/micro-kit/pkg/constant"
	plogger "github.com/kzmake/micro-kit/pkg/wrapper/logger"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
	"github.com/kzmake/micro-kit/service/task/pkg/config"
	"github.com/kzmake/micro-kit/service/task/pkg/registry"
)

var (
	name    = pconstant.TaskService
	version = "v0.1.0"
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

// New はサーバーを生成します。
func New() (micro.Service, error) {
	wg := new(sync.WaitGroup)
	service := micro.NewService(
		micro.Name(name),
		micro.Version(version),

		micro.WrapHandler(
			waitgroup(wg),
			plogger.NewHandlerWrapper(),
		),
		micro.WrapSubscriber(
			plogger.NewSubscriberWrapper(),
		),
		micro.WrapClient(
			plogger.NewClientWrapper(),
		),

		micro.BeforeStop(func() error {
			wg.Wait()
			return nil
		}),
	)

	c, err := pconfig.New("TASK", &config.Config{})
	if err != nil {
		return nil, err
	}
	cfg := c.(*config.Config)

	service.Init(
		micro.Action(func(c *cli.Context) error {
			return nil
		}),

		micro.Address(cfg.Endpoint),

		micro.RegisterTTL(30*time.Second),      // nolint:gomnd
		micro.RegisterInterval(10*time.Second), // nolint:gomnd
		micro.Registry(etcd.NewRegistry(
			mregistry.Addrs(cfg.ServiceDiscovery.Endpoint),
		)),
	)

	ctn, err := registry.New(cfg)
	if err != nil {
		return nil, err
	}

	h := ctn.Get("taskController").(proto.TaskServiceHandler)
	if err := proto.RegisterTaskServiceHandler(service.Server(), h); err != nil {
		return nil, err
	}

	return service, nil
}
