package output

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
)

// CreateTaskData はタスク作成のための OutputData です。
// DPO (Data Payload Object) として OutputData を生成します。
type CreateTaskData struct {
	Task  *aggregate.Task
	Error error
}

// CreateTaskPort はタスク作成のための OutputPort です。
type CreateTaskPort interface {
	Handle(ctx context.Context, task *aggregate.Task, err error) *CreateTaskData
}
