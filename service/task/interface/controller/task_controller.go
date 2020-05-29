package controller

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
	"github.com/kzmake/micro-kit/service/task/usecase/input"
)

type taskController struct {
	createTaskInputPort input.CreateTaskPort
}

// interfaces
var _ proto.TaskServiceHandler = (*taskController)(nil)

// NewTaskController はタスクに関する Controller を生成します。
func NewTaskController(createTaskInputPort input.CreateTaskPort) proto.TaskServiceHandler {
	return &taskController{
		createTaskInputPort: createTaskInputPort,
	}
}

// Create は input / output を制御し、タスク生成処理を行います。
func (c *taskController) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.CreateResponse) error {
	if err := req.Validate(); err != nil {
		return err
	}

	in := &input.CreateTaskData{
		Description: req.Description.Value,
	}

	out := c.createTaskInputPort.Handle(ctx, in)

	if out.Error != nil {
		rsp.Task = &proto.Task{
			Description: req.GetDescription(),
		}

		return nil
	}

	return out.Error
}
