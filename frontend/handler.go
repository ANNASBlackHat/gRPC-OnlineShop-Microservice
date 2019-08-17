package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"net/http"
	"strconv"
	pb "olshop-microservice/frontend/proto"
)

func (fe *frontendServer) getProduct(ctx *gin.Context)  {
	req := &empty.Empty{}
	if resp, err := fe.productClient().GetProduct(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
	}
}

func (fe *frontendServer) getProductById(ctx *gin.Context)  {
	id, _ := strconv.Atoi(ctx.Param("id"))
	prod := &pb.ProductRequest{Id: int32(id)}
	if resp, err := fe.productClient().GetProductById(ctx, prod); err == nil {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
	}
}

func (fe *frontendServer) addCart(ctx *gin.Context)  {
	userId, _ := strconv.Atoi(ctx.Param("userId"))
	name := ctx.PostForm("name")
	qty, _ := strconv.Atoi(ctx.PostForm("qty"))
	cart := &pb.AddCartRequest{User: &pb.User{Id: int32(userId)}, Cart: &pb.Cart{Name: name, Qty: int32(qty)}}
	if resp, err := fe.cartClient().AddCart(ctx, cart); err == nil {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
	}
}

func (fe *frontendServer) getCart(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("userId"))
	user := &pb.User{Id: int32(userId)}
	if resp, err := fe.cartClient().GetCart(ctx, user); err == nil {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": resp})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
	}
}

func (fe *frontendServer)  checkout(ctx *gin.Context)  {
	userId, _ := strconv.Atoi(ctx.Param("userId"))
	if resp, err := fe.checkoutClient().Checkout(ctx, &pb.User{Id: int32(userId)}); err == nil {
		ctx.JSON(http.StatusOK, resp)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
	}
}
