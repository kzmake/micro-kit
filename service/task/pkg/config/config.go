package config

import (
	// use tag
	_ "github.com/kelseyhightower/envconfig"
)

// Config は設定に関する定義です。
type Config struct {
	Endpoint string `default:"0.0.0.0:3000"`
	Database struct {
		Driver string `default:"mysql"`
		DSN    string `default:"test:test@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"`
	}
	ServiceDiscovery struct {
		Endpoint string `default:"127.0.0.1:2379"`
	} `envconfig:"SERVICE_DISCOVERY"`
	Log struct {
		Out   string `default:"stdout"`
		Level string `default:"info"`
	}
}
