package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func main() {
	str := `
{{.ID}}
{{.Name}}
{{.Age}}
{{.Gender}}
`

	u := struct {
		ID     int
		Name   string
		Age    int
		Gender int
	}{
		ID:     1,
		Name:   "yukpiz",
		Age:    1024,
		Gender: 1,
	}

	embeded, err := Embed(str, u, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("------------------")
	fmt.Println(embeded)
	fmt.Println("------------------")
}

func Embed(tmplstr string, data interface{}, trim bool) (string, error) {
	if trim {
		tmplstr = strings.TrimSpace(tmplstr)
	}
	tmpl, err := template.New("example").Parse(tmplstr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
