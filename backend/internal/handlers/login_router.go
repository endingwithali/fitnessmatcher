package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoginRouter() http.Handler {
	chi := chi.NewRouter()
	// chi.Use(AdminOnly)
	// chi.Get("/", adminIndex)
	// chi.Get("/accounts", adminListAccounts)
	return chi
}
