package task

import "errors"

type Helpers interface {
	GetArgument(args []Argument, key string) (Argument, error)
	GetOption(opts []Option, key string) (Option, error)
}

type helpers struct{}

func NewHelpers() Helpers {
	return helpers{}
}

func (h helpers) GetArgument(args []Argument, key string) (Argument, error) {
	for _, arg := range args {
		if arg.Key() == key {
			return arg, nil
		}
	}

	var c CommandLineArgument
	return c, errors.New("argument not found")
}

func (h helpers) GetOption(opts []Option, key string) (Option, error) {
	for _, arg := range opts {
		if arg.Key() == key {
			return arg, nil
		}
	}

	var c CommandLineArgument
	return c, errors.New("option not found")
}
