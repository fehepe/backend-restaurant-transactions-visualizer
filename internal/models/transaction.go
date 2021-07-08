package models

type Transaction struct {
	ID         string   `json:"id,,omitempty"`
	BuyerID    string   `json:"buyerID,omitempty"`
	IP         string   `json:"ip,,omitempty"`
	Device     string   `json:"device,omitempty"`
	ProductIDs []string `json:"productIDs,omitempty"`
	DType      []string `json:"dgraph.type,omitempty"`
}

type TransactionList []Transaction

type TransactionsListResponse struct {
	Transactions TransactionList `json:"transactions,omitempty"`
}
