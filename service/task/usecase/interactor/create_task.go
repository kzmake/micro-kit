package interactor

import (
	"context"
	"math/rand"
	"time"

	// "golang.org/x/xerrors"
	ulid "github.com/oklog/ulid/v2"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task/domain/vo"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

type createTaskInteractor struct {
}

// NewCreateTaskInteractor はタスクに関する Interactor を生成します。
func NewCreateTaskInteractor() port.CreateTask {
	return &createTaskInteractor{}
}

// Handle は InputData をもとにタスク生成を行いを OutputData を生成します。
func (i *createTaskInteractor) Handle(
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

	time.Sleep(1 * time.Second)

	return &port.CreateTaskOutputData{
		Task:  task,
		Error: nil,
	}
}
