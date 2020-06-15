package registry

import (
	"os"

	"github.com/jinzhu/gorm"
	di "github.com/sarulabs/di/v2"

	"github.com/kzmake/micro-kit/service/task/pkg/config"
)

// Production は本番環境用のDIコンテナ定義です。
var Production = []di.Def{
	{
		Name:  "logWriter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			c := ctn.Get("config").(*config.Config)
			switch c.Log.Out {
			case "stderr":
				return os.Stderr, nil
			default:
				return os.Stdout, nil
			}
		},
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
		Build: buildLoggerAssistant("logWriter"),
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
		Name:  "getTaskInputPort",
		Scope: di.App,
		Build: buildGetTaskInteractor("businessManager", "taskRepository"),
	},
	{
		Name:  "taskQueryController",
		Scope: di.App,
		Build: buildTaskQueryController("getTaskInputPort"),
	},
	{
		Name:  "taskCommandController",
		Scope: di.App,
		Build: buildTaskCommandController("createTaskInputPort"),
	},
	{
		Name:  "taskController",
		Scope: di.App,
		Build: buildTaskController("taskQueryController", "taskCommandController"),
	},
}