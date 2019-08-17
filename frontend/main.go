package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
)

type frontendServer struct {
	cartSvcConn     *grpc.ClientConn
	checkoutSvcConn *grpc.ClientConn
	productSvcConn  *grpc.ClientConn
}

func main() {
	svc := new(frontendServer)

	createConnection("localhost:1919", &svc.productSvcConn)
	createConnection("localhost:50052", &svc.cartSvcConn)
	createConnection("localhost:50053", &svc.checkoutSvcConn)

	g := gin.Default()
	g.GET("/product/:id", svc.getProductById)
	g.GET("/product", svc.getProduct)
	g.POST("/cart/:userId", svc.addCart)
	g.GET("/cart/:userId", svc.getCart)
	g.GET("/checkout/:userId", svc.checkout)

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("error - %v", err)
	}
}

func createConnection(address string, conn **grpc.ClientConn) {
	var err error
	*conn, err = grpc.Dial("localhost:1919", grpc.WithInsecure())
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", address))
	}
}
