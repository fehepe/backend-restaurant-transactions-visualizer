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
	Buyers       *models.BuyerList
	Products     *models.ProductList
	Transactions *models.TransactionList
}

type DataSource struct {
	apiClient http.Client
}

func NewDataSourceAPI(client http.Client) *DataSource {
	return &DataSource{apiClient: client}

}

func (ds DataSource) Get(route string, date string) (Responses, error) {

	buyerlist := models.BuyerList{}
	productslist := models.ProductList{}
	transactionslist := models.TransactionList{}
	baseUrl := os.Getenv("BASE_URL")

	parsedUrl, err := url.Parse(baseUrl + route)

	if err != nil {
		return Responses{}, err
	}

	queryParams := parsedUrl.Query()

	queryParams.Add("date", date)

	parsedUrl.RawQuery = queryParams.Encode()

	resp, err := ds.apiClient.Get(parsedUrl.String())

	if err != nil {
		return Responses{}, err
	}
	defer resp.Body.Close()

	switch route {
	case "buyers":
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Responses{}, err
		}
		buyerlist, err = converter.BuyersRespToObjList(body)
	case "products":
		productslist, err = converter.ProductsRespToObjList(resp.Body)
	case "transactions":
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Responses{}, err
		}
		transactionslist, err = converter.TransactionsRespToObjList(body)
	default:
		log.Fatal("The selected route was wrong.")
		return Responses{}, err
	}

	if err != nil {
		return Responses{}, err
	}

	return Responses{Buyers: &buyerlist, Products: &productslist, Transactions: &transactionslist}, nil
}
