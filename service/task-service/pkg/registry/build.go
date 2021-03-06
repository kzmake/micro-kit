package registry

import (
	"github.com/jinzhu/gorm"
	di "github.com/sarulabs/di/v2"

	log "github.com/kzmake/micro-kit/pkg/logger"

	"github.com/kzmake/micro-kit/service/task-service/domain/repository"
	"github.com/kzmake/micro-kit/service/task-service/infrastructure/mysql"
	"github.com/kzmake/micro-kit/service/task-service/infrastructure/ulid"
	"github.com/kzmake/micro-kit/service/task-service/interface/controller"
	"github.com/kzmake/micro-kit/service/task-service/usecase/business"
	"github.com/kzmake/micro-kit/service/task-service/usecase/business/logger"
	"github.com/kzmake/micro-kit/service/task-service/usecase/interactor"
	"github.com/kzmake/micro-kit/service/task-service/usecase/port"
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

func buildListTasksInteractor(businessManagerKey, taskRepositoryKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		businessManager := ctn.Get(businessManagerKey).(business.Manager)
		taskRepository := ctn.Get(taskRepositoryKey).(repository.Task)
		return interactor.NewListTasks(businessManager, taskRepository), nil
	}
}

func buildGetTaskInteractor(businessManagerKey, taskRepositoryKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		businessManager := ctn.Get(businessManagerKey).(business.Manager)
		taskRepository := ctn.Get(taskRepositoryKey).(repository.Task)
		return interactor.NewGetTask(businessManager, taskRepository), nil
	}
}

func buildDeleteTaskInteractor(businessManagerKey, taskRepositoryKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		businessManager := ctn.Get(businessManagerKey).(business.Manager)
		taskRepository := ctn.Get(taskRepositoryKey).(repository.Task)
		return interactor.NewDeleteTask(businessManager, taskRepository), nil
	}
}

func buildTaskQueryController(listTaskInputPortKey, getTaskInputPortKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		listTaskInputPort := ctn.Get(listTaskInputPortKey).(port.ListTasks)
		getTaskInputPort := ctn.Get(getTaskInputPortKey).(port.GetTask)
		return controller.NewTaskQuery(listTaskInputPort, getTaskInputPort), nil
	}
}

func buildTaskCommandController(createTaskInputPortKey, deleteTaskInputPortKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		createTaskInputPort := ctn.Get(createTaskInputPortKey).(port.CreateTask)
		deleteTaskInputPort := ctn.Get(deleteTaskInputPortKey).(port.DeleteTask)
		return controller.NewTaskCommand(createTaskInputPort, deleteTaskInputPort), nil
	}
}

func buildTaskController(taskQueryControllerKey, taskCommandControllerKey string) func(ctn di.Container) (interface{}, error) {
	return func(ctn di.Container) (interface{}, error) {
		taskQueryController := ctn.Get(taskQueryControllerKey).(*controller.TaskQuery)
		taskCommandController := ctn.Get(taskCommandControllerKey).(*controller.TaskCommand)
		return controller.NewTask(taskQueryController, taskCommandController), nil
	}
}
