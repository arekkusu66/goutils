package serve

import (
	"fmt"
	"net/http"
)

func Dir(dirname string) {
	http.Handle(fmt.Sprintf("%s/", dirname), http.StripPrefix(fmt.Sprintf("%s/", dirname), http.FileServer(http.Dir(fmt.Sprintf(".%s", dirname)))))
}