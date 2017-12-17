package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// LoadRoutes configures a Router with routes
func LoadRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
}
