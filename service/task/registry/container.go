package registry

import (
	"golang.org/x/xerrors"

	di "github.com/sarulabs/di/v2"
)

// New は defs をもとにDIコンテナを生成します。
func New(defs ...di.Def) (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, xerrors.Errorf("Builder生成に失敗しました: %w", err)
	}

	if err := builder.Add(defs...); err != nil {
		return nil, xerrors.Errorf("Definitions追加に失敗しました: %w", err)
	}

	return builder.Build(), nil
}
