package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kzmake/micro-kit/service/task-service/usecase/port"
)

type (
	mockCreateTaskInput struct {
		fn func(context.Context, *port.CreateTaskInputData) *port.CreateTaskOutputData
	}
	mockListTasksInput struct {
		fn func(context.Context, *port.ListTasksInputData) *port.ListTasksOutputData
	}
	mockGetTaskInput struct {
		fn func(context.Context, *port.GetTaskInputData) *port.GetTaskOutputData
	}
	mockDeleteTaskInput struct {
		fn func(context.Context, *port.DeleteTaskInputData) *port.DeleteTaskOutputData
	}
)

var _ port.CreateTask = &mockCreateTaskInput{}
var _ port.ListTasks = &mockListTasksInput{}
var _ port.GetTask = &mockGetTaskInput{}
var _ port.DeleteTask = &mockDeleteTaskInput{}

func (m *mockCreateTaskInput) Handle(ctx context.Context, in *port.CreateTaskInputData) *port.CreateTaskOutputData {
	return m.fn(ctx, in)
}

func (m *mockListTasksInput) Handle(ctx context.Context, in *port.ListTasksInputData) *port.ListTasksOutputData {
	return m.fn(ctx, in)
}

func (m *mockGetTaskInput) Handle(ctx context.Context, in *port.GetTaskInputData) *port.GetTaskOutputData {
	return m.fn(ctx, in)
}

func (m *mockDeleteTaskInput) Handle(ctx context.Context, in *port.DeleteTaskInputData) *port.DeleteTaskOutputData {
	return m.fn(ctx, in)
}

func TestNewTask(t *testing.T) {
	tq := NewTaskQuery(&mockListTasksInput{}, &mockGetTaskInput{})
	tc := NewTaskCommand(&mockCreateTaskInput{}, &mockDeleteTaskInput{})
	c := NewTask(tq, tc)

	require.NotNil(t, tq)
	require.NotNil(t, tc)
	require.NotNil(t, c)
}
