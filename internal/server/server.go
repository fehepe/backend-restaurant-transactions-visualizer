package server

import (
	"net/http"
	"os"

	"github.com/fehepe/backend-restaurant-transactions-visualizer/internal/buyers"
	"github.com/fehepe/backend-restaurant-transactions-visualizer/internal/loaddata"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router       *chi.Mux
	BuyerService buyers.Service
	LoadService  loaddata.Service
}

func Run(buyerService buyers.Service, loadService loaddata.Service) error {

	server := &Server{
		Router:       chi.NewRouter(),
		BuyerService: buyerService,
		LoadService:  loadService,
	}
	port := os.Getenv("API_PORT")

	server.initRoutes()

	return http.ListenAndServe(port, server.Router)
}

func (s *Server) initRoutes() {
	s.Router.Use(
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.SetHeader("Content-Type", "application/json"),
	)

	s.Router.Get("/buyer", buyers.ListBuyers(s.BuyerService))
	s.Router.Get("/buyer/{buyerId}", buyers.GetBuyerDetails(s.BuyerService))
	s.Router.Post("/load", loaddata.LoadData(s.LoadService))
	s.Router.Post("/load/{date}", loaddata.LoadData(s.LoadService))
}
