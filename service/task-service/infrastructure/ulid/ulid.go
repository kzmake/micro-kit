package ulid

import (
	"context"
	"math/rand"
	"time"

	ulid "github.com/oklog/ulid/v2"

	"github.com/kzmake/micro-kit/service/task-service/domain/repository"
	"github.com/kzmake/micro-kit/service/task-service/domain/vo"
)

type idRepository struct{}

// interfaces
var _ repository.ID = (*idRepository)(nil)

// NewIDRepository はIDに関するリポジトリを生成します。
func NewIDRepository() repository.ID { return &idRepository{} }

// Gen は ID を生成します。
func (r *idRepository) Gen(ctx context.Context) vo.ID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()

	return vo.ID(id)
}
