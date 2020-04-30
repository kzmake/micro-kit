package main

import (
	"fmt"

	// "github.com/xmlking/logger/log"

	"github.com/kzmake/micro-kit/service/task/infrastructure/grpc"
	"github.com/kzmake/micro-kit/service/task/interface/controller"
	"github.com/kzmake/micro-kit/service/task/interface/presenter"
	"github.com/kzmake/micro-kit/service/task/usecase/interactor"

	// myConfig "github.com/xmlking/micro-starter-kit/shared/config"
	constants "github.com/kzmake/micro-kit/shared/constants"
	// "github.com/xmlking/micro-starter-kit/shared/logger"
	// "github.com/xmlking/micro-starter-kit/shared/util"
	// transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
)

const (
	serviceName = constants.TASK_SERVICE
	version     = "v0.1.0"
)

var (
	createTaskOutputPort = presenter.NewCreateTaskPresenter()
	createTaskInputPort  = interactor.NewCreateTaskInteractor(createTaskOutputPort)
	taskController       = controller.NewTaskController(createTaskInputPort)
)

func main() {
	s := grpc.New(serviceName, version, taskController)

	if err := s.Run(); err != nil {
		fmt.Println("Error: サーバー起動に失敗しました")
	}
}
