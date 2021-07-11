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
	FindBuyerById(buyerId string) (*models.BuyerDetails, error)
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

func (d *dgraphRepository) FindBuyerById(buyerId string) (*models.BuyerDetails, error) {

	resp, err := d.db.Query(queries.FindBuyerById, nil)

	if err != nil {
		log.Fatal("Error running the query of Find all Buyers.")
		return nil, err
	}

	var dgraphResponse models.BuyerDetails

	if err := json.Unmarshal(resp.GetJson(), &dgraphResponse); err != nil {
		return nil, err
	}

	return &dgraphResponse, nil
}
