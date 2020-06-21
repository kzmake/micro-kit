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
