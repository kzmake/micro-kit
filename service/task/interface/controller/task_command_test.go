package controller

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"golang.org/x/xerrors"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	merrors "github.com/micro/go-micro/v2/errors"
	"github.com/stretchr/testify/require"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task/domain/errors"
	"github.com/kzmake/micro-kit/service/task/domain/vo"
	"github.com/kzmake/micro-kit/service/task/interface/proto"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

func TestNewTaskCommand(t *testing.T) {
	c := NewTaskCommand(&mockCreateTaskInput{}, &mockDeleteTaskInput{})

	require.NotNil(t, c)
}

func TestTaskCreate(t *testing.T) {
	const ulid = "01D78XZ44G0000000000000000"
	now := time.Now().UTC()

	testcases := []struct {
		req      *proto.CreateRequest
		expected *proto.CreateResponse
	}{
		{
			req: &proto.CreateRequest{},
			expected: &proto.CreateResponse{
				Result: &proto.Task{
					Id:          &wrappers.StringValue{Value: ulid},
					Description: &wrappers.StringValue{Value: ""},
					CreatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
					UpdatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
				},
			},
		},
		{
			req: &proto.CreateRequest{
				Description: &wrappers.StringValue{Value: "hoge"},
			},
			expected: &proto.CreateResponse{
				Result: &proto.Task{
					Id:          &wrappers.StringValue{Value: ulid},
					Description: &wrappers.StringValue{Value: "hoge"},
					CreatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
					UpdatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
				},
			},
		},
		{
			req: &proto.CreateRequest{
				Description: &wrappers.StringValue{Value: "あ𩸽🍣"},
			},
			expected: &proto.CreateResponse{
				Result: &proto.Task{
					Id:          &wrappers.StringValue{Value: ulid},
					Description: &wrappers.StringValue{Value: "あ𩸽🍣"},
					CreatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
					UpdatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
				},
			},
		},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.req.String()), func(t *testing.T) {
			mock := &mockCreateTaskInput{fn: func(_ context.Context, in *port.CreateTaskInputData) *port.CreateTaskOutputData {
				return &port.CreateTaskOutputData{
					Task: &aggregate.Task{
						ID:          vo.TaskID(ulid),
						Description: vo.Description(in.Description),
						CreatedAt:   now,
						UpdatedAt:   now,
						DeletedAt:   nil,
					},
				}
			}}
			c := NewTask(NewTaskQuery(nil, nil), NewTaskCommand(mock, nil))
			actual := &proto.CreateResponse{}

			_ = c.Create(context.TODO(), tc.req, actual)

			expected := tc.expected
			require.NotNil(t, actual.GetResult())
			require.EqualValues(t, expected, actual)
			require.Equal(t, expected.String(), actual.String())
		})
	}
}

func TestTaskCreate_WhenMaxLengthDescription(t *testing.T) {
	const ulid = "01D78XZ44G0000000000000000"
	now := time.Now().UTC()

	testcases := []struct {
		req *proto.CreateRequest
	}{
		{req: &proto.CreateRequest{Description: &wrappers.StringValue{Value: strings.Repeat("a", 255)}}},
		{req: &proto.CreateRequest{Description: &wrappers.StringValue{Value: strings.Repeat("あ", 255)}}},
		{req: &proto.CreateRequest{Description: &wrappers.StringValue{Value: strings.Repeat("𩸽", 255)}}},
		{req: &proto.CreateRequest{Description: &wrappers.StringValue{Value: strings.Repeat("🍣", 255)}}},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			mock := &mockCreateTaskInput{fn: func(_ context.Context, in *port.CreateTaskInputData) *port.CreateTaskOutputData {
				return &port.CreateTaskOutputData{
					Task: &aggregate.Task{
						ID:          vo.TaskID(ulid),
						Description: vo.Description(in.Description),
						CreatedAt:   now,
						UpdatedAt:   now,
						DeletedAt:   nil,
					},
				}
			}}
			c := NewTask(NewTaskQuery(nil, nil), NewTaskCommand(mock, nil))

			err := c.Create(context.TODO(), tc.req, &proto.CreateResponse{})

			require.NoError(t, err)
		})
	}
}

func TestTaskCreate_Error_WhenInvalidParams(t *testing.T) {
	testcases := []struct {
		req      *proto.CreateRequest
		expected error
	}{
		{
			req:      nil,
			expected: merrors.BadRequest(errors.IllegalInputBody.String(), "The request body is not appropriate."),
		},
		{
			req:      &proto.CreateRequest{Description: &wrappers.StringValue{Value: strings.Repeat("a", 256)}},
			expected: merrors.BadRequest(errors.IllegalInputDescription.String(), "The requested id is invalid."),
		},
		{
			req:      &proto.CreateRequest{Description: &wrappers.StringValue{Value: strings.Repeat("あ", 256)}},
			expected: merrors.BadRequest(errors.IllegalInputDescription.String(), "The requested id is invalid."),
		},
		{
			req:      &proto.CreateRequest{Description: &wrappers.StringValue{Value: strings.Repeat("𩸽", 256)}},
			expected: merrors.BadRequest(errors.IllegalInputDescription.String(), "The requested id is invalid."),
		},
		{
			req:      &proto.CreateRequest{Description: &wrappers.StringValue{Value: strings.Repeat("🍣", 256)}},
			expected: merrors.BadRequest(errors.IllegalInputDescription.String(), "The requested id is invalid."),
		},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			mock := &mockCreateTaskInput{}
			c := NewTask(NewTaskQuery(nil, nil), NewTaskCommand(mock, nil))

			err := c.Create(context.TODO(), tc.req, &proto.CreateResponse{})

			require.Error(t, err)
			require.Equal(t, tc.expected, err)
		})
	}
}

// nolint:dupl
func TestTaskCreate_Error_WhenUsecaseError(t *testing.T) {
	testcases := []struct {
		uerr     error
		expected error
	}{
		{
			uerr: fmt.Errorf("unknown"),
			expected: merrors.InternalServerError(
				"InternalServerError",
				"An internal error has occurred. Please try your query again at a later time.",
			),
		},
		{
			uerr:     xerrors.Errorf("タスク作成時に重複: %w", errors.WrapCode(errors.DuplicateTask, fmt.Errorf("failed"))),
			expected: merrors.BadRequest(errors.DuplicateTask.String(), "The task already exists."),
		},
		{
			uerr: xerrors.Errorf("タスク作成時に予期せぬエラー: %w", errors.WrapCode(errors.Unexpected, fmt.Errorf("failed"))),
			expected: merrors.InternalServerError(
				"InternalServerError",
				"An internal error has occurred. Please try your query again at a later time.",
			),
		},
	}

	// nolint:dupl
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.uerr), func(t *testing.T) {
			mock := &mockCreateTaskInput{fn: func(_ context.Context, in *port.CreateTaskInputData) *port.CreateTaskOutputData {
				return &port.CreateTaskOutputData{Error: tc.uerr}
			}}
			c := NewTask(NewTaskQuery(nil, nil), NewTaskCommand(mock, nil))

			err := c.Create(context.TODO(), &proto.CreateRequest{}, &proto.CreateResponse{})

			require.Error(t, err)
			require.Equal(t, tc.expected, err)
		})
	}
}

func TestTaskDelete(t *testing.T) {
	testcases := []struct {
		req      *proto.DeleteRequest
		expected *proto.DeleteResponse
	}{
		{
			req:      &proto.DeleteRequest{Id: &wrappers.StringValue{Value: "01D78XZ44G0000000000000000"}},
			expected: &proto.DeleteResponse{},
		},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.req.String()), func(t *testing.T) {
			mock := &mockDeleteTaskInput{fn: func(_ context.Context, in *port.DeleteTaskInputData) *port.DeleteTaskOutputData {
				return &port.DeleteTaskOutputData{}
			}}
			c := NewTask(NewTaskQuery(nil, nil), NewTaskCommand(nil, mock))
			actual := &proto.DeleteResponse{}

			_ = c.Delete(context.TODO(), tc.req, actual)

			expected := tc.expected
			require.EqualValues(t, expected, actual)
			require.Equal(t, expected.String(), actual.String())
		})
	}
}

// nolint:dupl
func TestTaskDelete_Error_WhenInvalidParams(t *testing.T) {
	testcases := []struct {
		req      *proto.DeleteRequest
		expected error
	}{

		{
			req:      nil,
			expected: merrors.BadRequest(errors.IllegalInputBody.String(), "The request body is not appropriate."),
		},
		{
			req:      &proto.DeleteRequest{Id: &wrappers.StringValue{Value: "01d78xZ44g0000000000000000"}},
			expected: merrors.BadRequest(errors.IllegalInputTaskID.String(), "The requested id is invalid."),
		},
		{
			req:      &proto.DeleteRequest{Id: &wrappers.StringValue{Value: "01D78XZ44G0000000000000000X"}},
			expected: merrors.BadRequest(errors.IllegalInputTaskID.String(), "The requested id is invalid."),
		},
		{
			req:      &proto.DeleteRequest{Id: &wrappers.StringValue{Value: "01D78XZ44G000000000000000"}},
			expected: merrors.BadRequest(errors.IllegalInputTaskID.String(), "The requested id is invalid."),
		},
		{
			req:      &proto.DeleteRequest{Id: &wrappers.StringValue{Value: ""}},
			expected: merrors.BadRequest(errors.IllegalInputTaskID.String(), "The requested id is invalid."),
		},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			mock := &mockDeleteTaskInput{}
			c := NewTask(NewTaskQuery(nil, nil), NewTaskCommand(nil, mock))

			err := c.Delete(context.TODO(), tc.req, &proto.DeleteResponse{})

			require.Error(t, err)
			require.Equal(t, tc.expected, err)
		})
	}
}

func TestTaskDelete_Error_WhenUsecaseError(t *testing.T) {
	testcases := []struct {
		uerr     error
		expected error
	}{
		{
			uerr: fmt.Errorf("unknown"),
			expected: merrors.InternalServerError(
				"InternalServerError",
				"An internal error has occurred. Please try your query again at a later time.",
			),
		},
		{
			uerr:     xerrors.Errorf("タスクが存在しない: %w", errors.WrapCode(errors.NotFoundTask, fmt.Errorf("failed"))),
			expected: merrors.NotFound(errors.NotFoundTask.String(), "The task does not found."),
		},
		{
			uerr: xerrors.Errorf("タスク削除時に予期せぬエラー: %w", errors.WrapCode(errors.Unexpected, fmt.Errorf("failed"))),
			expected: merrors.InternalServerError(
				"InternalServerError",
				"An internal error has occurred. Please try your query again at a later time.",
			),
		},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.uerr), func(t *testing.T) {
			mock := &mockDeleteTaskInput{fn: func(_ context.Context, in *port.DeleteTaskInputData) *port.DeleteTaskOutputData {
				return &port.DeleteTaskOutputData{Error: tc.uerr}
			}}
			c := NewTask(NewTaskQuery(nil, nil), NewTaskCommand(nil, mock))

			err := c.Delete(context.TODO(),
				&proto.DeleteRequest{Id: &wrappers.StringValue{Value: "01D78XZ44G0000000000000000"}},
				&proto.DeleteResponse{},
			)

			require.Error(t, err)
			require.Equal(t, tc.expected, err)
		})
	}
}
