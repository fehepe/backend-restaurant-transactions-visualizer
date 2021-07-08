package converter

import (
	"backend-restaurant-transactions-visualizer/internal/models"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"strconv"
	"strings"
)

// Convert bytes to buffer helper
func AnyToBytesBuffer(i interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(i)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

func BuyersRespToObjList(body []byte) (models.BuyerList, error) {

	var buyersList models.BuyerList

	err := json.Unmarshal(body, &buyersList)

	if err != nil {
		return nil, err
	}

	return buyersList, nil
}

func ProductsRespToObjList(body io.Reader) (models.ProductList, error) {

	var productsList models.ProductList

	reader := csv.NewReader(body)
	reader.Comma = '\''

	for {
		data, err := reader.Read()
		if err == io.EOF {
			break
		}
		id := string(data[0])
		name := string(data[1])
		price, _ := strconv.ParseFloat(data[2], 32)
		product := models.Product{Id: id, Name: name, Price: float32(price), DType: []string{"Product"}}
		productsList = append(productsList, product)
	}

	return productsList, nil
}

func TransactionsRespToObjList(body []byte) (models.TransactionList, error) {
	var transactionList models.TransactionList

	data := string(body)

	dataSplit := strings.Split(data, "\x00\x00")

	for _, tranx := range dataSplit {

		params := strings.Split(tranx, "\x00")

		if len(params) < 4 {
			continue
		}
		productIds := strings.Split(params[4][1:len(params[4])-1], ",")

		transaction := models.Transaction{ID: strings.ReplaceAll(params[0], "#", ""), BuyerID: params[1], IP: params[2], Device: params[3], ProductIDs: productIds, DType: []string{"Transaction"}}

		transactionList = append(transactionList, transaction)
	}

	return transactionList, nil

}
