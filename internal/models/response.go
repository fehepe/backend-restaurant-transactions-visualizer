package models

type BuyerDetails struct {
	Buyer               Buyer                `json:"buyer,omitempty"`
	TransactionsDetails []TransactionDetails `json:"transactions,omitempty"`
	BuyersEqIp          BuyerList            `json:"buyerEqIp,omitempty"`
	Products            []Product            `json:"products,omitempty"`
}

type TransactionDetails struct {
	Id       string    `json:"id,omitempty"`
	Ip       string    `json:"ip,omitempty"`
	Device   string    `json:"device,omitempty"`
	Products []Product `json:"products,omitempty"`
}
