package config

import (
//"github.com/go-kit/kit/log/level"
//"github.com/duhruh/tackle"
)

type Config interface {
	OptionMap
	// returns either an option or and array of options
	//HttpBindAddress() string
	//GrpcBindAddress() string
	//LogOption() level.Option
	//Environment() tackle.Environment
	//DatabaseConnection() map[string]string
}
type config struct {
	OptionMap
}

func NewConfig(opt OptionMap) Config {
	return config{OptionMap: opt}
}
