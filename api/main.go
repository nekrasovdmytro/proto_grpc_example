package main

import (
	"context"
	"encoding/json"
	"improve/core"
	pb "improve/proto"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:12345"
)

var clientConn *grpc.ClientConn

func main() {
	api := core.Api{
		Host: ":8080",
	}

	api.HandleHttpFunc("/", func(w http.ResponseWriter, req *http.Request) {
		conn := api.GrpcClientConn(address)
		defer conn.Close()

		client := pb.NewCarCollectionClient(conn)

		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := client.GetList(ctx, &pb.CarRequest{Year: 2018, Limit: 10})

		if err != nil {
			log.Fatalf("Err: %v", err)
		}

		b, err := json.Marshal(r)

		if err != nil {
			log.Fatalf("Json Err: %v", err)
		}

		w.Write(b)
	})
}
