package gtlws

import (
	"html/template"
)

var templates = map[string]*template.Template{}

type templateWriter struct {
	String string
}

func (hw *templateWriter) Write(p []byte) (n int, err error){
	hw.String = string(p)
	return len(p), nil
}

func ExecTemplateFile(path string) string {
	tpl, ok := templates[path]
	if ok == false {
		tplTwo, err := template.ParseFiles(path)
		if err != nil{
			return "error"
		}
		templates[path] = tplTwo
		tpl = tplTwo
	}
	writer := &templateWriter{}
	tpl.Execute(writer, nil)
	return writer.String
}
