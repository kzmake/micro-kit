package controller

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro/go-micro/v2/errors"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

// TaskCommand はタスクに関する Command 系の Controller の定義です。
type TaskCommand struct {
	createTaskInputPort port.CreateTask
}

// NewTaskCommand はタスクに関する Command 系の Controller を生成します。
func NewTaskCommand(
	createTaskInputPort port.CreateTask,
) *TaskCommand {
	return &TaskCommand{
		createTaskInputPort: createTaskInputPort,
	}
}

// Create は input / output を制御し、タスク生成処理を行います。
func (c *TaskCommand) Create(
	ctx context.Context,
	req *proto.CreateRequest,
	rsp *proto.CreateResponse,
) error {
	if err := req.Validate(); err != nil {
		return errors.BadRequest("InvalidParameterIllegalInput.Body", "The request body is not appropriate.")
	}

	in := &port.CreateTaskInputData{
		Description: req.GetDescription().GetValue(),
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
