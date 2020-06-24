package main

import (
	"os"

	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/pkg/logger/technical"

	"github.com/kzmake/micro-kit/service/task/infrastructure/grpc"
)

func main() {
	service, err := grpc.New()
	if err != nil {
		technical.Errorf("%+v", xerrors.Errorf("serviceの生成に失敗しました: %w", err))
		os.Exit(1)
	}

	if err := service.Run(); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("serviceの起動に失敗しました: %w", err))
		os.Exit(1)
	}
}
