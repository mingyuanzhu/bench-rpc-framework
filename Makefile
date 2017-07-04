TAG=0.0.1

all: grpc-client grpc-server grpc-gogo-client grpc-gogo-server thrift-client thrift-server

grpc-client:
	docker build --pull=true -t pdu/bench-grpc-client:${TAG} -f Dockerfile_grpc_client .

grpc-server:
	docker build --pull=true -t pdu/bench-grpc-server:${TAG} -f Dockerfile_grpc_server .

grpc-gogo-client:
	docker build --pull=true -t pdu/bench-grpc-gogo-client:${TAG} -f Dockerfile_grpc-gogo_client .

grpc-gogo-server:
	docker build --pull=true -t pdu/bench-grpc-gogo-server:${TAG} -f Dockerfile_grpc-gogo_server .

thrift-client:
	docker build --pull=true -t pdu/bench-thrift-client:${TAG} -f Dockerfile_thrift_client .

thrift-server:
	docker build --pull=true -t pdu/bench-thrift-server:${TAG} -f Dockerfile_thrift_server .

