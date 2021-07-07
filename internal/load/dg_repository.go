package load

import "backend-restaurant-transactions-visualizer/internal/models"

type Repository interface {
}
type BuyersListResponse struct {
	Buyers []models.Buyer `json:"buyers,omitempty"`
}

type ProductsListResponse struct {
	Products []models.Product `json:"products,omitempty"`
}

type TransactionsListResponse struct {
	Transactions []models.Transaction `json:"transactions,omitempty"`
}
