package task

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type GenerateTask struct {
	Helpers
	shortDescription string
	description      string
	name             string
	options          []Option
	arguments        []Argument
}

func NewGenerateTask() Task {
	return GenerateTask{
		Helpers:          NewHelpers(),
		name:             "generate:task",
		shortDescription: "Generate a new task",
		description:      "Generates a new task",
		arguments: []Argument{
			NewArgument("task", "the name of the task to generate"),
			NewArgument("directory", "the directory/package to generate the task"),
		},
	}
}

func (t GenerateTask) ShortDescription() string { return t.shortDescription }
func (t GenerateTask) Description() string      { return t.description }
func (t GenerateTask) Name() string             { return t.name }
func (t GenerateTask) Options() []Option        { return t.options }
func (t GenerateTask) Arguments() []Argument    { return t.arguments }

func (t GenerateTask) Run(w io.Writer) {
	newTask := struct {
		Name    string
		Command string
		Package string
	}{}

	taskArg, err := t.GetArgument(t.arguments, "task")
	if err != nil {
		panic(err)
	}

	newTask.Command = taskArg.Value().(string)

	directoryArg, err := t.GetArgument(t.arguments, "directory")
	if err != nil || directoryArg.Value() == nil {
		newTask.Package = "tasks"
	} else {
		newTask.Package = directoryArg.Value().(string)
	}

	newTask.Name = t.classname(newTask.Command)
	fileName := t.filename(newTask.Command)

	tmpl, err := template.New(fileName).Parse(t.template())
	if err != nil {
		panic(err)
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fullPath := filepath.Join(dir, newTask.Package, fileName)
	fmt.Println(fullPath)

	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(file, newTask)
	if err != nil {
		panic(err)
	}
}

func (t GenerateTask) filename(cmd string) string {
	parts := strings.Split(cmd, ":")

	var buf bytes.Buffer

	buf.WriteString(parts[0])
	for _, part := range parts[1:] {
		buf.WriteString("_" + part)
	}
	buf.WriteString(".go")

	return buf.String()

}
func (t GenerateTask) classname(cmd string) string {
	parts := strings.Split(cmd, ":")

	var buf bytes.Buffer

	for _, part := range parts {
		buf.WriteString(strings.Title(part))
	}

	return buf.String()
}

func (t GenerateTask) template() string {
	return `package {{.Package}}

import (
	"io"
	"bytes"

	"github.com/duhruh/tackle/task"
)

type {{.Name}}Task struct {
	task.Helpers
	shortDescription string
	description      string
	name             string
	options          []task.Option
	arguments        []task.Argument
}

func New{{.Name}}Task() task.Task {
	return {{.Name}}Task{
		Helpers:          task.NewHelpers(),
		name:             "{{.Command}}",
		shortDescription: "Short description here",
		description:      "Description here",
		options: []task.Option{},
		arguments: []task.Argument{},
	}
}

func (t {{.Name}}Task) ShortDescription() string   { return t.shortDescription }
func (t {{.Name}}Task) Description() string        { return t.description }
func (t {{.Name}}Task) Name() string               { return t.name }
func (t {{.Name}}Task) Options() []task.Option     { return t.options }
func (t {{.Name}}Task) Arguments() []task.Argument { return t.arguments }

func (t {{.Name}}Task) Run(w io.Writer) {
	w.Write(bytes.NewBufferString("Hello Tasks\n").Bytes())
}
`
}
