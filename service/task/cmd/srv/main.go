package main

import (
	"os"

	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/pkg/logger/technical"

	"github.com/kzmake/micro-kit/service/task/infrastructure/grpc"
	"github.com/kzmake/micro-kit/service/task/interface/proto"
	"github.com/kzmake/micro-kit/service/task/registry"
)

func main() {
	ctn, err := registry.New(registry.Production...)
	if err != nil {
		technical.Errorf("%+v", xerrors.Errorf("serverの起動に失敗しました: %w", err))
		os.Exit(1)
	}

	s := grpc.New(ctn.Get("taskController").(proto.TaskServiceHandler))

	if err := s.Run(); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("serverの起動に失敗しました: %w", err))
		os.Exit(1)
	}
}
