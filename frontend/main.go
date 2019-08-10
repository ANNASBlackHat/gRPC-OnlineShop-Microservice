package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"olshop-microservice/frontend/proto"
	pb "olshop-microservice/services/product-service/proto"
)

func main() {
	connProduct, err  := grpc.Dial("localhost:1919", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to service product. %v", err)
	}

	defer connProduct.Close()
	clientProduct := pb.NewProductServiceClient(connProduct)

	connCart, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to service cart. %v", err)
	}

	clientCart := proto.NewCartServiceClient(connCart)

	g := gin.Default()
	g.GET("/product/:id", func(context *gin.Context) {
		prod := &pb.ProductRequest{}
		if resp, err := clientProduct.GetProductById(context, prod); err == nil {
			context.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
		}else{
			context.JSON(http.StatusInternalServerError, gin.H{"success": false,"message": err.Error()})
		}
	})

	g.GET("/product", func(ctx *gin.Context) {
		req := &empty.Empty{}
		if resp, err := clientProduct.GetProduct(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false,"message": err.Error()})
		}
	})

	g.GET("/cart", func(ctx *gin.Context) {
		if resp, err := clientCart.GetCart(ctx, &empty.Empty{}); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false,"message": err.Error()})
		}
	})

	g.GET("/add-cart", func(ctx *gin.Context) {
		cart := &proto.Cart{Name:"Salad", Qty:1}
		if resp, err := clientCart.AddCart(ctx, cart); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false,"message": err.Error()})
		}
	})


	if err := g.Run(":8080"); err != nil {
		log.Fatalf("error - %v", err)
	}
}

