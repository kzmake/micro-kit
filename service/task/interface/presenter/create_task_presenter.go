package presenter

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task/usecase/port"
)

type createTaskPresenter struct{}

// interfaces
var _ port.CreateTaskOutputPort = (*createTaskPresenter)(nil)

// NewCreateTaskPresenter はタスク生成に関する  Presenter を生成します。
func NewCreateTaskPresenter() port.CreateTaskOutputPort {
	return &createTaskPresenter{}
}

// Handle は OutputData を生成します。
func (p *createTaskPresenter) Handle(ctx context.Context, task *aggregate.Task, err error) *port.CreateTaskOutputData {
	return &port.CreateTaskOutputData{
		Task:  task,
		Error: err,
	}
}
