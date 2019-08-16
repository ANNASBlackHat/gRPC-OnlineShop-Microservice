import services_pb2
import services_pb2_grpc

from concurrent import futures
import grpc
import logging
import redis
import time
import json

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

r = redis.Redis(host='localhost', port=6379, db=0)


class Cart(services_pb2_grpc.CartServiceServicer):

    def AddCart(self, request, context):
        user_id = request.user.id
        cart_json = r.get('cart:{}'.format(user_id))
        cart_arr = []
        if cart_json is not None:
            cart_arr = json.loads(cart_json)

        cart_arr.append({"name": request.cart.name, "qty": request.cart.qty})
        cart_json = json.dumps(cart_arr)
        r.set('cart:{}'.format(user_id), cart_json)
        return services_pb2.Response(status=True, message="success")

    def GetCart(self, request, context):
        user_id = request.id
        cart_json = r.get('cart:{}'.format(user_id))

        cart_list = []
        if cart_json is not None:
            cart_arr = json.loads(cart_json)
            cart_list = [services_pb2.Cart(name=c["name"], qty=c["qty"]) for c in cart_arr]
        # for c in cart_arr:
        #     cart_list.append(services_pb2.Cart(name=c["name"], qty=c["qty"]))

        return services_pb2.CartList(cart=cart_list)

    def RemoveCart(self, request, context):
        user_id = request.id
        cart_json = r.delete('cart:{}'.format(user_id))
        return services_pb2.Response(status=True, message="success")


def serve():
    with open('ssl/server.key', 'rb') as f:
        private_key = f.read()

    with open('ssl/server.crt', 'rb') as f:
        certificate_chain = f.read()

    server_credentials = grpc.ssl_server_credentials(((private_key, certificate_chain,),))

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    server.add_secure_port('localhost:50052', server_credentials)
    services_pb2_grpc.add_CartServiceServicer_to_server(Cart(), server)
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    logging.basicConfig()
    serve()
