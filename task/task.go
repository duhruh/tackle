package task

import "io"

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
func (t task) Run(w io.Writer) {}

type CommandLineArgument interface {
	Key() string
	DefaultValue() interface{}
	Value() interface{}
	Description() string
	SetValue(v interface{})
}
type Argument interface {
	CommandLineArgument
}

type argument struct {
	key          string
	defaultValue string
	value        interface{}
	description  string
}

func NewArgument(key string, description string, ops ...interface{}) Argument {
	return &argument{
		key:         key,
		description: description,
	}
}

func (o *argument) SetValue(v interface{}) {
	o.value = v
}
func (o *argument) Key() string {
	return o.key
}
func (o *argument) DefaultValue() interface{} {
	return o.defaultValue
}
func (o *argument) Value() interface{} {
	return o.value
}
func (o *argument) Description() string {
	return o.description
}

type Option interface {
	CommandLineArgument
}

type option struct {
	key          string
	defaultValue string
	value        interface{}
	description  string
}

func NewOption(key string, description string, ops ...interface{}) Option {
	return &option{
		key:         key,
		description: description,
	}
}

func (o *option) Key() string {
	return o.key
}
func (o *option) DefaultValue() interface{} {
	return o.defaultValue
}
func (o *option) Value() interface{} {
	return o.value
}
func (o *option) Description() string {
	return o.description
}
func (o *option) SetValue(v interface{}) {
	o.value = v
}
