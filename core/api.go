package core

import (
	"bytes"
	"encoding/json"
	"io"
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

func (api *Api) HandleHttpFunc(pattern string, callback func(w http.ResponseWriter, req *http.Request)) {
	http.HandleFunc(pattern, callback)

	log.Fatal(http.ListenAndServe(api.Host, nil))
}

func (api *Api) readerToBuffer(rc io.Reader) bytes.Buffer {
	buf := bytes.Buffer{}
	buf.ReadFrom(rc)

	return buf
}

func (api *Api) ParseJson(r io.Reader, object interface{}) {
	buf := api.readerToBuffer(r)

	json.Unmarshal(buf.Bytes(), object)
}

func (api *Api) HandleJson(w http.ResponseWriter, r interface{}) {
	b, err := json.Marshal(r)

	if err != nil {
		log.Fatalf("Json Err: %v", err)
	}

	w.Write(b)
}
