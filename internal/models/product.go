package models

type Product struct {
	Id    string   `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Price float32  `json:"price,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}
