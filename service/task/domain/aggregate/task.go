package aggregate

import (
	"time"

	"github.com/kzmake/micro-kit/service/task/domain/vo"
)

// Task は aggregate root の定義です。
type Task struct {
	ID          vo.ID
	Description vo.Description

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
