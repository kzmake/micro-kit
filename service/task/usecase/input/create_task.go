package input

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/usecase/output"
)

// CreateTaskData はタスク作成のための InputData です。
// DTO (Data Transfer Object) として InputData を生成します。
type CreateTaskData struct {
	Description string
}

// CreateTaskPort はタスク作成のための InputPort です。
type CreateTaskPort interface {
	Handle(ctx context.Context, input *CreateTaskData) *output.CreateTaskData
}
