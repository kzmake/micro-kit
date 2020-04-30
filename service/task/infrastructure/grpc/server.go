package grpc

import (
	"fmt"

	cli "github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"

	"github.com/kzmake/micro-kit/service/task/interface/proto"

	logWrapper "github.com/kzmake/micro-kit/pkg/wrapper/log"
)

// Server はサーバーとして動作するアプリケーションです。
type Server interface {
	Run() error
}

// New はサーバーを生成します。
func New(
	serviceName string, version string,
	taskController proto.TaskServiceHandler,
) Server {
	s := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
	)

	// Initialize service
	options := []micro.Option{
		// initialize
		micro.Action(func(c *cli.Context) error {
			// load config
			return nil
		}),

		// logging
		micro.WrapHandler(logWrapper.NewHandlerWrapper()),
		micro.WrapSubscriber(logWrapper.NewSubscriberWrapper()),
		micro.WrapClient(logWrapper.NewClientWrapper()),
	}

	// Initialize
	s.Init(options...)

	// Register Handler
	if err := proto.RegisterTaskServiceHandler(s.Server(), taskController); err != nil {
		fmt.Println("Error: Handlerの登録に失敗しました")
	}

	return s
}
