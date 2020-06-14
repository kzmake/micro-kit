package mysql

import (
	"context"

	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task/domain/repository"
	"github.com/kzmake/micro-kit/service/task/domain/vo"
)

// Task はタスクに関するスキーマです。
type Task struct {
	ID          string `gorm:"primary_key;size:26"`
	Description string `gorm:"not null;size:255"`

	BaseSchema
}

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository はタスクに関するリポジトリを生成します。
func NewTaskRepository(db *gorm.DB) repository.Task {
	return taskRepository{db}
}

// Save はタスクを永続化します。
func (r taskRepository) Save(ctx context.Context, task *aggregate.Task) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Task{
			ID:          string(task.ID),
			Description: string(task.Description),
		}).Error; err != nil {
			return xerrors.Errorf("Saveに失敗しました: %w", err)
		}

		return nil
	})
	if err != nil {
		return xerrors.Errorf("トランザクション内で失敗しました: %w", err)
	}

	return nil
}

// Find はタスクを取得します。
func (r taskRepository) Find(ctx context.Context, id vo.TaskID) (*aggregate.Task, error) {
	var res Task
	if err := r.db.Where("id = ?", id).First(&res); err != nil {
		return nil, xerrors.Errorf("Findに失敗しました: %w", err)
	}

	task := &aggregate.Task{
		ID:          vo.TaskID(res.ID),
		Description: vo.Description(res.Description),
	}

	return task, nil
}
