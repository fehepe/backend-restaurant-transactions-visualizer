package loaddata

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoadData(loadService Service) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		date := chi.URLParam(r, "date")
		err := loadService.LoadData(date)

		if err != nil {

			json.NewEncoder(rw).Encode(err.Error())
		} else {

			json.NewEncoder(rw).Encode("Data Loaded Successfully.")
		}
	}
}
