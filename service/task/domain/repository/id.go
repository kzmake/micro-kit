package repository

import (
	"context"

	"github.com/kzmake/micro-kit/service/task/domain/vo"
)

// ID はIDに関するリポジトリのIFです。
type ID interface {
	Gen(context.Context) (vo.ID, error)
}
