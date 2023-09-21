package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "net"
)

func main() {
	//conn, err := net.Dial("tcp","localhost:8080") -> gRPC는 5계층이므로 기본적으로 3,4계층의 TCP/IP, HTTP 프로토콜로 통신하는데 사용하기 때문에 연결은 가능, gRPC의 기능을 못 쓸뿐
	conn, err := grpc.Dial("localhost:8080",  grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("fail")
		return
	}
	fmt.Println("success")
	defer conn.Close()
}