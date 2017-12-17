package main

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi"

	"file-drop/api"
)

func main() {
	flag.Parse()

	r := chi.NewRouter()
	api.LoadRoutes(r)
	http.ListenAndServe(":3000", r)
}
