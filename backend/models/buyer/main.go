package buyer

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

func FindAll() ([]byte, error) {
	ctx := context.TODO()
	conn, err := grpc.Dial("10.0.0.6:9080", grpc.WithInsecure())
	if err != nil {
		log.Println("hi")
		log.Fatal(err)
	}
	defer conn.Close()
	Client := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	txn := Client.NewTxn()
	defer txn.Discard(ctx)

	query := `query {
		  Buyers(func: eq(DType,"Buyer")){
			uid 
			name
			age
		  }
	}`

	res, err := txn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	conn.Close()
	return res.Json, nil
}

func GetById(id string) ([]byte, error) {
	ctx := context.TODO()
	conn, err := grpc.Dial("10.0.0.6:9080", grpc.WithInsecure())
	if err != nil {
		log.Println("hi")
		log.Fatal(err)
	}
	defer conn.Close()
	Client := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	txn := Client.NewTxn()
	defer txn.Discard(ctx)

	query := `
		query var($a: string){
		  var(func: uid(0x58fe)){
			~buyer_id {
			  ips as ip
			}
		  }
		  
		  buyer(func: uid(0x58fe)){
			uid
			name
			Purchases: ~buyer_id {
			  uid
			  ip
			  device
			  Products {
				uid
				name
				price
				}
			}
			}
		  
			others(func: eq(ip,val(ips))){
				relation as uid
			ip
			Buyer: buyer_id {
			  uid
			  name
			}
		  }
		  
			var(func:uid(relation)){
				Products {
					recommended_products as uid
				}
			}
		  
		  recommendations(func: uid(recommended_products)){
			uid
			name
			price
		  }
		}`

	res, err := txn.QueryWithVars(ctx, query, map[string]string{"$a": id})
	if err != nil {
		return nil, err
	}
	conn.Close()
	return res.Json, nil
}
