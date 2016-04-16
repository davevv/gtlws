package gtlws

import (
	"html/template"
)

type templateWriter struct {
	String string
}

func (w *templateWriter) Write(p []byte) (n int, err error) {
	if w.String == "" {
		w.String = string(p)
	} else {
		w.String = w.String+string(p)
	}
	return len(p), nil
}

func ExecTemplate(tpl *template.Template, data interface{}) string {
	writer := &templateWriter{}
	err := tpl.Execute(writer, data)
	if err != nil {
		panic(err)
	}
	return writer.String
}
