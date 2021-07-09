package models

type Product struct {
	UId   string   `json:"uid,omitempty"`
	Id    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Price float32  `json:"price,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

type ProductList []Product

type ProductsListResponse struct {
	Products ProductList `json:"products,omitempty"`
}
