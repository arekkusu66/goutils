package serve

import (
	"fmt"
	"net/http"
)

//Dir serves the content of a directory to a route that has the same name as it, it must contain the slash /, example: serve.Dir("/scripts", nil)
//
//The mux parameter represents the http requests multiplexer and will default to http.DefaultServeMux if it is nil
func Dir(dirname string, mux *http.ServeMux) {
	var handler *http.ServeMux = http.DefaultServeMux

	if mux != nil {
		handler = mux
	}

	handler.Handle(fmt.Sprintf("%s/", dirname), http.StripPrefix(fmt.Sprintf("%s/", dirname), http.FileServer(http.Dir(fmt.Sprintf(".%s", dirname)))))
}
