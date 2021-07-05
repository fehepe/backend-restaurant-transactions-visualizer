package dgraph

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

// CancelFunc represents a function that should be deffered on its called after
// use
type CancelFunc func()

// Connect returns a new connection to a local graphdb database
func Connect() (*dgo.Dgraph, CancelFunc) {
	conn, err := grpc.Dial(os.Getenv("DB_URL"), grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	return dgraphClient, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection: %v", err)
		}
	}
}

// Save saves a marshed valid JSON into the database
func Save(client *dgo.Dgraph, element []byte) error {
	ctx := context.Background()

	mutation := &api.Mutation{
		CommitNow: true,
		SetJson:   element,
	}

	_, err := client.NewTxn().Mutate(ctx, mutation)
	if err != nil {
		return err
	}

	return nil
}

// BulkConnect takes an array of string and fields to make a bulk upsert
// (query + mutation) on a series of fields to connect related edges
// using the supplied fields with the edge name
func BulkConnect(client *dgo.Dgraph, field1, field2, edge string,
	values []string) error {
	ctx := context.Background()

	var query bytes.Buffer

	query.WriteString("query {")

	var mutations []*api.Mutation

	for idx, val := range values {

		// The field names are placeholder names for the query, they bear no
		// effect on the query itset
		q := fmt.Sprintf(`
			field_1_%d as var(func: eq(%s, "%s"))
			field_2_%d as var(func: eq(%s, "%s"))
		`, idx, field1, val, idx, field2, val)

		query.WriteString(q)

		mu := &api.Mutation{
			SetNquads: []byte(fmt.Sprintf(`uid(field_1_%d) <%s> uid(field_2_%d) .`,
				idx, edge, idx)),
		}

		mutations = append(mutations, mu)
	}

	req := &api.Request{
		Query:     query.String(),
		Mutations: mutations,
		CommitNow: true,
	}

	if _, err := client.NewTxn().Do(ctx, req); err != nil {
		return err
	}

	return nil
}

// Query perfoms a query to the database according to the parameters
// supplied and returns a response from the db
func Query(client *dgo.Dgraph, query string,
	variables map[string]string) (*api.Response, error) {

	ctx := context.Background()
	res, err := client.NewTxn().QueryWithVars(ctx, query, variables)

	if err != nil {
		return &api.Response{}, nil
	}
	return res, nil
}
