package buyers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandler(buyerService Service) chi.Router {

	router := chi.NewRouter()

	router.Get("/buyer", listBuyers(buyerService))
	router.Get("/buyer/{buyerId}", getBuyerDetails(buyerService))

	return router
}

func listBuyers(s Service) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		buyers, err := s.FindAllBuyers()

		for idx, _ := range buyers {
			buyers[idx].UId = ""
		}
		if err != nil {

			http.Error(rw, err.Error(), http.StatusInternalServerError)

		}

		json.NewEncoder(rw).Encode(buyers)
	}

}

func getBuyerDetails(s Service) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		buyerId := chi.URLParam(r, "buyerId")

		buyerInfo, err := s.FindBuyerById(buyerId)

		for idx, _ := range buyerInfo.BuyersEqIp {
			buyerInfo.BuyersEqIp[idx].Buyer.UId = ""
		}
		if err != nil {

			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		json.NewEncoder(rw).Encode(buyerInfo)
	}
}
