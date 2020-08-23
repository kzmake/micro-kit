package controller

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/kzmake/micro-kit/service/task-service/domain/errors"
	"github.com/kzmake/micro-kit/service/task-service/interface/proto"
	"github.com/kzmake/micro-kit/service/task-service/usecase/port"
)

// TaskCommand はタスクに関する Command 系の Controller の定義です。
type TaskCommand struct {
	createTaskInputPort port.CreateTask
	deleteTaskInputPort port.DeleteTask
}

// NewTaskCommand はタスクに関する Command 系の Controller を生成します。
func NewTaskCommand(
	createTaskInputPort port.CreateTask,
	deleteTaskInputPort port.DeleteTask,
) *TaskCommand {
	return &TaskCommand{
		createTaskInputPort: createTaskInputPort,
		deleteTaskInputPort: deleteTaskInputPort,
	}
}

// Create は input / output を制御し、タスク生成処理を行います。
func (c *TaskCommand) Create(
	ctx context.Context,
	req *proto.CreateRequest,
	rsp *proto.CreateResponse,
) error {
	if req == nil {
		return encodeError(ctx, errors.NewCode(errors.IllegalInputBody))
	}

	if err := req.Validate(); err != nil {
		var validationErr proto.CreateRequestValidationError
		if xerrors.As(err, &validationErr) {
			switch validationErr.Field() { // nolint:gocritic
			case "Description":
				return encodeError(ctx, errors.WrapCode(errors.IllegalInputDescription, err))
			}
		}
	}

	in := &port.CreateTaskInputData{
		Description: req.GetDescription().GetValue(),
	}

	out := c.createTaskInputPort.Handle(ctx, in)
	if err := out.Error; err != nil {
		return encodeError(ctx, out.Error)
	}

	rsp.Result = &proto.Task{
		Id:          &wrappers.StringValue{Value: string(out.Task.ID)},
		Description: &wrappers.StringValue{Value: string(out.Task.Description)},
		CreatedAt:   &timestamp.Timestamp{Seconds: out.Task.CreatedAt.Unix()},
		UpdatedAt:   &timestamp.Timestamp{Seconds: out.Task.UpdatedAt.Unix()},
	}

	return nil
}

// Delete は input / output を制御し、タスク削除処理を行います。
func (c *TaskCommand) Delete(
	ctx context.Context,
	req *proto.DeleteRequest,
	rsp *proto.DeleteResponse,
) error {
	if req == nil {
		return encodeError(ctx, errors.NewCode(errors.IllegalInputBody))
	}

	if err := req.Validate(); err != nil {
		var validationErr proto.DeleteRequestValidationError
		if xerrors.As(err, &validationErr) {
			switch validationErr.Field() { // nolint:gocritic
			case "Id": // nolint:goconst
				return encodeError(ctx, errors.WrapCode(errors.IllegalInputTaskID, err))
			}
		}
	}

	in := &port.DeleteTaskInputData{
		ID: req.GetId().GetValue(),
	}

	out := c.deleteTaskInputPort.Handle(ctx, in)
	if err := out.Error; err != nil {
		return encodeError(ctx, out.Error)
	}

	return nil
}
