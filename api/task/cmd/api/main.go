package main

import (
	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/pkg/logger/technical"

	"github.com/kzmake/micro-kit/service/task/infrastructure/grpc"
	"github.com/kzmake/micro-kit/service/task/interface/controller"
	"github.com/kzmake/micro-kit/service/task/usecase/interactor"
)

var (
	createTaskInputPort = interactor.NewCreateTask(nil, nil)
	taskController      = controller.NewTask(createTaskInputPort)
)

func main() {
	s := grpc.New(taskController)

	if err := s.Run(); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("server の起動に失敗しました: %w", err))
	}
}
