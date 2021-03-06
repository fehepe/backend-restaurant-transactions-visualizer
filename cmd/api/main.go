package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fehepe/backend-restaurant-transactions-visualizer/internal/buyers"
	"github.com/fehepe/backend-restaurant-transactions-visualizer/internal/loaddata"
	"github.com/fehepe/backend-restaurant-transactions-visualizer/internal/server"
	"github.com/fehepe/backend-restaurant-transactions-visualizer/pkg/db/dgraph"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal(err.Error())
	}

	dbConn := os.Getenv("CONN_DB")

	db, err := dgraph.ConnectDB(dbConn)

	db.LoadSchema()

	if err != nil {
		log.Fatalf("Error creating a new DGraph Client: %v", err)
	}

	client := http.Client{}

	buyerRepository := buyers.NewBuyersRepository(db)
	buyerService := buyers.NewBuyersService(buyerRepository)
	loadRepository := loaddata.NewLoadDataRepository(db)
	loadService := loaddata.NewLoadDataService(loadRepository, client)

	err = server.Run(buyerService, loadService)

	if err != nil {
		log.Fatalf("Error running the server: %v", err)
	}

}
