package DB

import (
	"log"

	"google.golang.org/grpc"
)

func GetClient() *grpc.ClientConn {
	conn, err := grpc.Dial("10.0.0.6:9080", grpc.WithInsecure())
	if err != nil {
		log.Println("hi")
		log.Fatal(err)
	}
	defer conn.Close()
	return conn
}

/*
func Setup() {
	schema := `
		type Buyer{
			id
			name
			age
			DType
			created_at
		}

		type Product {
			id
			name
			price
			DType
			created_at
		}

		type Transaction {
			id
			buyer_id
			ip
			device
			products
			DType
			created_at
		}

		device: string .
		id: string @index(exact) .
		name: string @index(exact) .
		age: int .
		DType: string index(exact) .
		date: datetime .
		products: [string] .
		ip: string .
		buyer_id: uid .
		price: float .
	`

	ctx := context.Background()
	op := &api.Operation{
		Schema:          schema,
		RunInBackground: true,
	}

	err := GetClient().Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}
}
*/
