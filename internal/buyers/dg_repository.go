package buyers

import (
	"backend-restaurant-transactions-visualizer/internal/models"
	"backend-restaurant-transactions-visualizer/pkg/db/dgraph"
)

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

	buyers := []models.Buyer{
		{
			Id:    "1",
			Age:   22,
			Name:  "Prueba",
			DType: []string{"Buyer"},
		},
	}

	return buyers, nil
}
