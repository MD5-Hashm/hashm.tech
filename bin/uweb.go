package main

import (
	"net/http"

	"github.com/kahoon/gopid"
)

func serve() {
	http.Handle("/", http.FileServer(http.Dir("/projects/hashmweb/site/updating/")))
	http.ListenAndServe(":80", nil)
}

func main() {
	gopid.Run("uweb")
	serve()
}
