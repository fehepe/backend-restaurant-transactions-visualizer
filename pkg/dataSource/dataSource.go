package datasource

import (
	"backend-restaurant-transactions-visualizer/internal/models"
	"backend-restaurant-transactions-visualizer/pkg/converter"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type dataSource interface {
	Get(route string, date string) (Responses, error)
}

type Responses struct {
	buyers       models.BuyerList
	products     models.ProductList
	transactions models.TransactionList
}

type DataSource struct {
	apiClient http.Client
}

func NewDataSourceAPI(client http.Client) *DataSource {
	return &DataSource{apiClient: client}

}

func (ds DataSource) Get(route string, date string) (*Responses, error) {

	buyerlist := models.BuyerList{}
	productslist := models.ProductList{}
	transactionslist := models.TransactionList{}
	baseUrl := os.Getenv("BASE_URL")

	parsedUrl, err := url.Parse(baseUrl + route)

	if err != nil {
		return nil, err
	}

	queryParams := parsedUrl.Query()

	queryParams.Add("date", date)

	parsedUrl.RawQuery = queryParams.Encode()

	resp, err := ds.apiClient.Get(parsedUrl.String())

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch route {
	case "buyers":
		buyerlist, err = converter.BuyersRespToObjList(body)
	case "products":
		productslist, err = converter.ProductsRespToObjList(body)
	case "transactions":
		transactionslist, err = converter.TransactionsRespToObjList(body)
	default:
		log.Fatal("The selected route was wrong.")
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &Responses{buyers: buyerlist, products: productslist, transactions: transactionslist}, nil
}
