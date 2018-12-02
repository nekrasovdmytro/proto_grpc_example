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

type CarInputObject struct {
	Type string `json:"type"`
	Year uint64 `json:"year"`
}

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
		var in *CarInputObject
		api.ParseJson(req.Body, &in)

		var carResponse *pb.CarResponse
		ch := make(chan *pb.CarResponse)
		go getListRpc(ch, in)

		carResponse = <-ch

		api.HandleJson(w, carResponse)
	})
}

func getListRpc(ch chan<- *pb.CarResponse, in *CarInputObject) {
	conn := api.GrpcClientConn(address)
	defer conn.Close()

	client := pb.NewCarCollectionClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.GetList(ctx, &pb.CarRequest{Year: in.Year, Type: in.Type, Limit: 10})

	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	ch <- r
}
