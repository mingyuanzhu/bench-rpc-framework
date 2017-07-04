package main

import (
	"bytes"
	"math/rand"
	"time"

	"github.com/pdu/bench-rpc-framework/thrift/proto/gen-go/server"
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

var msg = &server.Msg{
	Field1: str(1024),
	Field2: str(2048),
	Field3: str(4096),
	Field4: str(8192),
	Field5: str(16384),
}

type ServerHandler struct {
}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{}
}

func (s *ServerHandler) Ping() (*server.Msg, error) {
	return msg, nil
}
