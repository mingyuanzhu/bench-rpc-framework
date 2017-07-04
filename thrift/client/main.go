package main

import (
	"flag"
	"log"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/pdu/bench-rpc-framework/thrift/proto/gen-go/server"
)

var serverAddr = flag.String("server", ":10000", "The server address")
var qps = flag.Int("qps", 1000, "The qps setting")

func call() {
	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTBufferedTransportFactory(8192)
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(*serverAddr)
	if err != nil {
		log.Println("NewTSocket failed, ", err)
		return
	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		log.Println("transport Open failed, ", err)
		return
	}

	client := server.NewMsgServerClientFactory(transport, protocolFactory)

	msg, err := client.Ping()
	if err != nil {
		log.Println("client Ping failed, ", err)
		return
	}

	log.Println(len(msg.Field1), len(msg.Field2), len(msg.Field3), len(msg.Field4), len(msg.Field5))
}

func main() {
	flag.Parse()

	tick := time.Tick(time.Millisecond * time.Duration(1000/(*qps)))
	for range tick {
		go call()
	}
}
