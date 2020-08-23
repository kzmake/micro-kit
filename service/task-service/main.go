package main

import (
	"os"

	"golang.org/x/xerrors"

	cli "github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"

	"github.com/kzmake/micro-kit/pkg/constant"
	"github.com/kzmake/micro-kit/pkg/logger/technical"

	"github.com/kzmake/micro-kit/service/task-service/infrastructure/grpc"
	"github.com/kzmake/micro-kit/service/task-service/interface/proto"
	"github.com/kzmake/micro-kit/service/task-service/pkg/config"
	"github.com/kzmake/micro-kit/service/task-service/pkg/registry"
)

var (
	name    = constant.TaskService
	version = "v0.1.0"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		technical.Errorf("%+v", xerrors.Errorf("configの生成に失敗しました: %w", err))
		os.Exit(1)
	}

	s, err := grpc.New(
		micro.Name(name),
		micro.Version(version),
	)
	if err != nil {
		technical.Errorf("%+v", xerrors.Errorf("serviceの生成に失敗しました: %w", err))
		os.Exit(1)
	}

	s.Init(
		micro.Action(func(c *cli.Context) error { return nil }),
	)

	ctn, err := registry.New(cfg)
	if err != nil {
		technical.Errorf("%+v", xerrors.Errorf("DIコンテナの生成に失敗しました: %w", err))
		os.Exit(1)
	}

	h := ctn.Get("controller").(proto.TaskServiceHandler)
	if err := proto.RegisterTaskServiceHandler(s.Server(), h); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("handlerの取得に失敗しました: %w", err))
		os.Exit(1)
	}

	if err := s.Run(); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("serviceの起動に失敗しました: %w", err))
		os.Exit(1)
	}
}
