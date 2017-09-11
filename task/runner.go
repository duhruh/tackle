package task

import (
	"io"
	"strings"
)

type Runner interface {
	Register(task Task)
	Run(args []string)
	Tasks() []Task
	SetDefaultTask(task Task)
}

type runner struct {
	tasks          map[string]Task
	writer         io.Writer
	defaultTask    Task
	foundArguments map[string]Argument
}

func NewRunner(writer io.Writer, opts ...interface{}) Runner {
	r := &runner{writer: writer, tasks: make(map[string]Task), foundArguments: make(map[string]Argument)}
	r.SetDefaultTask(NewDefaultTask(r))
	r.Register(NewGenerateTask())
	return r
}

func (r *runner) Register(task Task) {
	r.tasks[task.Name()] = task
}

func (r *runner) SetDefaultTask(task Task) {
	r.defaultTask = task
}

func (r *runner) Tasks() []Task {
	var values []Task
	for _, val := range r.tasks {
		values = append(values, val)
	}
	return values
}

func (r *runner) Run(args []string) {
	if len(args) == 1 {
		r.defaultTask.Run(r.writer)
		return
	}

	command := args[1]

	task, ok := r.tasks[command]

	if !ok {
		r.defaultTask.Run(r.writer)
		return
	}

	for _, arg := range args[2:] {
		if strings.HasPrefix(arg, "--") {

			r.populateOption(arg, task)
			continue
		}

		r.populateArgument(arg, task)
	}

	task.Run(r.writer)
}

func (r *runner) populateArgument(raw string, task Task) {
	for _, arg := range task.Arguments() {
		if _, ok := r.foundArguments[arg.Key()]; ok {
			continue
		}

		arg.SetValue(raw)
		r.foundArguments[arg.Key()] = arg
		break
	}
}

func (r *runner) populateOption(raw string, task Task) {
	parts := strings.SplitN(raw, "=", 2)
	option := strings.TrimPrefix(parts[0], "--")
	for _, opt := range task.Options() {
		if opt.Key() == option {
			opt.SetValue(parts[1])
			break
		}
	}
}
