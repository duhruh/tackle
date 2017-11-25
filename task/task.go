package task

import "io"

type CommandLineArgument interface {
	Key() string
	DefaultValue() interface{}
	Value() interface{}
	Description() string
	SetValue(v interface{})
}

type Task interface {
	ShortDescription() string

	Description() string

	Name() string
	// The options allowed for this task
	Options() []Option
	// The arguments allowed for this task
	Arguments() []Argument
	// actually runs the task
	Run(w io.Writer)
}
