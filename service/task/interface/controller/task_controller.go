package controller

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

type taskController struct {
	createTaskInputPort port.CreateTaskInputPort
}

// interfaces
var _ proto.TaskServiceHandler = (*taskController)(nil)

// NewTaskController はタスクに関する Controller を生成します。
func NewTaskController(createTaskInputPort port.CreateTaskInputPort) proto.TaskServiceHandler {
	return &taskController{
		createTaskInputPort: createTaskInputPort,
	}
}

// Create は input / output を制御し、タスク生成処理を行います。
func (c *taskController) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.CreateResponse) error {
	if err := req.Validate(); err != nil {
		return err
	}

	input := &port.CreateTaskInputData{
		Description: req.Description.Value,
	}

	output := c.createTaskInputPort.Handle(ctx, input)

	if output.Error != nil {
		rsp.Task = &proto.Task{
			Description: req.GetDescription(),
		}

		return nil
	}

	return output.Error
}
