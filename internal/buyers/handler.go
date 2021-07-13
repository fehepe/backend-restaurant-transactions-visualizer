package buyers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ListBuyers(s Service) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		buyers, err := s.FindAllBuyers()

		for idx, _ := range buyers {
			buyers[idx].UId = ""
		}
		if err != nil {

			json.NewEncoder(rw).Encode(err.Error())
			return
		}

		json.NewEncoder(rw).Encode(buyers)
	}

}

func GetBuyerDetails(s Service) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		buyerId := chi.URLParam(r, "buyerId")

		if buyerId == "" {

			json.NewEncoder(rw).Encode("BuyerId cannot be empty.")
		}

		buyerInfo, err := s.FindBuyerById(buyerId)

		for idx, _ := range buyerInfo.BuyersEqIp {
			buyerInfo.BuyersEqIp[idx].Buyer.UId = ""
		}
		if err != nil {

			json.NewEncoder(rw).Encode(err)
			return
		}

		json.NewEncoder(rw).Encode(buyerInfo)
	}
}
