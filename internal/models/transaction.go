package models

type Transaction struct {
	UId        string    `json:"uid,omitempty"`
	Id         string    `json:"id,omitempty"`
	BuyerId    string    `json:"buyerID,omitempty"`
	Buyer      Buyer     `json:"buyer,omitempty"`
	IP         string    `json:"ip,omitempty"`
	Device     string    `json:"device,omitempty"`
	ProductIds []string  `json:"productIDs,omitempty"`
	Products   []Product `json:"products,omitempty"`
	DType      []string  `json:"dgraph.type,omitempty"`
}

type TransactionList []Transaction

type TransactionsListResponse struct {
	Transactions TransactionList `json:"transactions,omitempty"`
}
