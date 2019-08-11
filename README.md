

```$xslt
protoc -I proto/ proto/product-service.proto --go_out=plugins=grpc:services/product-service/proto/
```



Generate for python cart
```$xslt
cd services/cart-service
python3 -m grpc_tools.protoc -I../../proto --python_out=. --grpc_python_out=. ../../proto/cart-service.proto
```