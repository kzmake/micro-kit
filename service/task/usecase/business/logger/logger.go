package logger

import (
	"context"
	"io"
	"os"

	"github.com/kzmake/micro-kit/pkg/logger"

	"github.com/kzmake/micro-kit/service/task/usecase/business"
)

// New はビジネスロジックの Success / Failure を記憶する Logger を生成します。
func New(w io.Writer) business.Assistant {
	log := logger.New(logger.WithOutput(os.Stdout))

	return business.Assistant(func(nextFn business.Task) business.Task {
		return func(ctx context.Context) (interface{}, error) {
			v, err := nextFn(ctx)
			if err != nil {
				log.Errorf("Failed: %+v", err)
			}
			log.Infof("Success: %+v", v)
			return v, err
		}
	})
}
