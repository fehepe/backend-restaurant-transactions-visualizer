package models

type Response struct {
	buyers       *BuyerList
	products     *ProductList
	transactions *TransactionList
}
