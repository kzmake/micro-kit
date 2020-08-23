package config

import (
	"golang.org/x/xerrors"

	"github.com/micro/go-micro/v2/config/source/env"
	"github.com/micro/go-micro/v2/config/source/etcd"

	pconfig "github.com/kzmake/micro-kit/pkg/config"
)

const envPrefix = "TASK"

// Config は設定に関する定義です。
type Config struct {
	Database struct {
		Driver string
		DSN    string
	}
	Log struct {
		Out   string
		Level string
	}
	DynamicConfig struct {
		Provider string
		Endpoint string
		Prefix   string
	}
}

func newDefaultConfig() *Config {
	cfg := &Config{}

	cfg.Database.Driver = "mysql"
	cfg.Database.DSN = "test:test@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"

	cfg.Log.Out = "stdout"
	cfg.Log.Level = "info"

	return cfg
}

// New は Config を生成します。
func New() (*Config, error) {
	cfg := newDefaultConfig()

	// static: env
	err := pconfig.Apply(cfg, env.NewSource(
		env.WithPrefix(envPrefix),
		env.WithStrippedPrefix(envPrefix),
	))
	if err != nil {
		return nil, xerrors.Errorf("静的コンフィグの適用に失敗しました: %w", err)
	}

	// dynamic: etcd
	if cfg.DynamicConfig.Provider == "etcd" {
		err = pconfig.Apply(cfg, etcd.NewSource(
			etcd.WithAddress(cfg.DynamicConfig.Endpoint),
			etcd.WithPrefix(cfg.DynamicConfig.Prefix),
			etcd.StripPrefix(true),
		))
		if err != nil {
			return nil, xerrors.Errorf("動的コンフィグの適用に失敗しました: %w", err)
		}
	}

	return cfg, nil
}
