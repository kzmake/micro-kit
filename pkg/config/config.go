package config

import (
	"golang.org/x/xerrors"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"

	"github.com/kzmake/micro-kit/pkg/logger/technical"
)

// Apply は 指定したソースを Config に設定します。
func Apply(cfg interface{}, sources ...source.Source) error {
	conf, err := config.NewConfig()
	if err != nil {
		return xerrors.Errorf("Configの生成に失敗しました: %w", err)
	}

	_ = conf.Load(sources...)

	technical.Infof("Configを読み込みました: %#v", conf.Map())

	_ = conf.Scan(cfg)

	w, err := conf.Watch()
	if err != nil {
		return xerrors.Errorf("Watcherの生成に失敗しました: %w", err)
	}

	go hotReload(w, cfg)

	return nil
}

func hotReload(w config.Watcher, cfg interface{}) {
	for {
		v, err := w.Next()
		if err != nil {
			technical.Warnf("HotReloadに失敗しました: %+v", err)
			continue
		}

		_ = v.Scan(cfg)
		technical.Warnf("HotReloadに成功しました: %+v", cfg)
	}
}
