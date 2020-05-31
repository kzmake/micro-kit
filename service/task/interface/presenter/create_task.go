package presenter

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task/usecase/output"
)

type createTaskPresenter struct{}

// interfaces
var _ output.CreateTaskPort = (*createTaskPresenter)(nil)

// NewCreateTaskPresenter はタスク生成に関する  Presenter を生成します。
func NewCreateTaskPresenter() output.CreateTaskPort {
	return &createTaskPresenter{}
}

// Handle は OutputData を生成します。
func (p *createTaskPresenter) Handle(ctx context.Context, task *aggregate.Task, err error) *output.CreateTaskData {
	return &output.CreateTaskData{
		Task:  task,
		Error: err,
	}
}
