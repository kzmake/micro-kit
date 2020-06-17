package registry

import (
	"github.com/jinzhu/gorm"
	di "github.com/sarulabs/di/v2"

	log "github.com/kzmake/micro-kit/pkg/logger"

	"github.com/kzmake/micro-kit/service/task/domain/repository"
	"github.com/kzmake/micro-kit/service/task/infrastructure/mysql"
	"github.com/kzmake/micro-kit/service/task/infrastructure/ulid"
	"github.com/kzmake/micro-kit/service/task/interface/controller"
	"github.com/kzmake/micro-kit/service/task/usecase/business"
	"github.com/kzmake/micro-kit/service/task/usecase/business/logger"
	"github.com/kzmake/micro-kit/service/task/usecase/interactor"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

func buildULIDIDRepository() func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		return ulid.NewIDRepository(), nil
	}
}

func buildLoggerAssistant(loggerKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		lgr := ctn.Get(loggerKey).(*log.Logger)
		return logger.New(lgr), nil
	}
}

func buildBusinessManager(loggerAssistantKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		loggerAssistant := ctn.Get(loggerAssistantKey).(business.Assistant)
		return business.New(loggerAssistant), nil
	}
}

func buildMysqlTaskRepository(databaseKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		db := ctn.Get(databaseKey).(*gorm.DB)
		return mysql.NewTaskRepository(db), nil
	}
}

func buildCreateTaskInteractor(businessManagerKey, idRepositoryKey, taskRepositoryKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		businessManager := ctn.Get(businessManagerKey).(business.Manager)
		idRepository := ctn.Get(idRepositoryKey).(repository.ID)
		taskRepository := ctn.Get(taskRepositoryKey).(repository.Task)
		return interactor.NewCreateTask(businessManager, idRepository, taskRepository), nil
	}
}

func buildGetTaskInteractor(businessManagerKey, taskRepositoryKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		businessManager := ctn.Get(businessManagerKey).(business.Manager)
		taskRepository := ctn.Get(taskRepositoryKey).(repository.Task)
		return interactor.NewGetTask(businessManager, taskRepository), nil
	}
}

func buildTaskQueryController(getTaskInputPortKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		getTaskInputPort := ctn.Get(getTaskInputPortKey).(port.GetTask)
		return controller.NewTaskQuery(getTaskInputPort), nil
	}
}

func buildTaskCommandController(createTaskInputPortKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		createTaskInputPort := ctn.Get(createTaskInputPortKey).(port.CreateTask)
		return controller.NewTaskCommand(createTaskInputPort), nil
	}
}

func buildTaskController(taskQueryControllerKey, taskCommandControllerKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		taskQueryController := ctn.Get(taskQueryControllerKey).(*controller.TaskQuery)
		taskCommandController := ctn.Get(taskCommandControllerKey).(*controller.TaskCommand)
		return controller.NewTask(taskQueryController, taskCommandController), nil
	}
}
