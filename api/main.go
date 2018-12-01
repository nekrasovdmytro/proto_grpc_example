package main

import (
	"context"
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
var api core.Api

func main() {
	api := core.Api{
		Host: ":8080",
	}

	api.HandleHttpFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var carResponse *pb.CarResponse
		ch := make(chan *pb.CarResponse)
		go getListRpc(ch)

		carResponse = <-ch

		api.HandleJson(w, carResponse)
	})
}

func getListRpc(ch chan<- *pb.CarResponse) {
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

	ch <- r
}
