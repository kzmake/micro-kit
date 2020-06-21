package mysql

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/kzmake/micro-kit/service/task/domain/aggregate"
	"github.com/kzmake/micro-kit/service/task/domain/errors"
	"github.com/kzmake/micro-kit/service/task/domain/repository"
	"github.com/kzmake/micro-kit/service/task/domain/vo"
)

const (
	duplicateEntry uint16 = 1062
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
			var mysqlErr *mysql.MySQLError
			if xerrors.As(err, &mysqlErr) {
				if mysqlErr.Number == duplicateEntry {
					return errors.WrapCode(errors.DuplicateTask, err)
				}
			}
			return errors.WrapCode(errors.Unexpected, err)
		}

		return nil
	})
	if err != nil {
		return xerrors.Errorf("トランザクション内で失敗しました: %w", err)
	}

	return nil
}

// List はタスク一覧を取得します。
func (r taskRepository) List(ctx context.Context) ([]*aggregate.Task, error) {
	var resources []Task
	if err := r.db.Find(&resources).Error; err != nil {
		return nil, errors.WrapCode(errors.Unexpected, err)
	}

	tasks := make([]*aggregate.Task, 0, len(resources))
	for _, res := range resources {
		tasks = append(tasks, &aggregate.Task{
			ID:          vo.TaskID(res.ID),
			Description: vo.Description(res.Description),

			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
			DeletedAt: res.DeletedAt,
		})
	}

	return tasks, nil
}

// Find はタスクを取得します。
func (r taskRepository) Find(ctx context.Context, id vo.TaskID) (*aggregate.Task, error) {
	var res Task
	if err := r.db.Where("id = ?", string(id)).First(&res).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.WrapCode(errors.NotFoundTask, err)
		}
		return nil, errors.WrapCode(errors.Unexpected, err)
	}

	task := &aggregate.Task{
		ID:          vo.TaskID(res.ID),
		Description: vo.Description(res.Description),

		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		DeletedAt: res.DeletedAt,
	}

	return task, nil
}

// Delete はタスクを削除します。
func (r taskRepository) Delete(ctx context.Context, task *aggregate.Task) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if !r.isExist(tx, task.ID) {
			return errors.NewCode(errors.NotFoundTask)
		}

		ret := tx.Where("id = ?", string(task.ID)).Delete(&Task{})
		if err := ret.Error; err != nil {
			return errors.WrapCode(errors.Unexpected, err)
		}

		return nil
	})
	if err != nil {
		return xerrors.Errorf("トランザクション内で失敗しました: %w", err)
	}

	return nil
}

func (r taskRepository) isExist(db *gorm.DB, id vo.TaskID) bool {
	return !db.Where("id = ?", string(id)).First(&Task{}).RecordNotFound()
}
