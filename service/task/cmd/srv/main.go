package main

import (
	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/pkg/logger"
	"github.com/kzmake/micro-kit/service/task/infrastructure/grpc"
	"github.com/kzmake/micro-kit/service/task/interface/controller"
	"github.com/kzmake/micro-kit/service/task/interface/presenter"
	"github.com/kzmake/micro-kit/service/task/usecase/interactor"
)

var (
	createTaskOutputPort = presenter.NewCreateTaskPresenter()
	createTaskInputPort  = interactor.NewCreateTaskInteractor(createTaskOutputPort)
	taskController       = controller.NewTaskController(createTaskInputPort)
)

func main() {
	s := grpc.New(taskController)

	if err := s.Run(); err != nil {
		logger.Errorf("%+v", xerrors.Errorf("server の起動に失敗しました: %w", err))
	}
}
