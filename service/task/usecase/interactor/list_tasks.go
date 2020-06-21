package interactor

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task/domain/repository"
	"github.com/kzmake/micro-kit/service/task/usecase/business"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

type listTasks struct {
	manager        business.Manager
	taskRepository repository.Task
}

// NewListTasks はタスク一覧取得に関する Interactor を生成します。
func NewListTasks(
	manager business.Manager,
	taskRepository repository.Task,
) port.ListTasks {
	return &listTasks{
		manager:        manager,
		taskRepository: taskRepository,
	}
}

// Handle は InputData をもとにタスク一覧の取得を行い OutputData を生成します。
func (i *listTasks) Handle(
	ctx context.Context,
	in *port.ListTasksInputData,
) *port.ListTasksOutputData {
	v, err := i.manager.Execute(ctx, func(cctx context.Context) (interface{}, error) {
		return i.handle(cctx, in)
	})

	return &port.ListTasksOutputData{Tasks: v.([]*aggregate.Task), Error: err}
}

func (i *listTasks) handle(
	ctx context.Context,
	_ *port.ListTasksInputData,
) (interface{}, error) {
	tasks, err := i.taskRepository.List(ctx)
	if err != nil {
		return nil, xerrors.Errorf("Listに失敗しました: %w", err)
	}

	return tasks, nil
}
