package loaddata

import "backend-restaurant-transactions-visualizer/pkg/db/dgraph"

type Repository interface {
	Insert() error
}

type dgraphRepository struct {
	db *dgraph.Dgraph
}

func NewLoadDataRepository(db *dgraph.Dgraph) *dgraphRepository {
	return &dgraphRepository{db: db}
}

func (dr dgraphRepository) Insert() error {
	return nil
}
