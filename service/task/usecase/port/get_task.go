package port

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
)

// GetTaskInputData はタスク取得のための InputData です。
// DTO (Data Transfer Object) として InputData を生成します。
type GetTaskInputData struct {
	ID string
}

// GetTaskOutputData はタスク取得のための OutputData です。
// DPO (Data Payload Object) として OutputData を生成します。
type GetTaskOutputData struct {
	Task  *aggregate.Task
	Error error
}

// GetTask はタスク取得のための Port です。
type GetTask interface {
	Handle(ctx context.Context, in *GetTaskInputData) *GetTaskOutputData
}
