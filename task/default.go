package task

import (
	"bytes"
	"io"
)

type DefaultTask struct {
	runner Runner
}

func NewDefaultTask(runner Runner) Task {
	return DefaultTask{
		runner: runner,
	}
}

func (t DefaultTask) ShortDescription() string {
	return "Tackle default task"
}

func (t DefaultTask) Description() string {
	return "Prints out help information on all tasks"
}
func (t DefaultTask) Name() string {
	return "help"
}
func (t DefaultTask) Options() []Option {
	return []Option{}
}
func (t DefaultTask) Arguments() []Argument {
	return []Argument{}
}
func (t DefaultTask) Run(w io.Writer) {
	var buffer bytes.Buffer

	buffer.WriteString("\nWelcome to Tackle v1.0.0\n\n")

	buffer.WriteString("Tasks\n")
	for _, task := range t.runner.Tasks() {
		buffer.WriteString("\t" + task.Name() + " - " + task.ShortDescription() + "\n")
	}

	w.Write(buffer.Bytes())
}
