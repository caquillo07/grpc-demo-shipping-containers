generate:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/caquillo07/shipper/vessel-service \
    	proto/vessel/vessel.proto

build: generate docker-image

docker-image:
	docker build -t vssel-service .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service

