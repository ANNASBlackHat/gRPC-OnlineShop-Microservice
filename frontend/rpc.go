package main

import pb "olshop-microservice/frontend/proto"

func (fe *frontendServer) productClient() pb.ProductServiceClient  {
	return pb.NewProductServiceClient(fe.productSvcConn)
}

func (fe *frontendServer) cartClient() pb.CartServiceClient  {
	return pb.NewCartServiceClient(fe.cartSvcConn)
}

func (fe *frontendServer) checkoutClient() pb.CheckoutServiceClient  {
	return pb.NewCheckoutServiceClient(fe.cartSvcConn)
}