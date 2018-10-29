generate:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/caquillo07/grpc-demo-shipping-containers/user-service \
    	proto/user/user.proto

build: generate docker-image

docker-image:
	docker build -t user-service .

run:
	docker run -p 50053:50051 \
    	-e MICRO_SERVER_ADDRESS=:50051 \
    	-e MICRO_REGISTRY=mdns vessel-service
