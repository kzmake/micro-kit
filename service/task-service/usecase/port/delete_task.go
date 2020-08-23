package port

import (
	"context"
)

// DeleteTaskInputData はタスク削除のための InputData です。
// DTO (Data Transfer Object) として InputData を生成します。
type DeleteTaskInputData struct {
	ID string
}

// DeleteTaskOutputData はタスク削除のための OutputData です。
// DPO (Data Payload Object) として OutputData を生成します。
type DeleteTaskOutputData struct {
	Error error
}

// DeleteTask はタスク作成のための Port です。
type DeleteTask interface {
	Handle(ctx context.Context, in *DeleteTaskInputData) *DeleteTaskOutputData
}
