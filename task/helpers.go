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
	arg, err := h.get([]CommandLineArgument(args), key)

	if err != nil {
		return nil, err
	}

	return arg.(Argument), nil
}

func (h helpers) GetOption(opts []Option, key string) (Option, error) {
	arg, err := h.get([]CommandLineArgument(opts), key)

	if err != nil {
		return nil, err
	}

	return arg.(Option), nil
}

func (h helpers) get(args []CommandLineArgument, key string) (CommandLineArgument, error) {
	for _, arg := range args {
		if arg.Key() == key {
			return arg, nil
		}
	}

	var c CommandLineArgument
	return c, errors.New("argument not found")
}
