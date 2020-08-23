package registry

import (
	"golang.org/x/xerrors"

	di "github.com/sarulabs/di/v2"

	"github.com/kzmake/micro-kit/service/task-service/pkg/config"
)

// New は config をもとにDIコンテナを生成します。
func New(cfg *config.Config) (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, xerrors.Errorf("Builder生成に失敗しました: %w", err)
	}

	if err := builder.Add(append(append(cores, outputs...), di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) { return cfg, nil },
	})...); err != nil {
		return nil, xerrors.Errorf("Definitions追加に失敗しました: %w", err)
	}

	return builder.Build(), nil
}
