package loaddata

import (
	datasource "backend-restaurant-transactions-visualizer/pkg/dataSource"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Service interface {
	LoadData(date string) error
}

type loadService struct {
	loadRepo   Repository
	clientHttp http.Client
}

func NewLoadDataService(loadRepo Repository, clientHttp http.Client) *loadService {
	return &loadService{loadRepo: loadRepo, clientHttp: clientHttp}
}

func (ls loadService) LoadData(date string) error {

	dsAPI := datasource.NewDataSourceAPI(ls.clientHttp)

	if date == "" {
		date = fmt.Sprint(time.Now().Local().Unix())
	}
	err := ls.LoadDataBuyers(*dsAPI, date)
	if err != nil {
		log.Fatalf("Error LoadDataBuyers: %v", err)
		return err
	}
	err = ls.LoadDataTransactions(*dsAPI, date)
	if err != nil {
		log.Fatalf("Error LoadDataTransactions: %v", err)
		return err
	}
	err = ls.LoadDataProducts(*dsAPI, date)
	if err != nil {
		log.Fatalf("Error LoadDataProducts: %v", err)
		return err
	}

	return nil
}

func (ls loadService) LoadDataBuyers(dsAPI datasource.DataSource, date string) error {
	resp, err := dsAPI.Get("buyers", date)

	if err != nil {
		return err
	}

	buyersToInsert, err := ls.loadRepo.FilterBuyersAlreadyExist(*resp.Buyers)
	if err != nil {
		return err
	}
	json, err := json.Marshal(buyersToInsert)
	if err != nil {
		return err
	}
	ls.loadRepo.Insert(json)
	return nil
}

func (ls loadService) LoadDataProducts(dsAPI datasource.DataSource, date string) error {

	resp, err := dsAPI.Get("products", date)

	if err != nil {
		return err
	}

	productsToInsert, err := ls.loadRepo.FilterProductsAlreadyExist(*resp.Products)
	if err != nil {
		return err
	}
	json, err := json.Marshal(productsToInsert)
	if err != nil {
		return err
	}
	ls.loadRepo.Insert(json)

	return nil
}

func (ls loadService) LoadDataTransactions(dsAPI datasource.DataSource, date string) error {
	resp, err := dsAPI.Get("transactions", date)
	fmt.Println(*resp.Transactions)
	if err != nil {
		return err
	}
	transactionsToInsert, err := ls.loadRepo.FilterTransactionsAlreadyExist(*resp.Transactions)
	if err != nil {
		return err
	}
	json, err := json.Marshal(transactionsToInsert)
	if err != nil {
		return err
	}
	ls.loadRepo.Insert(json)

	return nil
}
