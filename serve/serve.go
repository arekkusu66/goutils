package serve

import (
	"fmt"
	"net/http"
)

//Dir serves the content of a directory to a route that has the same name as it, it must contain the slash /, example: serve.Dir("/scripts")
func Dir(dirname string) {
	http.Handle(fmt.Sprintf("%s/", dirname), http.StripPrefix(fmt.Sprintf("%s/", dirname), http.FileServer(http.Dir(fmt.Sprintf(".%s", dirname)))))
}