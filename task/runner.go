package task

import (
	"bytes"
	"io"
	"strings"
)

const (
	helpOption       = "help"
	optionPrefix     = "--"
	optionAssignment = "="

	optionText   = "options"
	commandText  = "commands"
	argumentText = "arguments"

	newline = "\n"
	tab     = "\t"
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

	commandArgs := args[2:]
	for _, arg := range commandArgs {
		if r.isOption(arg) {
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
	return strings.Contains(arg, helpOption)
}

func (r *runner) isOption(arg string) bool {
	return strings.HasPrefix(arg, optionPrefix)
}

// Prints out the complete help for a give task
func (r *runner) showCommandHelp(task Task) {
	var fullHelp bytes.Buffer

	fullHelp.WriteString(commandText + ": " + task.Name() + newline)
	fullHelp.WriteString(tab + task.Description() + newline)

	if len(task.Options()) > 0 {
		fullHelp.WriteString(optionText + newline)
		fullHelp = r.formatCommandLine(fullHelp, r.optionsToCommandLine(task.Options()))
	}

	if len(task.Arguments()) > 0 {
		fullHelp.WriteString(argumentText + newline)
		fullHelp = r.formatCommandLine(fullHelp, r.argumentsToCommandLine(task.Arguments()))
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
	parts := strings.SplitN(raw, optionAssignment, 2)
	option := strings.TrimPrefix(parts[0], optionPrefix)
	for _, opt := range task.Options() {
		if opt.Key() == option {
			opt.SetValue(parts[1])
			break
		}
	}
}

func (r *runner) formatCommandLine(buf bytes.Buffer, cmd []CommandLineArgument) bytes.Buffer {
	for _, arg := range cmd {
		buf.WriteString(tab + arg.Key() + " - " + arg.Description() + newline)
	}
	return buf
}

func (r *runner) argumentsToCommandLine(args []Argument) []CommandLineArgument {
	var cmd []CommandLineArgument
	for _, c := range args {
		cmd = append(cmd, CommandLineArgument(c))
	}

	return cmd
}

func (r *runner) optionsToCommandLine(opts []Option) []CommandLineArgument {
	var cmd []CommandLineArgument
	for _, c := range opts {
		cmd = append(cmd, CommandLineArgument(c))
	}

	return cmd
}
