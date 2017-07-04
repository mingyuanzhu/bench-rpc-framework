package main

import (
	"bytes"
	"math/rand"
	"time"

	pb "github.com/pdu/bench-rpc-framework/grpc/proto"
	"golang.org/x/net/context"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomChar() rune {
	return letterRunes[rand.Intn(len(letterRunes))]
}

func str(size int) string {
	var buffer bytes.Buffer

	for i := 0; i < size; i++ {
		buffer.WriteString(string(randomChar()))
	}

	return buffer.String()
}

var msg = &pb.GrpcMsg{
	Field1: str(1024),
	Field2: str(2048),
	Field3: str(4096),
	Field4: str(8192),
	Field5: str(16384),
}

// GrpcServer is the server handler
type GrpcServer struct {
}

func newGrpcServer() (*GrpcServer, error) {
	return &GrpcServer{}, nil
}

// Ping implementation
func (g *GrpcServer) Ping(ctx context.Context, input *pb.Empty) (*pb.GrpcMsg, error) {
	return msg, nil
}
