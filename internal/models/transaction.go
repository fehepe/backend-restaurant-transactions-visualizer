package models

type Transaction struct {
	Id         string   `json:"id,,omitempty"`
	BuyerId    string   `json:"buyerID,omitempty"`
	IP         string   `json:"ip,,omitempty"`
	Device     string   `json:"device,omitempty"`
	ProductIds []string `json:"productIDs,omitempty"`
	DType      []string `json:"dgraph.type,omitempty"`
}

type TransactionList []Transaction

type TransactionsListResponse struct {
	Transactions TransactionList `json:"transactions,omitempty"`
}
