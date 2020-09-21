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

	exec.Command("g++", "./data/parser.cpp")
	cmd := exec.Command("./data/a.out")
	stdout, err := cmd.Output()
	if err != nil {
		log.Print("error")
	}
	log.Printf("%s", stdout)
}

func Setup() {
	conn, err := grpc.Dial("10.0.0.6:9080", grpc.WithInsecure())
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

	/*
		op := &api.Operation{
			Schema:          schema,
			RunInBackground: true,
		}

		err = client.Alter(ctx, op)
		if err != nil {
			log.Fatal(err)
		}
	*/
	UploadData()
}

func main() {
	Setup()
}
