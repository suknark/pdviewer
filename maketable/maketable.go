package maketable

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"pdviewer/tools"
)

func MakeTempl(i []tools.Inc, sw string, oncall string) {
	content, err := ioutil.ReadFile("templates/table.html")
	if err != nil {
		log.Fatal("Could not read file")
		return
	}

	t := template.Must(template.New("tmpl").Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j == 0 }}).Parse(string(content)))

	var html bytes.Buffer
	err = t.Execute(&html, i)
	if err != nil {
		log.Println("executing template:", err)
	}
	maketemplate := "{{define \"incidents\"}}\n <h1 class=\"header\">Time without high incidents - " + sw + "</h1><h2 class=\"header\"> OnCall now - " + oncall + "</h2> <div class=\"row\">\n" + html.String() + "{{end}}"
	d1 := []byte(maketemplate)
	ioutil.WriteFile("templates/post.html", d1, 0644)
}
