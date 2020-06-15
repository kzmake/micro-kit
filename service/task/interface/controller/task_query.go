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

// TaskQuery はタスクに関する Query 系の Controller の定義です。
type TaskQuery struct {
	getTaskInputPort port.GetTask
}

// NewTaskQuery はタスクに関する Query 系の Controller を生成します。
func NewTaskQuery(
	getTaskInputPort port.GetTask,
) *TaskQuery {
	return &TaskQuery{
		getTaskInputPort: getTaskInputPort,
	}
}

// Get は input / output を制御し、タスク取得処理を行います。
func (c *TaskQuery) Get(
	ctx context.Context,
	req *proto.GetRequest,
	rsp *proto.GetResponse,
) error {
	if err := req.Validate(); err != nil {
		var validationErr proto.GetRequestValidationError
		if xerrors.As(err, &validationErr) {
			switch validationErr.Field() { // nolint:gocritic
			case "Id":
				if validationErr.Reason() == "value is required and must not be nil." {
					return errors.BadRequest("InvalidParameterRequired.Id", "The request is missing a required parameter.")
				}
				return errors.BadRequest("InvalidParameterFormat.Id", "The parameter id is not valid format.")
			}
		}

		return errors.InternalServerError("InternalServerError", "An internal error has occurred. Please try your query again at a later time.")
	}

	in := &port.GetTaskInputData{
		ID: req.GetId().GetValue(),
	}

	out := c.getTaskInputPort.Handle(ctx, in)
	if err := out.Error; err != nil {
		return out.Error
	}

	task := &proto.Task{
		Id:          &wrappers.StringValue{Value: string(out.Task.ID)},
		Description: &wrappers.StringValue{Value: string(out.Task.Description)},
		CreatedAt:   &timestamp.Timestamp{Seconds: out.Task.CreatedAt.Unix()},
		UpdatedAt:   &timestamp.Timestamp{Seconds: out.Task.UpdatedAt.Unix()},
	}
	if out.Task.DeletedAt != nil {
		task.DeletedAt = &timestamp.Timestamp{Seconds: out.Task.DeletedAt.Unix()}
	}

	rsp.Result = task

	return nil
}
