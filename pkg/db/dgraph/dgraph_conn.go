package dgraph

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/fehepe/backend-restaurant-transactions-visualizer/pkg/queries"
	"google.golang.org/grpc"
)

type Dgraph struct {
	dbClient *dgo.Dgraph
	Cancel   CancelFunc
	ctx      context.Context
}

type CancelFunc func()

func ConnectDB(dgraphUrl string) (*Dgraph, error) {
	context := context.Background()
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

	return &Dgraph{dbClient: dgraphClient, Cancel: cancelFunc, ctx: context}, nil
}

func (d Dgraph) LoadSchema() error {
	op := &api.Operation{
		Schema: queries.Schema,
	}

	if err := d.dbClient.Alter(d.ctx, op); err != nil {
		log.Fatalf("Error while mutating schema: %v\n", err)
		return err
	}

	return nil
}

func (d Dgraph) Query(query string,
	variables map[string]string) (*api.Response, error) {

	res, err := d.dbClient.NewTxn().QueryWithVars(d.ctx, query, variables)

	if err != nil {
		return &api.Response{}, nil
	}
	return res, nil
}

func (d Dgraph) Save(element []byte) error {

	mutation := &api.Mutation{
		CommitNow: true,
	}
	mutation.SetJson = element
	_, err := d.dbClient.NewTxn().Mutate(d.ctx, mutation)
	if err != nil {
		return err
	}

	return nil
}

func (d Dgraph) Insert(element []byte) error {

	mutation := &api.Mutation{
		CommitNow: true,
		SetJson:   element,
	}

	_, err := d.dbClient.NewTxn().Mutate(d.ctx, mutation)
	if err != nil {
		return err
	}

	return nil
}
