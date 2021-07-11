package buyers

import (
	"backend-restaurant-transactions-visualizer/internal/models"
	"backend-restaurant-transactions-visualizer/pkg/db/dgraph"
	"backend-restaurant-transactions-visualizer/pkg/queries"
	"encoding/json"
	"log"
)

type Repository interface {
	FindAllBuyers() (models.BuyerList, error)
	FindBuyerById(buyerId string) (*BuyerDetailsResponse, error)
}
type BuyerDetailsResponse struct {
	Buyer               []models.Buyer              `json:"buyer,omitempty"`
	TransactionsDetails []models.TransactionDetails `json:"transactions,omitempty"`
	BuyersEqIp          []models.BuyerEqIp          `json:"buyerEqIp,omitempty"`
}

type dgraphRepository struct {
	db *dgraph.Dgraph
}

func NewBuyersRepository(db *dgraph.Dgraph) *dgraphRepository {
	return &dgraphRepository{db: db}
}

func (d *dgraphRepository) FindAllBuyers() (models.BuyerList, error) {

	resp, err := d.db.Query(queries.FindBuyers, nil)

	if err != nil {
		log.Fatal("Error running the query of Find all Buyers.")
		return nil, err
	}

	var dgraphResponse models.BuyersListResponse

	if err := json.Unmarshal(resp.GetJson(), &dgraphResponse); err != nil {
		return nil, err
	}

	return dgraphResponse.Buyers, nil
}

func (d *dgraphRepository) FindBuyerById(buyerId string) (*BuyerDetailsResponse, error) {

	vars := map[string]string{"$id": buyerId}
	resp, err := d.db.Query(queries.FindBuyerDetailsById, vars)

	if err != nil {
		log.Fatal("Error running the query of Find Buyer Details By Id.")
		return nil, err
	}

	var dgraphResponse BuyerDetailsResponse

	if err := json.Unmarshal(resp.Json, &dgraphResponse); err != nil {
		return nil, err
	}
	for idx, tranx := range dgraphResponse.TransactionsDetails {
		dgraphResponse.TransactionsDetails[idx], err = d.getRecommendations(tranx)
		if err != nil {
			return nil, err
		}
	}

	return &dgraphResponse, nil
}

func (d *dgraphRepository) getRecommendations(traxn models.TransactionDetails) (models.TransactionDetails, error) {

	for _, prod := range traxn.Products {
		vars := map[string]string{"$id": prod.Id}
		resp, err := d.db.Query(queries.FindRecomendationsByProdId, vars)

		if err != nil {
			log.Fatal("Error running the query of get Recommendations By Id.")
			return models.TransactionDetails{}, err
		}

		var dgraphResponse models.RecommendationsResponse

		if err := json.Unmarshal(resp.Json, &dgraphResponse); err != nil {
			return models.TransactionDetails{}, err
		}

		traxn.Recommendations = append(traxn.Recommendations, dgraphResponse)
	}

	return traxn, nil

}
