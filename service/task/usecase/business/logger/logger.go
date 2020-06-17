package logger

import (
	"context"

	"github.com/kzmake/micro-kit/pkg/logger"

	"github.com/kzmake/micro-kit/service/task/usecase/business"
)

// New はビジネスロジックの Success / Failure を記憶する Logger を生成します。
func New(l *logger.Logger) business.Assistant {
	return business.Assistant(func(nextFn business.Task) business.Task {
		return func(ctx context.Context) (interface{}, error) {
			v, err := nextFn(ctx)
			if err != nil {
				l.Errorf("Failed: %+v", err)
			}
			l.Infof("Success: %+v", v)
			return v, err
		}
	})
}
