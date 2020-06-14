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
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

type createTask struct {
	idRepository   repository.ID
	taskRepository repository.Task
}

// NewCreateTask はタスクに関する Interactor を生成します。
func NewCreateTask(
	idRepository repository.ID,
	taskRepository repository.Task,
) port.CreateTask {
	return &createTask{
		idRepository:   idRepository,
		taskRepository: taskRepository,
	}
}

// Handle は InputData をもとにタスク生成を行いを OutputData を生成します。
func (i *createTask) Handle(
	ctx context.Context,
	in *port.CreateTaskInputData,
) *port.CreateTaskOutputData {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	task := &aggregate.Task{
		ID:          vo.ID(id),
		Description: vo.Description(in.Description),
	}

	if err := i.taskRepository.Save(ctx, task); err != nil {
		return &port.CreateTaskOutputData{Error: xerrors.Errorf("Saveに失敗しました: %w", err)}
	}

	task, err := i.taskRepository.Find(ctx, task.ID)
	if err != nil {
		return &port.CreateTaskOutputData{Error: xerrors.Errorf("Findに失敗しました: %w", err)}
	}

	return &port.CreateTaskOutputData{Task: task, Error: nil}
}
