package converter

import (
	"backend-restaurant-transactions-visualizer/internal/models"
	"bytes"
	"encoding/json"
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

	// var validBuyersToLoad []models.Buyer
	// duplicate := make(map[string]bool)

	// for _, item := range buyersList {
	// 	buyer, err := models.Buyer{Id: item.Id, Name: item.Name, Age: item.Age}

	// 	if err != nil {
	// 		log.Println(err)
	// 		continue
	// 	}

	// 	exist := duplicate[item.Id]

	// 	if exist {
	// 		continue
	// 	} else {
	// 		duplicate[item.Id] = true
	// 	}

	// 	validBuyersToLoad = append(validBuyersToLoad, *buyer)
	// }

	return buyersList, nil
}

func ProductsRespToObjList(body []byte) (models.ProductList, error) {
	var productsList models.ProductList

	// err := json.Unmarshal(body, &productsList)

	// if err != nil {
	// 	return nil, err
	// }
	return productsList, nil
}
func TransactionsRespToObjList(body []byte) (models.TransactionList, error) {
	var transactionList models.TransactionList

	// err := json.Unmarshal(body, &transactionList)

	// if err != nil {
	// 	return nil, err
	// }
	return transactionList, nil

}
