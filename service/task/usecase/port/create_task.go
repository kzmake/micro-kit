package port

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
)

// CreateTaskInputData はタスク作成のための InputData です。
// DTO (Data Transfer Object) として InputData を生成します。
type CreateTaskInputData struct {
	Description string
}

// CreateTaskOutputData はタスク作成のための OutputData です。
// DPO (Data Payload Object) として OutputData を生成します。
type CreateTaskOutputData struct {
	Task  *aggregate.Task
	Error error
}

// CreateTask はタスク作成のための Port です。
type CreateTask interface {
	Handle(ctx context.Context, in *CreateTaskInputData) *CreateTaskOutputData
}
