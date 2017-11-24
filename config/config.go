package config

type Config interface {
	OptionMap
}
type config struct {
	OptionMap
}

func NewConfig(opt OptionMap) Config {
	return config{OptionMap: opt}
}
