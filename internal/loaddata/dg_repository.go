package loaddata

import (
	"encoding/json"
	"github.com/fehepe/backend-restaurant-transactions-visualizer/internal/models"
	"github.com/fehepe/backend-restaurant-transactions-visualizer/pkg/db/dgraph"
	"github.com/fehepe/backend-restaurant-transactions-visualizer/pkg/queries"
	"log"
)

type Repository interface {
	Insert(entity []byte) error

	GetKvpBuyers() (map[string]string, error)
	GetKvpProducts() (map[string]string, error)
	FilterBuyersAlreadyExist(buyers models.BuyerList) (models.BuyerList, error)
	FilterProductsAlreadyExist(buyers models.ProductList) (models.ProductList, error)
	FilterTransactionsAlreadyExist(buyers models.TransactionList) (models.TransactionList, error)
}

type dgraphRepository struct {
	db *dgraph.Dgraph
}

func NewLoadDataRepository(db *dgraph.Dgraph) *dgraphRepository {
	return &dgraphRepository{db: db}
}

func (dr dgraphRepository) Insert(entity []byte) error {

	err := dr.db.Save(entity)
	if err != nil {
		log.Fatalf("Error Inserting the entity: %v", err)
		return err
	}
	return nil
}

func (dr dgraphRepository) FilterBuyersAlreadyExist(buyers models.BuyerList) (models.BuyerList, error) {
	resultList := models.BuyerList{}

	resp, err := dr.db.Query(queries.FindBuyers, nil)

	if err != nil {
		log.Fatal("Error running the query of AlreadyExist by ID.")
		return nil, err
	}

	var dgraphResponse models.BuyersListResponse

	if err := json.Unmarshal(resp.GetJson(), &dgraphResponse); err != nil {
		return nil, err
	}

	for _, buyer := range buyers {
		exist := containsBuyer(dgraphResponse.Buyers, buyer)
		if exist {
			continue
		} else {
			resultList = append(resultList, buyer)
		}

	}
	return resultList, nil
}

func (dr dgraphRepository) FilterProductsAlreadyExist(prods models.ProductList) (models.ProductList, error) {
	resultList := models.ProductList{}
	resp, err := dr.db.Query(queries.FindProducts, nil)

	if err != nil {
		log.Fatal("Error running the query of AlreadyExist.")
		return nil, err
	}

	var dgraphResponse models.ProductsListResponse

	if err := json.Unmarshal(resp.GetJson(), &dgraphResponse); err != nil {
		return nil, err
	}

	for _, prod := range prods {
		exist := containsProduct(dgraphResponse.Products, prod)
		if exist {
			continue
		} else {
			resultList = append(resultList, prod)
		}

	}
	return resultList, nil
}
func (dr dgraphRepository) FilterTransactionsAlreadyExist(tranxs models.TransactionList) (models.TransactionList, error) {
	resultList := models.TransactionList{}
	resp, err := dr.db.Query(queries.FindTransactions, nil)

	if err != nil {
		log.Fatal("Error running the query of AlreadyExist.")
		return nil, err
	}

	var dgraphResponse models.TransactionsListResponse

	if err := json.Unmarshal(resp.GetJson(), &dgraphResponse); err != nil {
		return nil, err
	}

	for _, tranx := range tranxs {
		exist := containsTransactions(dgraphResponse.Transactions, tranx)
		if exist {
			continue
		} else {
			resultList = append(resultList, tranx)
		}

	}
	return resultList, nil
}
func (dr dgraphRepository) GetKvpBuyers() (map[string]string, error) {
	resp, err := dr.db.Query(queries.FindBuyers, nil)

	if err != nil {
		log.Fatal("Error running the query of Find all Buyers.")
		return nil, err
	}

	var dgraphResponse models.BuyersListResponse

	if err := json.Unmarshal(resp.GetJson(), &dgraphResponse); err != nil {
		return nil, err
	}

	kvpBuyers := make(map[string]string)
	for _, buyers := range dgraphResponse.Buyers {
		kvpBuyers[buyers.Id] = buyers.UId
	}

	return kvpBuyers, nil
}

func (dr dgraphRepository) GetKvpProducts() (map[string]string, error) {
	resp, err := dr.db.Query(queries.FindProducts, nil)

	if err != nil {
		log.Fatal("Error running the query of Find all Buyers.")
		return nil, err
	}

	var dgraphResponse models.ProductsListResponse

	if err := json.Unmarshal(resp.GetJson(), &dgraphResponse); err != nil {
		return nil, err
	}

	kvpProducts := make(map[string]string)
	for _, prod := range dgraphResponse.Products {
		kvpProducts[prod.Id] = prod.UId
	}

	return kvpProducts, nil
}

func containsBuyer(list models.BuyerList, buyer models.Buyer) bool {
	for _, v := range list {
		if v.Id == buyer.Id {
			return true
		}
	}

	return false
}
func containsProduct(list models.ProductList, prod models.Product) bool {
	for _, v := range list {
		if v.Id == prod.Id {
			return true
		}
	}

	return false
}
func containsTransactions(list models.TransactionList, tranx models.Transaction) bool {
	for _, v := range list {
		if v.Id == tranx.Id {
			return true
		}
	}

	return false
}
