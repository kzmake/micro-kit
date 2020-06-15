package main

import (
	"os"

	"golang.org/x/xerrors"

	di "github.com/sarulabs/di/v2"

	conf "github.com/kzmake/micro-kit/pkg/config"
	"github.com/kzmake/micro-kit/pkg/logger/technical"

	"github.com/kzmake/micro-kit/service/task/config"
	"github.com/kzmake/micro-kit/service/task/registry"

	"github.com/kzmake/micro-kit/service/task/infrastructure/grpc"
	"github.com/kzmake/micro-kit/service/task/interface/proto"
)

var app = []di.Def{
	{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return conf.New("TASK", &config.Config{})
		},
	},
	{
		Name:  "app",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			c := ctn.Get("config").(*config.Config)
			h := ctn.Get("taskController").(proto.TaskServiceHandler)
			return grpc.New(c, h), nil
		},
	},
}

func main() {
	ctn, err := registry.New(append(app, registry.Production...)...)
	if err != nil {
		technical.Errorf("%+v", xerrors.Errorf("DIコンテナの生成に失敗しました: %w", err))
		os.Exit(1)
	}

	s := ctn.Get("app").(grpc.Server)

	if err := s.Run(); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("serverの起動に失敗しました: %w", err))
		os.Exit(1)
	}
}
