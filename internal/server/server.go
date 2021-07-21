package server

import (
	"github.com/fehepe/backend-restaurant-transactions-visualizer/internal/buyers"
	"github.com/fehepe/backend-restaurant-transactions-visualizer/internal/loaddata"
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

	router.Get("/buyer", buyers.ListBuyers(buyerService))
	router.Get("/buyer/{buyerId}", buyers.GetBuyerDetails(buyerService))
	router.Post("/load", loaddata.LoadData(loadService))
	router.Post("/load/{date}", loaddata.LoadData(loadService))

	return http.ListenAndServe(port, router)
}
