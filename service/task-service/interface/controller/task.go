package controller

import (
	"github.com/kzmake/micro-kit/service/task-service/interface/proto"
)

type task struct {
	*TaskQuery
	*TaskCommand
}

// interfaces
var _ proto.TaskServiceHandler = (*task)(nil)

// NewTask はタスクに関する Controller を生成します。
func NewTask(query *TaskQuery, command *TaskCommand) proto.TaskServiceHandler {
	return &task{TaskQuery: query, TaskCommand: command}
}
