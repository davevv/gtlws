package gtlws

import (
	"io"
	"net/http"
)

func Start(cbk Page) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, cbk())
	})
	http.ListenAndServe(":8080", nil)
}
