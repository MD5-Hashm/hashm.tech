package main

import (
	"net/http"
)

func serve() {
	http.Handle("/", http.FileServer(http.Dir("/projects/hashmweb/site/dist/")))
	http.ListenAndServe(":80", nil)
}

func main() {
	serve()
}
