package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "olshop-microservice/services/product-service/proto"
)

type server struct {}

func main() {
	listener, err := net.Listen("tcp",":1919")
	if err != nil {
		log.Fatalf("failed to listen to port 1919. %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterProductServiceServer(srv, &server{})
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to server : %v", err)
	}
}

func (s *server) GetProduct(ctx context.Context, p *empty.Empty) (*pb.ProductList, error) {
	products := []*pb.Product{
		{Id:1, Name:"Salad Buah", Price:11000},
		{Id:2, Name:"Es Teh", Price:3000},
	}
	return &pb.ProductList{Product: products},nil
}

func (s *server) GetProductById(ctx context.Context, p *pb.ProductRequest) (*pb.Product, error) {
	return &pb.Product{Id: 1, Name:"Salad Buah", Price:11000},nil
}


