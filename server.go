package gtlws

import (
	"io"
	"net/http"
)

type Pages []*Page

func (pages *Pages) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	for _, page := range *pages {
		pars := page.Regexp.FindStringSubmatch(req.URL.Path)
		if pars != nil {
			f := *(page.Handler)
			body, _ := f(page, pars)
			io.WriteString(res, body)
			return
		}
	}
	http.NotFound(res, req)
}

func Start(pages Pages) {
	http.ListenAndServe(":8080", &pages)
}
