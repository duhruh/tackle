package task

import (
	"bytes"
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

// Register a new task to be run
func (r *runner) Register(task Task) {
	r.tasks[task.Name()] = task
}

// Sets the task to run by default
func (r *runner) SetDefaultTask(task Task) {
	r.defaultTask = task
}

// Returns a list of all the tasks currently registered
func (r *runner) Tasks() []Task {
	var values []Task
	for _, val := range r.tasks {
		values = append(values, val)
	}
	return values
}

// Kicks off the task runner
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
			if r.isHelp(arg) {
				r.showCommandHelp(task)
				return
			}

			r.populateOption(arg, task)
			continue
		}

		r.populateArgument(arg, task)
	}

	task.Run(r.writer)
}

// Simple check to see if the argument is the help option
func (r *runner) isHelp(arg string) bool {
	return strings.Contains(arg, "help")
}

// Prints out the complete help for a give task
func (r *runner) showCommandHelp(task Task) {
	var fullHelp bytes.Buffer

	fullHelp.WriteString("command: " + task.Name() + "\n")
	fullHelp.WriteString("\t" + task.Description() + "\n")

	if len(task.Options()) > 0 {
		fullHelp.WriteString("options\n")
		for _, opt := range task.Options() {
			fullHelp.WriteString("\t" + opt.Key() + " - " + opt.Description() + "\n")
		}
	}

	if len(task.Arguments()) > 0 {
		fullHelp.WriteString("arguments\n")
		for _, arg := range task.Arguments() {
			fullHelp.WriteString("\t" + arg.Key() + " - " + arg.Description() + "\n")
		}
	}

	r.writer.Write(fullHelp.Bytes())
}

// Adds the given raw argument to the running list of found arguments
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
