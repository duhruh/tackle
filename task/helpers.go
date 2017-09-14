package task

import (
	"bytes"
	"errors"
	"io"
)

type Helpers interface {
	GetArgument(args []Argument, key string) (Argument, error)
	GetOption(opts []Option, key string) (Option, error)
	Say(w io.Writer, str string)
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

func (h helpers) Say(w io.Writer, str string) {
	var sentence bytes.Buffer
	sentence.WriteString(str + "\n")
	w.Write(sentence.Bytes())
}
