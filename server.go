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

/*

import (
	"io"
	"net/http"
	"regexp"
)

type Handler struct {
	routes []RegRoute
}

func (han *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	for _, regRoute := range han.routes {
		pars := regRoute.Regexp.FindStringSubmatch(req.URL.Path)
		if pars != nil {
			body, _ := regRoute.Handler(pars)
			io.WriteString(res, body)
			return
		}
	}
	http.NotFound(res, req)
}

func Start(routes Routes) {
	regRoutes := make([]RegRoute,len(routes))
	for indx, route := range routes {
		regRoutes[indx] = RegRoute{
			regexp.MustCompile(route.String),
			route.Handler,
		}
	}
	http.ListenAndServe(":8080", &Handler{regRoutes})
}
*/
