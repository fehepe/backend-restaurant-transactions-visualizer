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

	dbConn := os.Getenv("CONN_DB")
	db, err := dgraph.ConnectDB(dbConn)

	//db.LoadSchema()

	if err != nil {
		log.Fatalf("Error creating a new DGraph Client: %v", err)
	}

	client := http.Client{}

	buyerRepository := buyers.NewBuyersRepository(db)
	buyerService := buyers.NewBuyersService(buyerRepository)
	loadRepository := loaddata.NewLoadDataRepository(db)
	loadService := loaddata.NewLoadDataService(loadRepository, client)

	buyerDetails, err := buyerService.FindBuyerById("89722b5")
	fmt.Println(buyerDetails)
	//err = loadService.LoadData("")
	fmt.Println(err, loadService)
}
