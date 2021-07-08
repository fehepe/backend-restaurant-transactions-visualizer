package main

import (
	"backend-restaurant-transactions-visualizer/internal/buyers"
	"backend-restaurant-transactions-visualizer/internal/loaddata"
	"backend-restaurant-transactions-visualizer/pkg/db/dgraph"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal(err.Error())
	}
	schemaLoaded := os.Getenv("SCHEMA_LOAD")

	if schemaLoaded == "0" {
		os.Setenv("SCHEMA_LOAD", "1")
	}

	dbConn := os.Getenv("CONN_DB")
	db, err := dgraph.ConnectDB(dbConn)

	if err != nil {
		log.Fatalf("Error creating a new DGraph Client: %v", err)
	}

	client := http.Client{}
	//context := context.Background()

	buyerRepository := buyers.NewBuyersRepository(db)
	buyerService := buyers.NewBuyersService(buyerRepository)
	loadRepository := loaddata.NewLoadDataRepository(db)
	loadService := loaddata.NewLoadDataService(loadRepository, client)

	loadService.LoadData("")

	fmt.Println(buyerService)

}
