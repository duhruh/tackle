package config

import "io"

type Loader interface {
	LoadFromFile(io.Reader) (Config, error)
}
