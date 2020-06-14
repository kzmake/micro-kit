package main

import (
	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/pkg/logger/technical"

	"github.com/kzmake/micro-kit/service/task/infrastructure/api"
)

func main() {
	s := api.New()

	if err := s.Run(); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("server の起動に失敗しました: %w", err))
	}
}
