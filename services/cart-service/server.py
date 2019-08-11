import cart_service_pb2
import cart_service_pb2_grpc
import messages_pb2

from concurrent import futures
import grpc
import logging
import redis
import time
import json

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

r = redis.Redis(host='localhost', port=6379, db=0)

class Cart(cart_service_pb2_grpc.CartServiceServicer):

    def AddCart(self, request, context):
        user_id = request.user.id
        cart_json = r.get('cart:{}'.format(user_id))
        cart_arr = []
        if (cart_json != None):
            cart_arr = json.loads(cart_json)    

        cart_arr.append({"name": request.cart.name, "qty": request.cart.qty})
        cart_json = json.dumps(cart_arr)
        r.set('cart:{}'.format(user_id), cart_json)
        return cart_service_pb2.Response(status=True, message="success")

    def GetCart(self, request, context):
        user_id = request.id
        cart_json = r.get('cart:{}'.format(user_id))

        cart_list = []
        if cart_json != None:
            cart_arr = json.loads(cart_json)
            cart_list = [messages_pb2.Cart(name=c["name"], qty=c["qty"]) for c in cart_arr]
        # for c in cart_arr:
        #     cart_list.append(messages_pb2.Cart(name=c["name"], qty=c["qty"]))

        return messages_pb2.CartList(cart=cart_list)

    def RemoveCart(self, request, context):
        user_id = request.id
        cart_json = r.delete('cart:{}'.format(user_id))
        return cart_service_pb2.Response(status=True, message="success")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    cart_service_pb2_grpc.add_CartServiceServicer_to_server(Cart(), server)
    server.add_insecure_port('[::]:50052')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    logging.basicConfig()
    serve()
