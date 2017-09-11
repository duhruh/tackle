package task

import (
	"bytes"
	"io"
	"os"
	"text/template"
)

type GenerateTask struct {
	shortDescription string
	description      string
	name             string
	options          []Option
	arguments        []Argument
}

func NewGenerateTask() Task {
	return GenerateTask{
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

	//var directory string

	newTask := struct {
		Name    string
		Command string
		Package string
	}{}

	for _, val := range t.arguments {
		if val.Key() == "task" {
			newTask.Name = val.Value().(string)
		}
		if val.Key() == "directory" {
			//_ := val.Value().(string)
		}
	}

	tmpl, err := template.New("task").Parse(t.template())
	if err != nil {
		// do something
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, newTask)
	if err != nil {
		// do something
		panic(err)
	}

	w.Write(bytes.NewBufferString("Hello Tasks").Bytes())
}

func (t GenerateTask) template() string {
	return `
package {{.Package}}

import (
	"io"
	"bytes"
	"github.com/duhruh/tackle/task"
)

type {{.Name}}Task struct {
	shortDescription string
	description      string
	name             string
	options          []task.Option
	arguments        []task.Argument
}

func New{{.Name}}Task() Task {
	return {{.Name}}Task{
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
	w.Write(bytes.NewBufferString("Hello Tasks").Bytes())
}
`
}
