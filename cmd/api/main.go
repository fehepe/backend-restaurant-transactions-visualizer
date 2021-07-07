package main

import (
	"backend-restaurant-transactions-visualizer/internal/buyers"
	"backend-restaurant-transactions-visualizer/pkg/db/dgraph"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err.Error())
	}
	dbConn := os.Getenv("CONN_DB")
	db, err := dgraph.ConnectDB(dbConn)

	if err != nil {
		log.Fatalf("Error creating a new DGraph Client: %v", err)
	}

	buyerRepository := buyers.NewBuyersRepository(db)
	buyerService := buyers.NewBuyersService(buyerRepository)

	buyers, _ := buyerService.FindAllBuyers()
	fmt.Println(buyers)

}
