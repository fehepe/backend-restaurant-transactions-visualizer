package server

import (
	"backend-restaurant-transactions-visualizer/internal/buyers"
	"backend-restaurant-transactions-visualizer/internal/loaddata"
	"log"
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
		SetJsonResponseContentType,
	)

	router.Get("/buyer", buyers.ListBuyers(buyerService))
	router.Get("/buyer/{buyerId}", buyers.GetBuyerDetails(buyerService))
	router.Post("/load", loaddata.LoadData(loadService))
	router.Post("/load/{date}", loaddata.LoadData(loadService))

	log.Fatal(http.ListenAndServe(port, router))
	log.Printf("Starting server on: http://localhost%s/api/\n", port)
	return nil
}

func SetJsonResponseContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
