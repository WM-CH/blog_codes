package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/liuliqiang/blog_codes/rpc/grpc/go"

	"google.golang.org/grpc"
)

var (
	port *int
)

func main() {
	port = new(int)
	flag.IntVar(port, "port", 8080, "port to serve")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, &helloworld.HelloLiqiangIO{})
	grpcServer.Serve(lis)
}
