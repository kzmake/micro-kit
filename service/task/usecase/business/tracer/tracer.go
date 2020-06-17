package tracer

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/kzmake/micro-kit/service/task/usecase/business"
)

// New はビジネスロジックの処理を追跡する Tracer を生成します。
func New() business.Assistant {
	return business.Assistant(func(nextFn business.Task) business.Task {
		return func(ctx context.Context) (interface{}, error) {
			span, ctx := opentracing.StartSpanFromContext(ctx, "interactor")
			defer span.Finish()

			v, err := nextFn(ctx)
			if err != nil {
				ext.Error.Set(span, true)
				return nil, err
			}

			return v, nil
		}
	})
}
