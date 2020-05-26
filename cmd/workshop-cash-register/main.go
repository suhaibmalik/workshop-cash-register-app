package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/suhaibmalik/workshop-cash-register-app/internal/cart"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/list", cart.HandleListItems)
	// todo: add the other endpoints

	http.ListenAndServe(":3000", r)
}
