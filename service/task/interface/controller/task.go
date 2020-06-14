package controller

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro/go-micro/v2/errors"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

type taskController struct {
	createTaskInputPort port.CreateTask
}

// interfaces
var _ proto.TaskServiceHandler = (*taskController)(nil)

// NewTaskController はタスクに関する Controller を生成します。
func NewTaskController(
	createTaskInputPort port.CreateTask,
) proto.TaskServiceHandler {
	return &taskController{
		createTaskInputPort: createTaskInputPort,
	}
}

// Create は input / output を制御し、タスク生成処理を行います。
func (c *taskController) Create(
	ctx context.Context,
	req *proto.CreateRequest,
	rsp *proto.CreateResponse,
) error {
	if err := req.Validate(); err != nil {
		return errors.BadRequest("InvalidParameterIllegalInput.Body", "The request body is not appropriate.")
	}

	in := &port.CreateTaskInputData{
		Description: req.Description.Value,
	}

	out := c.createTaskInputPort.Handle(ctx, in)
	if err := out.Error; err != nil {
		return out.Error
	}

	rsp.Result = &proto.Task{
		Id:          &wrappers.StringValue{Value: string(out.Task.ID)},
		Description: &wrappers.StringValue{Value: string(out.Task.Description)},
	}

	return nil
}
