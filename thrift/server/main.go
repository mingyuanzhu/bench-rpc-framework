package main

import (
	"flag"

	"github.com/pdu/bench-rpc-framework/thrift/proto/gen-go/server"

	log "github.com/Sirupsen/logrus"
	"github.com/apache/thrift/lib/go/thrift"
)

var addr = flag.String("serverAddr", ":10002", "The default server address")

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}
	handler := NewServerHandler()
	processor := server.NewMsgServerProcessor(handler)
	thriftServer := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	return thriftServer.Serve()
}

func main() {
	flag.Parse()

	var protocolFactory thrift.TProtocolFactory
	var transportFactory thrift.TTransportFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory = thrift.NewTBufferedTransportFactory(8192)
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	if err := runServer(transportFactory, protocolFactory, *addr); err != nil {
		log.Errorf("runServer failed, ", err)
	}
}
