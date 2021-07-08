package models

type Buyer struct {
	Id    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Age   int      `json:"age,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

type BuyerList []Buyer

type BuyersListResponse struct {
	Buyers BuyerList `json:"buyers,omitempty"`
}
