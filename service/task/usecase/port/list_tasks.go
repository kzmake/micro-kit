package port

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
)

// ListTasksInputData はタスク一覧取得のための InputData です。
// DTO (Data Transfer Object) として InputData を生成します。
type ListTasksInputData struct {
}

// ListTasksOutputData はタスク一覧取得のための OutputData です。
// DPO (Data Payload Object) として OutputData を生成します。
type ListTasksOutputData struct {
	Tasks []*aggregate.Task
	Error error
}

// ListTasks はタスク一覧取得のための Port です。
type ListTasks interface {
	Handle(ctx context.Context, in *ListTasksInputData) *ListTasksOutputData
}
