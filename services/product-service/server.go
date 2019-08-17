package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	pb "olshop-microservice/services/product-service/proto"
)

type server struct{}

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

	certificate, err := tls.LoadX509KeyPair(
		"services/product-service/ssl/mydomain.com.crt",
		"services/product-service/ssl/mydomain.com.key",
	)

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile("ssl/Root_CA.crt")
	if err != nil {
		log.Fatalf("failed to read client ca cert: %s", err)
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		log.Fatal("failed to append client certs")
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	}

	serverOption := grpc.Creds(credentials.NewTLS(tlsConfig))

	listener, err := net.Listen("tcp", ":1919")
	if err != nil {
		log.Fatalf("failed to listen to port 1919. %v", err)
	}

	srv := grpc.NewServer(serverOption)
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
