package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("www.miguelestevez.xyz:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	type Person struct {
		Uid   string   `json:"uid,omitempty"`
		Name  string   `json:"name,omitempty"`
		DType []string `json:"dgraph.type,omitempty"`
	}

	type ans struct {
		Persons []Person
	}

	txn := dgraphClient.NewTxn()
	ctx := context.Background()
	defer txn.Discard(ctx)

	query := ` query { persons(func: has(name)) {
				uid
				name
				DType
		}}
		`
	res, err := txn.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	var resp ans
	json.Unmarshal(res.GetJson(), &resp)
	fmt.Print(resp)
}
