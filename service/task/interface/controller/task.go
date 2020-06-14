package controller

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro/go-micro/v2/errors"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

type task struct {
	createTaskInputPort port.CreateTask
}

// interfaces
var _ proto.TaskServiceHandler = (*task)(nil)

// NewTask はタスクに関する Controller を生成します。
func NewTask(
	createTaskInputPort port.CreateTask,
) proto.TaskServiceHandler {
	return &task{
		createTaskInputPort: createTaskInputPort,
	}
}

// Create は input / output を制御し、タスク生成処理を行います。
func (c *task) Create(
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
		CreatedAt:   &timestamp.Timestamp{Seconds: out.Task.CreatedAt.Unix()},
		UpdatedAt:   &timestamp.Timestamp{Seconds: out.Task.UpdatedAt.Unix()},
	}
	if out.Task.DeletedAt != nil {
		rsp.Result.DeletedAt = &timestamp.Timestamp{Seconds: out.Task.DeletedAt.Unix()}
	}

	return nil
}
