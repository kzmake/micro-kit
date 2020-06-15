package main

import (
	"os"

	"golang.org/x/xerrors"

	"github.com/jinzhu/gorm"

	"github.com/kzmake/micro-kit/pkg/logger/technical"

	"github.com/kzmake/micro-kit/service/task/infrastructure/grpc"
	"github.com/kzmake/micro-kit/service/task/infrastructure/mysql"
	"github.com/kzmake/micro-kit/service/task/infrastructure/ulid"
	"github.com/kzmake/micro-kit/service/task/interface/controller"
	"github.com/kzmake/micro-kit/service/task/usecase/business"
	"github.com/kzmake/micro-kit/service/task/usecase/business/logger"
	"github.com/kzmake/micro-kit/service/task/usecase/interactor"
)

func main() {
	loggerAssistant := logger.New(os.Stdout)
	manager := business.New(loggerAssistant)

	db, err := gorm.Open("mysql", "test:test@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		technical.Errorf("%+v", xerrors.Errorf("DBの接続に失敗しました: %w", err))
	}

	idRepository := ulid.NewIDRepository()
	taskRepository := mysql.NewTaskRepository(db)
	createTaskPort := interactor.NewCreateTask(manager, idRepository, taskRepository)
	getTaskPort := interactor.NewGetTask(manager, taskRepository)
	taskQueryController := controller.NewTaskQuery(getTaskPort)
	taskCommandController := controller.NewTaskCommand(createTaskPort)
	requestController := controller.NewTask(taskQueryController, taskCommandController)

	s := grpc.New(requestController)

	if err := s.Run(); err != nil {
		technical.Errorf("%+v", xerrors.Errorf("serverの起動に失敗しました: %w", err))
	}
}
