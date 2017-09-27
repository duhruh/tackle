package tackle

import "github.com/go-kit/kit/log/level"

type Config interface {
	HttpBindAddress() string
	GrpcBindAddress() string
	LogOption() level.Option
	Environment() Environment
	DatabaseConnection() map[string]string
}
