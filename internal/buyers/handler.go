package buyers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ListBuyers(s Service) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		buyers, err := s.FindAllBuyers()

		for idx := range buyers {
			buyers[idx].UId = ""
		}
		if err != nil {

			http.Error(rw, err.Error(), http.StatusInternalServerError)

		}

		json.NewEncoder(rw).Encode(buyers)
	}

}

func GetBuyerDetails(s Service) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		buyerId := chi.URLParam(r, "buyerId")

		buyerInfo, err := s.FindBuyerById(buyerId)

		for idx := range buyerInfo.BuyersEqIp {
			buyerInfo.BuyersEqIp[idx].Buyer.UId = ""
		}
		if err != nil {

			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		json.NewEncoder(rw).Encode(buyerInfo)
	}
}
