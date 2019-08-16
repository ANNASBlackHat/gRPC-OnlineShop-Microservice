package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
	pb "olshop-microservice/frontend/proto"
	"strconv"
)

var cert = "ssl/server.crt"

func main() {
	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	connProduct, err  := grpc.Dial("localhost:1919", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("failed to connect to service product. %v", err)
	}

	defer connProduct.Close()
	clientProduct := pb.NewProductServiceClient(connProduct)

	connCart, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to service cart. %v", err)
	}

	clientCart := pb.NewCartServiceClient(connCart)

	connCheckout, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to service cart. %v", err)
	}

	clientCheckout := pb.NewCheckoutServiceClient(connCheckout)

	g := gin.Default()
	g.GET("/product/:id", func(context *gin.Context) {
		id,_ := strconv.Atoi(context.Param("id"))
		prod := &pb.ProductRequest{Id: int32(id)}
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

	g.GET("/cart/:userId", func(ctx *gin.Context) {
		userId,_ := strconv.Atoi(ctx.Param("userId"))
		user := &pb.User{Id:int32(userId)}
		if resp, err := clientCart.GetCart(ctx, user); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false,"message": err.Error()})
		}
	})

	g.POST("/cart/:userId", func(ctx *gin.Context) {
		userId,_ := strconv.Atoi(ctx.Param("userId"))
		name := ctx.PostForm("name")
		qty,_ := strconv.Atoi(ctx.PostForm("qty"))
		cart := &pb.AddCartRequest{User: &pb.User{Id:int32(userId)}, Cart: &pb.Cart{Name: name, Qty:int32(qty)}}
		if resp, err := clientCart.AddCart(ctx, cart); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false,"message": err.Error()})
		}
	})

	g.GET("/checkout/:userId", func(ctx *gin.Context) {
		userId,_ := strconv.Atoi(ctx.Param("userId"))
		if resp, err := clientCheckout.Checkout(ctx, &pb.User{Id:int32(userId)}); err == nil {
			ctx.JSON(http.StatusOK, resp)
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false,"message": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("error - %v", err)
	}
}

