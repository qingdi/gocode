package ch4

import (
	"fmt"
	"html/template"
	"os"
)

func Autoescape() {
	const templ = `<p>A:{{.A}}</p> <p>B:{{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string
		B template.HTML
	}

	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		fmt.Println("err")
	}
}
