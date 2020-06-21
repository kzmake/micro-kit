package controller

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro/go-micro/v2/errors"
	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/service/task/interface/proto"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
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
	if err := req.Validate(); err != nil {
		var validationErr proto.CreateRequestValidationError
		if xerrors.As(err, &validationErr) {
			switch validationErr.Field() { // nolint:gocritic
			case "Description":
				return errors.BadRequest("InvalidParameterFormat.Description", "The parameter description is not valid format.")
			}
		}

		return errors.InternalServerError("InternalServerError", "An internal error has occurred. Please try your query again at a later time.")
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

// Delete は input / output を制御し、タスク削除処理を行います。
func (c *TaskCommand) Delete(
	ctx context.Context,
	req *proto.DeleteRequest,
	rsp *proto.DeleteResponse,
) error {
	if err := req.Validate(); err != nil {
		var validationErr proto.DeleteRequestValidationError
		if xerrors.As(err, &validationErr) {
			switch validationErr.Field() { // nolint:gocritic
			case "Id": // nolint:goconst
				return errors.BadRequest("InvalidParameterFormat.Id", "The parameter id is not valid format.")
			}
		}

		return errors.InternalServerError("InternalServerError", "An internal error has occurred. Please try your query again at a later time.")
	}

	in := &port.DeleteTaskInputData{
		ID: req.GetId().GetValue(),
	}

	out := c.deleteTaskInputPort.Handle(ctx, in)
	if err := out.Error; err != nil {
		return out.Error
	}

	return nil
}
