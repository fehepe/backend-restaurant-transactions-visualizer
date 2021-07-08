package loaddata

import (
	datasource "backend-restaurant-transactions-visualizer/pkg/dataSource"
	"fmt"
	"net/http"
	"time"
)

type Service interface {
	LoadData(date string) error
}

type buyerService struct {
	buyerRepo  Repository
	clientHttp http.Client
}

func NewLoadDataService(buyerRepo Repository, clientHttp http.Client) *buyerService {
	return &buyerService{buyerRepo: buyerRepo, clientHttp: clientHttp}
}

func (bs buyerService) LoadData(date string) error {

	dsAPI := datasource.NewDataSourceAPI(bs.clientHttp)

	if date == "" {
		date = fmt.Sprint(time.Now().Local().Unix())
	}
	buyers, err := dsAPI.Get("buyers", date)

	if err != nil {
		return err
	}

	fmt.Println(buyers)

	return nil
}