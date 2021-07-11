package models

type RecommendationsResponse struct {
	Product         []Product `json:"product,omitempty"`
	Recommendations []Product `json:"productsRecomendation,omitempty"`
}

type TransactionDetails struct {
	Id              string                    `json:"id,omitempty"`
	Ip              string                    `json:"ip,omitempty"`
	Device          string                    `json:"device,omitempty"`
	Products        ProductList               `json:"products,omitempty"`
	Recommendations []RecommendationsResponse `json:"recommendations,omitempty"`
}

type BuyerEqIp struct {
	Id     string `json:"id,omitempty"`
	Ip     string `json:"ip,omitempty"`
	Device string `json:"device,omitempty"`
	Buyer  Buyer  `json:"buyer,omitempty"`
}
