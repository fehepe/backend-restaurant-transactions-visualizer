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
	resp, err := dsAPI.Get("buyers", date)

	if err != nil {
		return err
	}

	buyersToInsert, err := ls.loadRepo.FilterBuyersAlreadyExist(*resp.Buyers)
	if err != nil {
		return err
	}
	ls.loadRepo.InsertBuyers(buyersToInsert)

	resp, err = dsAPI.Get("products", date)

	if err != nil {
		return err
	}

	productsToInsert, err := ls.loadRepo.FilterProductsAlreadyExist(*resp.Products)
	if err != nil {
		return err
	}
	ls.loadRepo.InsertProduct(productsToInsert)

	resp, err = dsAPI.Get("transactions", date)
	fmt.Println(*resp.Transactions)
	if err != nil {
		return err
	}
	transactionsToInsert, err := ls.loadRepo.FilterTransactionsAlreadyExist(*resp.Transactions)
	if err != nil {
		return err
	}
	ls.loadRepo.InsertTransactions(transactionsToInsert)

	return nil
}
