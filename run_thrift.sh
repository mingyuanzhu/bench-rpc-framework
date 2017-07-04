TAG=0.0.1
docker run --name=bench-thrift-server -p 10002:10002 -it --rm pdu/bench-thrift-server:${TAG}
