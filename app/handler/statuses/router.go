package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	sr repository.Status
}

// Create Handler for `/v1/statuses/`
func NewRouter(sr repository.Status, ar repository.Account) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr}
	r.Route("/", func(r chi.Router) {
		r.Use(auth.Middleware(ar))
		r.Post("/", h.Create)
	})
	r.Get("/{id:[0-9]+}", h.FindStatusByID)
	return r
}
