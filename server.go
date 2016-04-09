package gtlws

import (
	"io"
	"net/http"
)

type Handler struct {
	routes Routes
}

func (han *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	body, _ := han.routes[0].Handler([]string{"Karova"})
	io.WriteString(res, body)
}

func Start(routes Routes) {
	http.ListenAndServe(":8080", &Handler{routes})
}
