package main

import (
	"log"
	"net"

	"goPractice/rpc/server/proto/hello"

	"goPractice/rpc/server/controller/hello_controller"

	"google.golang.org/grpc"
)

const Address = "0.0.0.0:9090"

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatal("fail to listen,err: ", err)
	}

	s := grpc.NewServer()

	hello.RegisterHelloServer(s, &hello_controller.HelloController{})

	log.Println("Listen on: ", Address)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
