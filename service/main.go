package main

import (
	"context"
	pb "improve/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var list []*pb.Car

const (
	port = ":12345"
)

type server struct{}

func (s *server) GetList(ctx context.Context, r *pb.CarRequest) (*pb.CarResponse, error) {

	list = []*pb.Car{
		&pb.Car{Type: "AMMM", Year: 2018},
		&pb.Car{Type: "BMMM", Year: 2017},
		&pb.Car{Type: "CMMM", Year: 2018},
		&pb.Car{Type: "DMMM", Year: 2018},
	}

	var result []*pb.Car

	for _, i := range list {
		if r.Year == i.Year {
			result = append(result, i)
		}
	}

	return &pb.CarResponse{List: result}, nil
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
