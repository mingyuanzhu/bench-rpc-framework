package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/pdu/bench-rpc-framework/grpc-gogo/proto"
	"google.golang.org/grpc"
)

var server = flag.String("server", ":10001", "The server address")
var qps = flag.Int("qps", 1000, "The qps setting")

func main() {
	flag.Parse()

	tick := time.Tick(time.Millisecond * time.Duration(1000/(*qps)))
	for range tick {
		go func() {
			var opts []grpc.DialOption
			opts = append(opts, grpc.WithInsecure())

			conn, err := grpc.Dial(*server, opts...)
			if err != nil {
				log.Println("grpc Dial fail")
				return
			}
			defer conn.Close()

			client := pb.NewGrpcServerClient(conn)

			msg, err := client.Ping(context.Background(), &pb.Empty{})
			if err != nil {
				log.Println("client Ping fail")
				return
			}

			log.Println(len(msg.Field1), len(msg.Field2), len(msg.Field3), len(msg.Field4), len(msg.Field5))
		}()
	}
}
