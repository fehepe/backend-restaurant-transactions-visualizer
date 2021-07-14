package server

import (
	"backend-restaurant-transactions-visualizer/internal/buyers"
	"backend-restaurant-transactions-visualizer/internal/loaddata"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run(buyerService buyers.Service, loadService loaddata.Service) error {
	port := os.Getenv("API_PORT")
	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.SetHeader("Content-Type", "application/json"),
	)
	router.Mount("/buyer", buyers.NewHandler(buyerService))
	router.Mount("/load", loaddata.NewHandler(loadService))

	return http.ListenAndServe(port, router)
}
