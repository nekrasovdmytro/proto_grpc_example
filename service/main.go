package main

import (
	"context"
	pb "improve/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":12345"
)

type server struct{}

func (s *server) GetList(ctx context.Context, r *pb.CarRequest) (*pb.CarResponse, error) {
	var list []*pb.Car
	list = append(list, &pb.Car{Type: "MMMM"})

	return &pb.CarResponse{List: list}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Err : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCarCollectionServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Err : %v", err)
	}
}
