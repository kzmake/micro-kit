package repository

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/domain/vo"
)

// IDRepository はIDに関するリポジトリのIFです。
type IDRepository interface {
	Gen(context.Context) (vo.ID, error)
}
