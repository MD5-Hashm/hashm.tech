package main

import (
	"net/http"

	"github.com/kahoon/gopid"
)

func serve() {
	http.Handle("/", http.FileServer(http.Dir("/projects/hashmweb/site/dist/")))
	http.Handle("/slope", http.FileServer(http.Dir("/projects/hashmweb/site/slope/slope-game/")))
	http.ListenAndServe(":80", nil)
}

func main() {
	gopid.Run("hashmtech")
	serve()
}
