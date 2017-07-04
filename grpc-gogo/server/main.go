package main

import (
	"flag"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/natefinch/lumberjack.v2"

	"google.golang.org/grpc"

	log "github.com/Sirupsen/logrus"
	"github.com/dropbox/godropbox/errors"
	pb "github.com/pdu/bench-rpc-framework/grpc-gogo/proto"
)

var protoPort = flag.String("protoPort", ":10001", "Proto server address")
var logFile = flag.String("logFile", "/var/log/preview/preview.log", "preview log directory")
var logDays = flag.Int("logDays", 28, "the max days to remain log files")
var logCount = flag.Int("logCount", 10, "the max number of log files")

func runProtoServer(ch chan error) {
	lis, err := net.Listen("tcp", *protoPort)
	if err != nil {
		ch <- errors.Wrap(err, "runProto fail")
		return
	}

	grpcServer := grpc.NewServer()
	handler, err := newGrpcServer()
	if err != nil {
		ch <- errors.Wrap(err, "runProtoServer fail")
		return
	}
	pb.RegisterGrpcServerServer(grpcServer, handler)

	log.Info("runProtoServer start")
	err = grpcServer.Serve(lis)
	ch <- errors.Wrap(err, "runProtoServer fail")
}

func main() {
	flag.Parse()

	log.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:   *logFile,
		MaxBackups: *logCount,
		MaxAge:     *logDays, //days
		LocalTime:  true,
	}, os.Stdout))
	log.SetLevel(log.DebugLevel)

	ch := make(chan error)

	go runProtoServer(ch)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err := <-ch:
		log.Fatalf("server fail", err)
	case <-sigs:
		log.Fatal("received quit signal")
	}
}
