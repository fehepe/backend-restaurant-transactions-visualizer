package loaddata

import (
	"encoding/json"
	"fmt"
	"github.com/fehepe/backend-restaurant-transactions-visualizer/internal/models"
	datasource "github.com/fehepe/backend-restaurant-transactions-visualizer/pkg/dataSource"
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

	if err != nil {
		log.Fatalf("Error LoadDataBuyers: %v", err)
		return err
	}

	productsUIds, err := ls.LoadDataProducts(*dsAPI, date)

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
func eliminateDuplicates(resp datasource.Responses) datasource.Responses {
	kvpResponse := make(map[string]string)
	buyerlist := models.BuyerList{}
	productslist := models.ProductList{}
	transactionslist := models.TransactionList{}

	if len(*resp.Buyers) != 0 {

		for _, buyer := range *resp.Buyers {
			if kvpResponse[buyer.Id] == "" {

				kvpResponse[buyer.Id] = buyer.Id
				buyerlist = append(buyerlist, buyer)
			} else {
				fmt.Println(buyer.Id)
			}

		}

	} else if len(*resp.Products) != 0 {
		for _, prod := range *resp.Products {
			if kvpResponse[prod.Id] == "" {

				kvpResponse[prod.Id] = prod.Id
				productslist = append(productslist, prod)
			} else {
				fmt.Println(prod.Id)
			}

		}
	} else {
		for _, trnx := range *resp.Transactions {
			if kvpResponse[trnx.Id] == "" {

				kvpResponse[trnx.Id] = trnx.Id
				transactionslist = append(transactionslist, trnx)
			} else {
				fmt.Println(trnx.Id)
			}

		}
	}
	return datasource.Responses{Buyers: &buyerlist, Products: &productslist, Transactions: &transactionslist}

}
func (ls loadService) LoadDataBuyers(dsAPI datasource.DataSource, date string) (map[string]string, error) {
	resp, err := dsAPI.Get("buyers", date)
	if err != nil {
		return nil, err
	}
	resp = eliminateDuplicates(resp)

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
	resp = eliminateDuplicates(resp)
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

	if err != nil {
		return err
	}
	resp = eliminateDuplicates(resp)
	transactionsToInsert, err := ls.loadRepo.FilterTransactionsAlreadyExist(*resp.Transactions)
	if err != nil {
		return err
	}

	for idxT := range transactionsToInsert {
		buyerDbUid := kvpBuyers[transactionsToInsert[idxT].Buyer.UId]
		transactionsToInsert[idxT].Buyer.UId = buyerDbUid
		for idx := range transactionsToInsert[idxT].Products {
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
