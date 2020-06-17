package main

import (
	"os"
	"time"

	"golang.org/x/xerrors"

	"github.com/jinzhu/gorm"
	di "github.com/sarulabs/di/v2"

	"github.com/kzmake/micro-kit/pkg/config"
	"github.com/kzmake/micro-kit/pkg/logger"
	"github.com/kzmake/micro-kit/pkg/logger/technical"
	"github.com/kzmake/micro-kit/pkg/tracer"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
	conf "github.com/kzmake/micro-kit/service/task/pkg/config"
	"github.com/kzmake/micro-kit/service/task/pkg/registry"

	"github.com/kzmake/micro-kit/service/task/infrastructure/grpc"
)

var app = []di.Def{
	{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			l := logger.New(
				logger.WithOutput(os.Stdout),
				logger.WithTimeFormat(time.RFC3339Nano),
				logger.WithSkipFrameCount(4), // nolint:gomnd
			)
			technical.Logger = l
			return l, nil
		},
	},
	{
		Name:  "tracer",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) { return tracer.New("task") },
	},
	{
		Name:  "database",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			c := ctn.Get("config").(*conf.Config)
			return gorm.Open(c.Database.Driver, c.Database.DSN)
		},
		Close: func(obj interface{}) error {
			return obj.(*gorm.DB).Close()
		},
	},
}

func main() {
	cfg, err := config.New("TASK", &conf.Config{})
	if err != nil {
		technical.Errorf("%+v", xerrors.Errorf("configの取得に失敗しました: %w", err))
		os.Exit(1)
	}

	s := grpc.New(cfg.(*conf.Config))

	c := []di.Def{{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) { return cfg, nil },
	}}

	ctn, err := registry.New(append(append(c, app...), registry.Task...)...)
	if err != nil {
		technical.Errorf("%+v", xerrors.Errorf("diコンテナの生成に失敗しました: %w", err))
		os.Exit(1)
	}

	h := ctn.Get("taskController").(proto.TaskServiceHandler)
	if err := proto.RegisterTaskServiceHandler(s.Server(), h); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("handlerの登録に失敗しました: %w", err))
		os.Exit(1)
	}

	if err := s.Run(); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("serverの起動に失敗しました: %w", err))
		os.Exit(1)
	}
}
