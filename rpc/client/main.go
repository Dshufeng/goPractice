package main

import (
	"log"

	"goPractice/rpc/client/proto/hello"

	"context"

	"io"

	"google.golang.org/grpc"
)

const Address = "0.0.0.0:9090"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := hello.NewHelloClient(conn)

	res, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "hello"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)

	stream, err := c.LotsOfReplies(context.Background(), &hello.HelloRequest{Name: "hello world"})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("stream.Recv: %v", err)
		}
		log.Println(res)
	}
}
