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

func (ls *loadService) LoadData(date string) error {

	dsAPI := datasource.NewDataSourceAPI(ls.clientHttp)

	if date == "" {
		date = fmt.Sprint(time.Now().Local().Unix())
	}

	buyersUIds, err := ls.LoadDataBuyers(*dsAPI, date)

	fmt.Println(buyersUIds)

	if err != nil {
		log.Fatalf("Error LoadDataBuyers: %v", err)
		return err
	}

	productsUIds, err := ls.LoadDataProducts(*dsAPI, date)

	fmt.Println(productsUIds)

	if err != nil {
		log.Fatalf("Error LoadDataProducts: %v", err)
		return err
	}

	err = ls.LoadDataTransactions(*dsAPI, date, buyersUIds, productsUIds)
	if err != nil {
		log.Fatalf("Error LoadDataTransactions: %v", err)
		return err
	}

	return nil
}

func (ls loadService) LoadDataBuyers(dsAPI datasource.DataSource, date string) (map[string]string, error) {
	resp, err := dsAPI.Get("buyers", date)

	if err != nil {
		return nil, err
	}

	buyersToInsert, err := ls.loadRepo.FilterBuyersAlreadyExist(*resp.Buyers)
	if err != nil {
		return nil, err
	}

	json, err := json.Marshal(buyersToInsert)
	if err != nil {
		return nil, err
	}
	err = ls.loadRepo.Insert(json)

	if err != nil {
		return nil, err
	}

	kvpBuyers, err := ls.loadRepo.GetKvpBuyers()
	if err != nil {
		return nil, err
	}
	return kvpBuyers, nil
}

func (ls loadService) LoadDataProducts(dsAPI datasource.DataSource, date string) (map[string]string, error) {

	resp, err := dsAPI.Get("products", date)

	if err != nil {
		return nil, err
	}

	productsToInsert, err := ls.loadRepo.FilterProductsAlreadyExist(*resp.Products)
	if err != nil {
		return nil, err
	}
	json, err := json.Marshal(productsToInsert)
	if err != nil {
		return nil, err
	}
	err = ls.loadRepo.Insert(json)

	if err != nil {
		return nil, err
	}

	kvpProds, err := ls.loadRepo.GetKvpProducts()
	if err != nil {
		return nil, err
	}
	return kvpProds, nil
}

func (ls loadService) LoadDataTransactions(dsAPI datasource.DataSource, date string, kvpBuyers, kvpProds map[string]string) error {
	resp, err := dsAPI.Get("transactions", date)
	fmt.Println(*resp.Transactions)
	if err != nil {
		return err
	}
	transactionsToInsert, err := ls.loadRepo.FilterTransactionsAlreadyExist(*resp.Transactions)
	if err != nil {
		return err
	}

	for idxT, _ := range transactionsToInsert {
		buyerDbUid := kvpBuyers[transactionsToInsert[idxT].Buyer.UId]
		transactionsToInsert[idxT].Buyer.UId = buyerDbUid
		for idx, _ := range transactionsToInsert[idxT].Products {
			prodDbUid := kvpProds[transactionsToInsert[idxT].Products[idx].UId]

			transactionsToInsert[idxT].Products[idx].UId = prodDbUid
		}
	}
	json, err := json.Marshal(transactionsToInsert)
	if err != nil {
		return err
	}
	err = ls.loadRepo.Insert(json)

	if err != nil {
		return err
	}

	return nil
}
