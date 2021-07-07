package dgraph

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type Dgraph struct {
	dbClient *dgo.Dgraph
	Cancel   CancelFunc
}

type CancelFunc func()

func ConnectDB(dgraphUrl string) (*Dgraph, error) {

	conn, err := grpc.Dial(dgraphUrl, grpc.WithInsecure())

	if err != nil {
		log.Println("Error Connecting to Dgraph")
		return nil, err
	}

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	cancelFunc := func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection: %v", err)
		}
	}

	return &Dgraph{dbClient: dgraphClient, Cancel: cancelFunc}, nil
}

func (d Dgraph) LoadSchema() error {
	op := &api.Operation{
		Schema: `
		id:          string   @index(exact) .
		name:        string                 .
		age:         int                    .
		price:       int                    .
		buyerID:     string   @index(exact) .
		ip:          string   @index(exact) .
		device:      string                 .
		productIDs:  [string] @index(exact) .
		date:        string   @index(exact) .
		transaction: [uid]    @reverse      .
		product:     [uid]    @reverse      .
		type Buyer {
			id:   string
			name: string 
			age:  int
			
		}
		
		type Product {
			id:    string
			name:  string
			price: int
			
		}
		
		type Transaction {
			id:         string
			buyerID:    string
			ip:         string
			device:     string
			productIDs: [string]
			
		}`,
	}

	ctx := context.Background()

	if err := d.dbClient.Alter(ctx, op); err != nil {
		log.Fatalf("Error while mutating schema: %v\n", err)
		return err
	}

	return nil
}

func (d Dgraph) Query(query string,
	variables map[string]string) (*api.Response, error) {

	ctx := context.Background()
	res, err := d.dbClient.NewTxn().QueryWithVars(ctx, query, variables)

	if err != nil {
		return &api.Response{}, nil
	}
	return res, nil
}

func (d Dgraph) Save(element []byte) error {
	ctx := context.Background()

	mutation := &api.Mutation{
		CommitNow: true,
		SetJson:   element,
	}

	_, err := d.dbClient.NewTxn().Mutate(ctx, mutation)
	if err != nil {
		return err
	}

	return nil
}
