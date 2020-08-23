package controller

import (
	"context"
	"fmt"
	"testing"
	"time"

	"golang.org/x/xerrors"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	merrors "github.com/micro/go-micro/v2/errors"
	"github.com/stretchr/testify/require"

	"github.com/kzmake/micro-kit/service/task-service/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task-service/domain/errors"
	"github.com/kzmake/micro-kit/service/task-service/domain/vo"
	"github.com/kzmake/micro-kit/service/task-service/interface/proto"
	"github.com/kzmake/micro-kit/service/task-service/usecase/port"
)

func TestNewTaskQuery(t *testing.T) {
	c := NewTaskQuery(&mockListTasksInput{}, &mockGetTaskInput{})

	require.NotNil(t, c)
}

func TestTaskList(t *testing.T) {
	now := time.Now().UTC()

	testcases := []struct {
		tasks    []*aggregate.Task
		req      *proto.ListRequest
		expected *proto.ListResponse
	}{
		{
			req:   &proto.ListRequest{},
			tasks: []*aggregate.Task{},
			expected: &proto.ListResponse{
				Results: []*proto.Task{},
			},
		},
		{
			req: &proto.ListRequest{},
			tasks: []*aggregate.Task{
				{
					ID:          vo.TaskID("01D78XZ44G0000000000000000"),
					Description: vo.Description("hoge"),
					CreatedAt:   now,
					UpdatedAt:   now,
					DeletedAt:   nil,
				},
			},
			expected: &proto.ListResponse{
				Results: []*proto.Task{
					{
						Id:          &wrappers.StringValue{Value: "01D78XZ44G0000000000000000"},
						Description: &wrappers.StringValue{Value: "hoge"},
						CreatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
						UpdatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
					},
				},
			},
		},
		{
			req: &proto.ListRequest{},
			tasks: []*aggregate.Task{
				{
					ID:          vo.TaskID("01D78XZ44G0000000000000000"),
					Description: vo.Description("hoge"),
					CreatedAt:   now,
					UpdatedAt:   now,
					DeletedAt:   nil,
				},
				{
					ID:          vo.TaskID("01D78XZ44G0000000000000001"),
					Description: vo.Description(""),
					CreatedAt:   now,
					UpdatedAt:   now,
					DeletedAt:   nil,
				},
			},
			expected: &proto.ListResponse{
				Results: []*proto.Task{
					{
						Id:          &wrappers.StringValue{Value: "01D78XZ44G0000000000000000"},
						Description: &wrappers.StringValue{Value: "hoge"},
						CreatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
						UpdatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
					},
					{
						Id:          &wrappers.StringValue{Value: "01D78XZ44G0000000000000001"},
						Description: &wrappers.StringValue{Value: ""},
						CreatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
						UpdatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
					},
				},
			},
		},
	}

	// nolint:dupl
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.req.String()), func(t *testing.T) {
			mock := &mockListTasksInput{fn: func(_ context.Context, in *port.ListTasksInputData) *port.ListTasksOutputData {
				return &port.ListTasksOutputData{Tasks: tc.tasks}
			}}
			c := NewTask(NewTaskQuery(mock, nil), NewTaskCommand(nil, nil))
			actual := &proto.ListResponse{}

			_ = c.List(context.TODO(), tc.req, actual)

			expected := tc.expected
			require.NotNil(t, actual.GetResults())
			require.EqualValues(t, expected, actual)
			require.Equal(t, expected.String(), actual.String())
		})
	}
}

func TestTaskList_Error_WhenUsecaseError(t *testing.T) {
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
			uerr: xerrors.Errorf("タスク一覧取得時に予期せぬエラー: %w", errors.WrapCode(errors.Unexpected, fmt.Errorf("failed"))),
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
			mock := &mockListTasksInput{fn: func(_ context.Context, in *port.ListTasksInputData) *port.ListTasksOutputData {
				return &port.ListTasksOutputData{Error: tc.uerr}
			}}
			c := NewTask(NewTaskQuery(mock, nil), NewTaskCommand(nil, nil))

			err := c.List(context.TODO(), &proto.ListRequest{}, &proto.ListResponse{})

			require.Error(t, err)
			require.Equal(t, tc.expected, err)
		})
	}
}

func TestTaskGet(t *testing.T) {
	now := time.Now().UTC()

	testcases := []struct {
		task     *aggregate.Task
		req      *proto.GetRequest
		expected *proto.GetResponse
	}{
		// nolint:dupl
		{
			req: &proto.GetRequest{Id: &wrappers.StringValue{Value: "01D78XZ44G0000000000000000"}},
			task: &aggregate.Task{
				ID:          vo.TaskID("01D78XZ44G0000000000000000"),
				Description: vo.Description("hoge"),
				CreatedAt:   now,
				UpdatedAt:   now,
				DeletedAt:   nil,
			},
			expected: &proto.GetResponse{
				Result: &proto.Task{
					Id:          &wrappers.StringValue{Value: "01D78XZ44G0000000000000000"},
					Description: &wrappers.StringValue{Value: "hoge"},
					CreatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
					UpdatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
				},
			},
		},
		// nolint:dupl
		{
			req: &proto.GetRequest{Id: &wrappers.StringValue{Value: "01D78XZ44G0000000000000001"}},
			task: &aggregate.Task{
				ID:          vo.TaskID("01D78XZ44G0000000000000001"),
				Description: vo.Description(""),
				CreatedAt:   now,
				UpdatedAt:   now,
				DeletedAt:   nil,
			},
			expected: &proto.GetResponse{
				Result: &proto.Task{
					Id:          &wrappers.StringValue{Value: "01D78XZ44G0000000000000001"},
					Description: &wrappers.StringValue{Value: ""},
					CreatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
					UpdatedAt:   &timestamp.Timestamp{Seconds: now.UTC().Unix()},
				},
			},
		},
	}

	// nolint:dupl
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d: %s", i, tc.req.String()), func(t *testing.T) {
			mock := &mockGetTaskInput{fn: func(_ context.Context, in *port.GetTaskInputData) *port.GetTaskOutputData {
				return &port.GetTaskOutputData{Task: tc.task}
			}}
			c := NewTask(NewTaskQuery(nil, mock), NewTaskCommand(nil, nil))
			actual := &proto.GetResponse{}

			_ = c.Get(context.TODO(), tc.req, actual)

			expected := tc.expected
			require.NotNil(t, actual.GetResult())
			require.EqualValues(t, expected, actual)
			require.Equal(t, expected.String(), actual.String())
		})
	}
}

// nolint:dupl
func TestTaskGet_Error_WhenInvalidParams(t *testing.T) {
	testcases := []struct {
		req      *proto.GetRequest
		expected error
	}{

		{
			req:      nil,
			expected: merrors.BadRequest(errors.IllegalInputBody.String(), "The request body is not appropriate."),
		},
		{
			req:      &proto.GetRequest{Id: &wrappers.StringValue{Value: "01d78xZ44g0000000000000000"}},
			expected: merrors.BadRequest(errors.IllegalInputTaskID.String(), "The requested id is invalid."),
		},
		{
			req:      &proto.GetRequest{Id: &wrappers.StringValue{Value: "01D78XZ44G0000000000000000X"}},
			expected: merrors.BadRequest(errors.IllegalInputTaskID.String(), "The requested id is invalid."),
		},
		{
			req:      &proto.GetRequest{Id: &wrappers.StringValue{Value: "01D78XZ44G000000000000000"}},
			expected: merrors.BadRequest(errors.IllegalInputTaskID.String(), "The requested id is invalid."),
		},
		{
			req:      &proto.GetRequest{Id: &wrappers.StringValue{Value: ""}},
			expected: merrors.BadRequest(errors.IllegalInputTaskID.String(), "The requested id is invalid."),
		},
	}
	for i, tc := range testcases {
		i := i
		tc := tc

		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			mock := &mockGetTaskInput{}
			c := NewTask(NewTaskQuery(nil, mock), NewTaskCommand(nil, nil))

			err := c.Get(context.TODO(), tc.req, &proto.GetResponse{})

			require.Error(t, err)
			require.Equal(t, tc.expected, err)
		})
	}
}

func TestTaskGet_Error_WhenUsecaseError(t *testing.T) {
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
			uerr: xerrors.Errorf("タスク一覧取得時に予期せぬエラー: %w", errors.WrapCode(errors.Unexpected, fmt.Errorf("failed"))),
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
			mock := &mockGetTaskInput{fn: func(_ context.Context, in *port.GetTaskInputData) *port.GetTaskOutputData {
				return &port.GetTaskOutputData{Error: tc.uerr}
			}}
			c := NewTask(NewTaskQuery(nil, mock), NewTaskCommand(nil, nil))

			err := c.Get(context.TODO(),
				&proto.GetRequest{Id: &wrappers.StringValue{Value: "01D78XZ44G0000000000000000"}},
				&proto.GetResponse{},
			)

			require.Error(t, err)
			require.Equal(t, tc.expected, err)
		})
	}
}
