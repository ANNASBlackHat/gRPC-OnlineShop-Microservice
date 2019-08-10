import cart_service_pb2
import cart_service_pb2_grpc
import grpc
from concurrent import futures
import logging

import time

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

class Cart(cart_service_pb2_grpc.CartServiceServicer):

    def AddCart(self, request, context):
        return cart_service_pb2.Response(status=True, message="success")

    def GetCart(self, request, context):
        cart = cart_service_pb2.Cart(name='salad', qty=1)
        return cart

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    cart_service_pb2_grpc.add_CartServiceServicer_to_server(Cart(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    logging.basicConfig()
    serve()
