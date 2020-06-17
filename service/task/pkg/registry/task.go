package registry

import (
	di "github.com/sarulabs/di/v2"
)

// Task はタスクに関するDIコンテナ定義です。
var Task = []di.Def{
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
