package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	// "github.com/rizqdwan/e-wallet/internal/service"
)


func NewRouter() http.Handler {

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})
	})


	return r
}


