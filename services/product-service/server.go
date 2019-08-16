package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	pb "olshop-microservice/services/product-service/proto"
)

type server struct{}

var (
	crt = "ssl/server.crt"
	key = "ssl/server.key"
)

var products []*pb.Product

func main() {
	//add fake data
	products = append(products,
		&pb.Product{Id:1, Name: "Es Teh", Price: 3000},
		&pb.Product{Id:2, Name: "Ayam", Price: 10000},
		&pb.Product{Id: 3, Name: "Terong Goreng", Price: 5000},
		&pb.Product{Id: 4, Name: "Kubis Goreng", Price: 4000},
		&pb.Product{Id: 5, Name: "Tempe", Price: 1000},
	)

	listener, err := net.Listen("tcp", ":1919")
	if err != nil {
		log.Fatalf("failed to listen to port 1919. %v", err)
	}

	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile(crt, key)
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	srv := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterProductServiceServer(srv, &server{})
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to server : %v", err)
	}
}

func (s *server) GetProduct(ctx context.Context, p *empty.Empty) (*pb.ProductList, error) {
	return &pb.ProductList{Product: products}, nil
}

func (s *server) GetProductById(ctx context.Context, req *pb.ProductRequest) (*pb.Product, error) {
	var res  = &pb.Product{}
	for _,p := range products {
		if p.Id == req.Id {
			res = p
		}
	}
	return res, nil
}
