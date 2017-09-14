package task

import (
	"bytes"
	"io"
	"sort"
	"unicode"
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
	tasks := t.runner.Tasks()

	sort.Sort(alphabetically(tasks))
	for _, task := range tasks {
		buffer.WriteString("\t" + task.Name() + " - " + task.ShortDescription() + "\n")
	}

	w.Write(buffer.Bytes())
}

type alphabetically []Task

func (s alphabetically) Len() int      { return len(s) }
func (s alphabetically) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s alphabetically) Less(i, j int) bool {
	iRunes := []rune(s[i].Name())
	jRunes := []rune(s[j].Name())

	max := len(iRunes)
	if max > len(jRunes) {
		max = len(jRunes)
	}

	for idx := 0; idx < max; idx++ {
		ir := iRunes[idx]
		jr := jRunes[idx]

		lir := unicode.ToLower(ir)
		ljr := unicode.ToLower(jr)

		if lir != ljr {
			return lir < ljr
		}

		if ir != jr {
			return ir < jr
		}
	}

	return false
}
