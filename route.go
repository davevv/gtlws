package gtlws

import (
	"regexp"
	"html/template"
)

type Page struct {
	Regexp  *regexp.Regexp
	Template *template.Template
	Handler *func(*Page, []string) (string, []string)
}

type PageParams struct {
	RString string
	TPath string
	Handler func(*Page, []string) (string, []string)
}

func MakePage (params PageParams) *Page{
	if params.TPath != "" {
		tpl, err := template.ParseFiles(params.TPath)
		if err != nil {
		}
		return &Page{
			Regexp:regexp.MustCompile(params.RString),
			Template:tpl,
			Handler:&params.Handler,
		}
	}
	return &Page{
		Regexp:regexp.MustCompile(params.RString),
		Handler:&params.Handler,
	}
}

/*
package gtlws

import (
	"regexp"
)

type Routes []Route

type Route struct {
	String  string
	Handler func([]string) (string, []string)
}

type RegRoute struct {
	Regexp  *regexp.Regexp
	Handler func([]string) (string, []string)
}
*/
