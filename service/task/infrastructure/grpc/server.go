package grpc

import (
	"context"
	"sync"
	"time"

	cli "github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/server"
	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/pkg/constant"
	"github.com/kzmake/micro-kit/pkg/logger/technical"
	logWrapper "github.com/kzmake/micro-kit/pkg/wrapper/logger"

	"github.com/kzmake/micro-kit/service/task/config"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
)

var (
	service = constant.Service.Task
	version = "v0.1.0"

	waitGroup = new(sync.WaitGroup)
)

func waitgroup(waitGroup *sync.WaitGroup) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			waitGroup.Add(1)
			defer waitGroup.Done()
			return h(ctx, req, rsp)
		}
	}
}

// Server はサーバーとして動作するアプリケーションです。
type Server interface {
	Run() error
}

// New はサーバーを生成します。
func New(conf *config.Config, handler proto.TaskServiceHandler) Server {
	service := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Address(conf.Endpoint),

		micro.RegisterTTL(30*time.Second),      // nolint:gomnd
		micro.RegisterInterval(10*time.Second), // nolint:gomnd
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(conf.ServiceDiscovery.Endpoint),
		)),

		micro.WrapHandler(
			waitgroup(waitGroup),
			logWrapper.NewHandlerWrapper(),
		),
		micro.WrapSubscriber(
			logWrapper.NewSubscriberWrapper(),
		),
		micro.WrapClient(
			logWrapper.NewClientWrapper(),
		),

		micro.BeforeStop(func() error {
			waitGroup.Wait()
			return nil
		}),
	)
	service.Init(
		micro.Action(func(c *cli.Context) error {
			return nil
		}),
	)

	// Register Handler
	if err := proto.RegisterTaskServiceHandler(service.Server(), handler); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("handler の登録に失敗しました: %w", err))
	}

	return service
}
