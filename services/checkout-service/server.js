var PROTO_PATH = __dirname + '/../../proto/services.proto';

var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var pb = grpc.loadPackageDefinition(packageDefinition).proto;

var client = new pb.CartService('localhost:50052',
    grpc.credentials.createInsecure());

function Checkout(call, callback) {
    const userId = call.request.id
    client.GetCart({id: userId}, function (err, response) {
        console.log(response);
        client.RemoveCart({id:userId}, function(err, response){})
        let resp = {status: false, message: 'cart not found'};
        if (response && response.cart && response.cart.length > 0) {
            resp = {cart: response.cart, shipping: {from: "jakarta", to: "bali", price: 50000}}
        }
        callback(null, resp);
    })
}


function main() {
    var server = new grpc.Server();
    server.addService(pb.CheckoutService.service, {Checkout: Checkout});
    server.bind('0.0.0.0:50053', grpc.ServerCredentials.createInsecure());
    server.start();
}

main()