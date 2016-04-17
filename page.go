package gtlws

import (
	"html/template"
	"regexp"
	"net/http"
)

type PageHandler func(*Page, []string, *http.ResponseWriter)

type Page struct {
	Regexp   *regexp.Regexp
	Template *template.Template
	Handler  *PageHandler
}

type PageParams struct {
	RString string
	TPath   string
	Handler PageHandler
}

func MakePage(params PageParams) *Page {
	if params.TPath != "" {
		tpl, err := template.ParseFiles(params.TPath)
		if err != nil {
			panic(err)
		}
		return &Page{
			Regexp:   regexp.MustCompile(params.RString),
			Template: tpl,
			Handler:  &params.Handler,
		}
	}
	return &Page{
		Regexp:  regexp.MustCompile(params.RString),
		Handler: &params.Handler,
	}
}
