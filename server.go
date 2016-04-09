package gtlws

import (
	"io"
	"net/http"
	"regexp"
)

type Handler struct {
	routes Routes
}

func (han *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	body, _ := han.routes[0].Handler([]string{"Karova"})
	io.WriteString(res, body)
}

func Start(routes Routes) {
	regRoutes := make([]RegRoute,len(routes))
	for indx, route := range routes {
		regRoutes[indx] = RegRoute{
			regexp.MustCompile(route.String),
			route.Handler,
		}
	}
	http.ListenAndServe(":8080", &Handler{routes})
}
