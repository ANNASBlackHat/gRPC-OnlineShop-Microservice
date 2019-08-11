package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"olshop-microservice/frontend/proto"
	pb "olshop-microservice/services/product-service/proto"
	"strconv"
)

func main() {
	connProduct, err  := grpc.Dial("localhost:1919", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to service product. %v", err)
	}

	defer connProduct.Close()
	clientProduct := pb.NewProductServiceClient(connProduct)

	connCart, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to service cart. %v", err)
	}

	clientCart := proto.NewCartServiceClient(connCart)

	connCheckout, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to service cart. %v", err)
	}

	clientCheckout := proto.NewCheckoutServiceClient(connCheckout)

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
		user := &proto.User{Id:int32(userId)}
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
		cart := &proto.AddCartRequest{User: &proto.User{Id:int32(userId)}, Cart: &proto.Cart{Name: name, Qty:int32(qty)}}
		if resp, err := clientCart.AddCart(ctx, cart); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false,"message": err.Error()})
		}
	})

	g.GET("/checkout/:userId", func(ctx *gin.Context) {
		userId,_ := strconv.Atoi(ctx.Param("userId"))
		if resp, err := clientCheckout.Checkout(ctx, &proto.User{Id:int32(userId)}); err == nil {
			ctx.JSON(http.StatusOK, resp)
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false,"message": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("error - %v", err)
	}
}

