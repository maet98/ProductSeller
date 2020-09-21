package buyer

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/maet98/sellerApp/config"
	"google.golang.org/grpc"
)

func FindAll() ([]byte, error) {
	ctx := context.TODO()
	conn, err := grpc.Dial(config.Config.Url(), grpc.WithInsecure())
	if err != nil {
		log.Println("hi")
		log.Fatal(err)
	}
	defer conn.Close()
	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	txn := client.NewTxn()
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
	conn, err := grpc.Dial(config.Config.Url(), grpc.WithInsecure())
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
			var(func: uid($a)){
			~buyer_id {
				ips as ip
				Products {
				  myProducts as uid
				}
			}
		  }
		  
			buyer(func: uid($a)){
				uid
				name
				Purchases: ~buyer_id {
				  uid
				  ip
				  Date
				  device
				  Products {
						uid
						name
						price
					}
				}
			}
		  
			var(func: eq(ip,val(ips))){
				relation as uid
					others as Buyer: buyer_id @filter(not(uid($a)))
			  }
      
		  Others(func: uid(others)){
			uid
			name
		  }
		  
			var(func:uid(relation)){
				Products @filter(not(uid(myProducts))) {
					recommended_products as uid
				}
			}
		  
		  recommendations(func: uid(recommended_products)){
			uid
			name
			price
			Date
		  }
		}`

	res, err := txn.QueryWithVars(ctx, query, map[string]string{"$a": id})
	if err != nil {
		return nil, err
	}
	conn.Close()
	return res.Json, nil
}
