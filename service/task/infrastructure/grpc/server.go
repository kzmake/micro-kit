package grpc

import (
	cli "github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
	registry "github.com/micro/go-plugins/registry/etcdv3/v2"
	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/pkg/constant"
	"github.com/kzmake/micro-kit/pkg/logger"
	logWrapper "github.com/kzmake/micro-kit/pkg/wrapper/log"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
)

var (
	service = constant.Service.Task
	version = "v0.1.0"
)

// Server はサーバーとして動作するアプリケーションです。
type Server interface {
	Run() error
}

// New はサーバーを生成します。
func New(taskController proto.TaskHandler) Server {
	service := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	service.Init(
		// init
		micro.Action(func(c *cli.Context) error {
			// load config
			return nil
		}),

		// logging
		micro.WrapHandler(logWrapper.NewHandlerWrapper()),
		micro.WrapSubscriber(logWrapper.NewSubscriberWrapper()),
		micro.WrapClient(logWrapper.NewClientWrapper()),

		// service registry
		micro.Registry(registry.NewRegistry()),
	)

	s := service.Server()
	s.Init(
		server.Wait(nil),
	)

	// Register Handler
	if err := proto.RegisterTaskHandler(s, taskController); err != nil {
		logger.Errorf("%+v", xerrors.Errorf("handler の登録に失敗しました: %w", err))
	}

	return service
}
