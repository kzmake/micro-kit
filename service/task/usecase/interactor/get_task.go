package interactor

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task/domain/repository"
	"github.com/kzmake/micro-kit/service/task/domain/vo"
	"github.com/kzmake/micro-kit/service/task/usecase/business"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

type getTask struct {
	manager        business.Manager
	taskRepository repository.Task
}

// NewGetTask はタスク取得に関する Interactor を生成します。
func NewGetTask(
	manager business.Manager,
	taskRepository repository.Task,
) port.GetTask {
	return &getTask{
		manager:        manager,
		taskRepository: taskRepository,
	}
}

// Handle は InputData をもとにタスク取得を行い OutputData を生成します。
func (i *getTask) Handle(
	ctx context.Context,
	in *port.GetTaskInputData,
) *port.GetTaskOutputData {
	v, err := i.manager.Execute(ctx, func(cctx context.Context) (interface{}, error) {
		return i.handle(cctx, in)
	})

	return &port.GetTaskOutputData{Task: v.(*aggregate.Task), Error: err}
}

func (i *getTask) handle(
	ctx context.Context,
	in *port.GetTaskInputData,
) (interface{}, error) {
	taskID := vo.TaskID(in.ID)

	task, err := i.taskRepository.Find(ctx, taskID)
	if err != nil {
		return nil, xerrors.Errorf("Findに失敗しました: %w", err)
	}

	return task, nil
}
