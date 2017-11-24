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

	ArgumentsToCommandLine() []CommandLineArgument

	OptionsToCommandLine() []CommandLineArgument
}

type task struct {
	name            string
	description     string
	longDescription string
	options         []Option
	arguments       []Argument
}

func NewTask(name string) Task {
	return task{name: name}
}

func (t task) ShortDescription() string {
	return t.longDescription
}

func (t task) Description() string {
	return t.description
}
func (t task) Name() string {
	return t.name
}
func (t task) Options() []Option {
	return t.options
}
func (t task) Arguments() []Argument {
	return t.arguments
}

func (t task) ArgumentsToCommandLine() []CommandLineArgument {
	var cmd []CommandLineArgument
	for _, c := range t.Arguments() {
		cmd = append(cmd, CommandLineArgument(c))
	}

	return cmd
}

func (t task) OptionsToCommandLine() []CommandLineArgument {
	var cmd []CommandLineArgument
	for _, c := range t.Options() {
		cmd = append(cmd, CommandLineArgument(c))
	}

	return cmd
}
func (t task) Run(w io.Writer) {}
