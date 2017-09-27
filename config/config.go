package config

import (
//"github.com/go-kit/kit/log/level"
//"github.com/duhruh/tackle"
)

type Config interface {
	// returns either an option or and array of options
	Get(string) interface{}
	//HttpBindAddress() string
	//GrpcBindAddress() string
	//LogOption() level.Option
	//Environment() tackle.Environment
	//DatabaseConnection() map[string]string
}
type config struct {
	options []Option
}

func NewConfig(opt []Option) Config {
	return config{options: opt}
}

func (c config) Get(opt string) interface{} {
	for _, option := range c.options {
		if option.Name() == opt {
			return option
		}
	}
	return nil
}

type Option interface {
	Name() string
	Value() interface{}
}
