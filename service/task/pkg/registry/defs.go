package registry

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	di "github.com/sarulabs/di/v2"

	"github.com/kzmake/micro-kit/pkg/logger"
	"github.com/kzmake/micro-kit/pkg/logger/technical"
	"github.com/kzmake/micro-kit/pkg/tracer"

	"github.com/kzmake/micro-kit/service/task/pkg/config"
)

var outputs = []di.Def{
	{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			l := logger.New(
				logger.WithOutput(os.Stdout),
				logger.WithTimeFormat(time.RFC3339Nano),
				logger.WithSkipFrameCount(4), // nolint:gomnd
			)
			technical.Logger = l
			return l, nil
		},
	},
	{
		Name:  "tracer",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) { return tracer.New("task") },
	},
	{
		Name:  "database",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			c := ctn.Get("config").(*config.Config)
			return gorm.Open(c.Database.Driver, c.Database.DSN)
		},
		Close: func(obj interface{}) error {
			return obj.(*gorm.DB).Close()
		},
	},
}

var cores = []di.Def{
	{
		Name:  "idRepository",
		Scope: di.App,
		Build: buildULIDIDRepository(),
	},
	{
		Name:  "taskRepository",
		Scope: di.App,
		Build: buildMysqlTaskRepository("database"),
	},
	{
		Name:  "loggerAssistant",
		Scope: di.App,
		Build: buildLoggerAssistant("logger"),
	},
	{
		Name:  "businessManager",
		Scope: di.App,
		Build: buildBusinessManager("loggerAssistant"),
	},
	{
		Name:  "createTaskInputPort",
		Scope: di.App,
		Build: buildCreateTaskInteractor("businessManager", "idRepository", "taskRepository"),
	},
	{
		Name:  "listTasksInputPort",
		Scope: di.App,
		Build: buildListTasksInteractor("businessManager", "taskRepository"),
	},
	{
		Name:  "getTaskInputPort",
		Scope: di.App,
		Build: buildGetTaskInteractor("businessManager", "taskRepository"),
	},
	{
		Name:  "deleteTaskInputPort",
		Scope: di.App,
		Build: buildDeleteTaskInteractor("businessManager", "taskRepository"),
	},
	{
		Name:  "taskQueryController",
		Scope: di.App,
		Build: buildTaskQueryController("listTasksInputPort", "getTaskInputPort"),
	},
	{
		Name:  "taskCommandController",
		Scope: di.App,
		Build: buildTaskCommandController("createTaskInputPort", "deleteTaskInputPort"),
	},
	{
		Name:  "taskController",
		Scope: di.App,
		Build: buildTaskController("taskQueryController", "taskCommandController"),
	},
}
