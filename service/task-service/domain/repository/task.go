package repository

import (
	"context"

	"github.com/kzmake/micro-kit/service/task-service/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task-service/domain/vo"
)

// Task はタスクに関するリポジトリのIFです。
type Task interface {
	Save(context.Context, *aggregate.Task) error
	Find(context.Context, vo.TaskID) (*aggregate.Task, error)
	List(context.Context) ([]*aggregate.Task, error)
	Delete(context.Context, *aggregate.Task) error
}
