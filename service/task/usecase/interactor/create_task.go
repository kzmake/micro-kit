package interactor

import (
	"context"

	// "golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task/domain/vo"
	"github.com/kzmake/micro-kit/service/task/usecase/input"
	"github.com/kzmake/micro-kit/service/task/usecase/output"
)

type createTaskInteractor struct {
	outputPort output.CreateTaskPort
}

// NewCreateTaskInteractor はタスクに関する Interactor を生成します。
func NewCreateTaskInteractor(outputPort output.CreateTaskPort) input.CreateTaskPort {
	return &createTaskInteractor{
		outputPort: outputPort,
	}
}

// Handle は InputData をもとにタスク生成を行いを OutputData を生成します。
func (i *createTaskInteractor) Handle(ctx context.Context, in *input.CreateTaskData) *output.CreateTaskData {
	out := &aggregate.Task{
		ID:          vo.ID("uniq_id"),
		Description: vo.Description("hogehoge"),
	}

	// createdTask, err := repository.Create(ctx, task)
	// if err != nil {
	//		xerrors.Errorf("error: %w", err)
	// }

	return i.outputPort.Handle(ctx, out, nil)
}
