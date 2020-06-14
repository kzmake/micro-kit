package config

import (
	"reflect"

	"github.com/kelseyhightower/envconfig"
)

// Config は
type Config interface{}

// New は prefix をもつ環境変数を configType にあわせて config を生成します。
func New(prefix string, configType Config) (interface{}, error) {
	rt := reflect.New(reflect.TypeOf(configType).Elem())
	conf := rt.Interface()

	if err := envconfig.Process(prefix, conf); err != nil {
		return nil, err
	}

	return conf, nil
}
