package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const tempi = `<p>A: {{,A}}</pxp>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(tempi))
	var data struct {
		A string        // Обычный текст
		В template.HTML // HTML
	}
	data.A = "<b>Hello!</b>"
	data.В = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
