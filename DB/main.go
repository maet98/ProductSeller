package main

import (
	"context"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

func UploadData() {
	exec.Command("g++", "parser.cpp").Run()
	exec.Command("./a.out").Run()
	log.Print("DONE")
}

func Setup() {
	conn, err := grpc.Dial("wolfisc.ddns.net:9080", grpc.WithInsecure())
	if err != nil {
		log.Println("hi")
		log.Fatal(err)
	}
	defer conn.Close()

	ctx := context.Background()
	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	txn := client.NewTxn()
	defer txn.Discard(ctx)

	buffer, err := ioutil.ReadFile("./schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	schema := string(buffer)
	log.Println(schema)

	op := &api.Operation{
		Schema:          schema,
		RunInBackground: true,
	}

	err = client.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}
	UploadData()
	buffer, err = ioutil.ReadFile("./query.json")

	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		SetJson: buffer,
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Setup()
}
