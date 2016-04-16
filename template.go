package gtlws

import (
	"html/template"
)

type templateWriter struct {
	String string
}

func (w *templateWriter) Write(p []byte) (n int, err error) {
	w.String = string(p)
	return len(p), nil
}

func ExecTemplate(tpl *template.Template, data interface{}) string {
	writer := &templateWriter{}
	tpl.Execute(writer, nil)
	return writer.String
}
