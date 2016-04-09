package gtlws

import (
	"io"
	"net/http"
)

type Handler struct {
	cbk func(res http.ResponseWriter, req *http.Request)
}

func (han *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	han.cbk(res, req)
}

func Start(cbk Page) {
	http.ListenAndServe(":8080", &Handler{func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, cbk())
	}})
}
