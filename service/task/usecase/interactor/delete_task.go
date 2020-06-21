package interactor

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/service/task/domain/repository"
	"github.com/kzmake/micro-kit/service/task/domain/vo"
	"github.com/kzmake/micro-kit/service/task/usecase/business"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

type deleteTask struct {
	manager        business.Manager
	taskRepository repository.Task
}

// NewDeleteTask はタスク削除に関する Interactor を生成します。
func NewDeleteTask(
	manager business.Manager,
	taskRepository repository.Task,
) port.DeleteTask {
	return &deleteTask{
		manager:        manager,
		taskRepository: taskRepository,
	}
}

// Handle は InputData をもとにタスク削除を行いを OutputData を生成します。
func (i *deleteTask) Handle(
	ctx context.Context,
	in *port.DeleteTaskInputData,
) *port.DeleteTaskOutputData {
	_, err := i.manager.Execute(ctx, func(cctx context.Context) (interface{}, error) {
		return nil, i.handle(cctx, in)
	})

	return &port.DeleteTaskOutputData{Error: err}
}

func (i *deleteTask) handle(
	ctx context.Context,
	in *port.DeleteTaskInputData,
) error {
	taskID := vo.TaskID(in.ID)

	task, err := i.taskRepository.Find(ctx, taskID)
	if err != nil {
		return xerrors.Errorf("Findに失敗しました: %w", err)
	}

	if err := i.taskRepository.Delete(ctx, task); err != nil {
		return xerrors.Errorf("Deleteに失敗しました: %w", err)
	}

	return nil
}
