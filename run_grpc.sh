TAG=0.0.1
docker run --name=bench-grpc-server -p 10000:10000 -it --rm pdu/bench-grpc-server:${TAG}
