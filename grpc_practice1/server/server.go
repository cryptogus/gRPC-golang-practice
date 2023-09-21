package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	
	if err != nil {
		//err
		fmt.Println("fail server")
		return
	}
	// ln.Accept() -> gRPC 대신 net 패키지로 받아도 연결은 된다.
	grpcServer := grpc.NewServer()
	grpcServer.Serve(ln)
}