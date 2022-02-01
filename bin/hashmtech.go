package main

import (
	"net/http"

	"github.com/kahoon/gopid"
)

func serve() {
	http.Handle("/", http.FileServer(http.Dir("/projects/hashmweb/site/dist/")))
	http.ListenAndServe(":80", nil)
}

func main() {
	gopid.Run("hashmtech")
	serve()
}
