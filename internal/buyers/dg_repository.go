package buyers

import (
	"backend-restaurant-transactions-visualizer/internal/models"
	"backend-restaurant-transactions-visualizer/pkg/db/dgraph"
	"backend-restaurant-transactions-visualizer/pkg/queries"
	"encoding/json"
)

type BuyerListResponse struct {
	Buyers models.BuyerList `json:"buyers,omitempty"`
}

type Repository interface {
	FindAllBuyers() ([]models.Buyer, error)
}

type dgraphRepository struct {
	db *dgraph.Dgraph
}

func NewBuyersRepository(db *dgraph.Dgraph) *dgraphRepository {
	return &dgraphRepository{db: db}
}

func (d *dgraphRepository) FindAllBuyers() ([]models.Buyer, error) {

	resp, err := d.db.Query(queries.AllBuyers, nil)

	if err != nil {
		return nil, err
	}

	var dgraphResponse BuyerListResponse

	if err := json.Unmarshal(resp.GetJson(), &dgraphResponse); err != nil {
		return nil, err
	}

	return dgraphResponse.Buyers, nil
}
