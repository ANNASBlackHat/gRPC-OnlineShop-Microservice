## Service Architecture
This sample app is composed of manu microservices written in different languages that talk to each other over gRPC.

[![Architecture of
microservices](./img/architecture-diagram.png)](./img/architecture-diagram.png)

**Services List**

| Service | Language | Description |
|---------|----------|-------------|
|[frontend](./frontend) | Go | Exposes an HTTP server |
|[productservice](./services/product-service) | Go | Provides the list of products |
|[cartservice](./services/cart-service) | Python | Stores the items in the user's shipping cart in Redis and retrieves it |
|[checkoutservice](./services/checkout-service) | Node.js | Retrieves user cart, prepares order and orchestrates the payment and shipping detail |
|[shippingservice](./services) | Java | (Not Ready Yet) |


## API Endpoint
| Endpoint | Method | Description |
|---------|----------|-------------|
| /product | GET | Get all product catalog |
| /product/:id | GET | Get product with specific id |
| /cart/:userId | POST | Add cart item for specific user (detail below) |
| /cart/:userId | GET | Get cart item |
| /checkout/:userId | GET | Checkout process |


**Add Cart Params** (form-data)


`name`  : Name of Product

`qty`   : Item quantity


## Installation
Follow the instructions for installation process for each programming languages  https://grpc.io/docs/quickstart/


## Run Program
To run the application, run all services
```$xslt
go run services/product-service/server.go
node services/checkout-service/server.js
python3 services/cart-service/server.py
go run frontend/main.go
```
This app in running on **port: 8080**

## Protobuf Generator

Generate for Go
```$xslt
protoc -I proto/ proto/product-service.proto --go_out=plugins=grpc:services/product-service/proto/
```



Generate for Python
```$xslt
cd services/cart-service
python3 -m grpc_tools.protoc -I../../proto --python_out=. --grpc_python_out=. ../../proto/cart-service.proto
```