package interactor

import (
	"context"
	"math/rand"
	"time"

	"golang.org/x/xerrors"

	ulid "github.com/oklog/ulid/v2"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task/domain/repository"
	"github.com/kzmake/micro-kit/service/task/domain/vo"
	"github.com/kzmake/micro-kit/service/task/usecase/business"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

type createTask struct {
	manager        business.Manager
	idRepository   repository.ID
	taskRepository repository.Task
}

// NewCreateTask はタスクに関する Interactor を生成します。
func NewCreateTask(
	manager business.Manager,
	idRepository repository.ID,
	taskRepository repository.Task,
) port.CreateTask {
	return &createTask{
		manager:        manager,
		idRepository:   idRepository,
		taskRepository: taskRepository,
	}
}

// Handle は InputData をもとにタスク生成を行いを OutputData を生成します。
func (i *createTask) Handle(
	ctx context.Context,
	in *port.CreateTaskInputData,
) *port.CreateTaskOutputData {
	v, err := i.manager.Execute(ctx, func(cctx context.Context) (interface{}, error) {
		return i.handle(cctx, in)
	})

	return &port.CreateTaskOutputData{Task: v.(*aggregate.Task), Error: err}
}

func (i *createTask) handle(
	ctx context.Context,
	in *port.CreateTaskInputData,
) (interface{}, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	task := &aggregate.Task{
		ID:          vo.ID(id),
		Description: vo.Description(in.Description),
	}

	if err := i.taskRepository.Save(ctx, task); err != nil {
		return nil, xerrors.Errorf("Saveに失敗しました: %w", err)
	}

	task, err := i.taskRepository.Find(ctx, task.ID)
	if err != nil {
		return nil, xerrors.Errorf("Findに失敗しました: %w", err)
	}

	return task, nil
}
