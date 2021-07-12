package loaddata

import (
	"encoding/json"
	"net/http"
)

func LoadData(loadService Service) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		date := r.URL.Query().Get("date")

		err := loadService.LoadData(date)

		if err != nil {

			json.NewEncoder(rw).Encode(err.Error())
		} else {

			json.NewEncoder(rw).Encode("Data Loaded Successfully.")
		}
	}
}
