package core

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

type Api struct {
	Host           string
	grpcClientConn *grpc.ClientConn
}

func (api *Api) GrpcClientConn(address string) *grpc.ClientConn {
	if api.grpcClientConn != nil {
		return api.grpcClientConn
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Err no connection: %v", err)
	}

	return conn
}

func (api *Api) HandleHttpFunc(pattern string, callback func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, callback)

	log.Fatal(http.ListenAndServe(api.Host, nil))
}

func (api *Api) HandleJson(w http.ResponseWriter, r interface{}) {
	b, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("Json Err: %v", err)
	}

	w.Write(b)
}
