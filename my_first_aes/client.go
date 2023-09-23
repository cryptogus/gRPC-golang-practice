package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "net"

	"encoding/hex"
	"context"
	pb "practice/pb/pb_in"
)

func main() {
	//conn, err := net.Dial("tcp","localhost:8080") -> gRPC는 5계층이므로 기본적으로 3,4계층의 TCP/IP, HTTP 프로토콜로 통신하는데 사용하기 때문에 연결은 가능, gRPC의 기능을 못 쓸뿐
	
	// var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8081",  grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("dial fail")
		return
	}
	defer conn.Close()

	
	//outeguide/route_guide_grpc.pb.go에 New<service>Client 함수 생겨 있음
	client := pb.NewAESClient(conn) // protoc --proto_path=.\ --go_out=.\ --go-grpc_out=.\ .\practice_crypto.proto
	// --go-grpc_out 옵션을 채워줘야 New<proto파일의 service 이름>Client 함수를 사용할 수 있다. .grpc.go 파일이 하나 생김
	// ctx := context.TODO() //https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
	key, _ := hex.DecodeString("2b7e151628aed2a6abf7158809cf4f3c")
	plaintext := []byte("Hello, AES encryption!")
	dst ,err2 := client.EncryptAES(context.Background(), &pb.Input{
		Text: plaintext,
		Key: key,
	}) // 왜 ctx가 들어가야하는가? -> proto파일에서 stream 키워드는 사용하면 어떻게 해결해야하는가?
	
	// 아래 코드도 위와 동일하게 동작함
	// var src pb.Input
	// src.Text = plaintext
	// src.Key = key
	// fmt.Printf("%s",src.Text)
	// var err2 error
	// var dst *pb.Output
	// dst ,err2 = client.EncryptAES(context.Background(), &src)
	
	/*
	https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

	응답을 제공할 때 클라이언트가 응답을 받기 전에 연결을 끊는 등의 상황이 항상 발생할 수 있습니다. 응답을 제공하는 함수가 클라이언트의 연결이 끊어진 것을 알지 못하면 서버 소프트웨어는 사용되지 않을 응답을 계산하는 데 필요한 것보다 더 많은 컴퓨팅 시간을 소비하게 될 수 있습니다.
	이 경우 클라이언트의 연결 상태와 같은 요청의 컨텍스트를 인식하면 클라이언트가 연결을 끊으면 서버가 요청 처리를 중지할 수 있습니다. 이렇게 하면 사용량이 많은 서버의 귀중한 컴퓨팅 리소스를 절약하고 다른 클라이언트의 요청을 처리할 수 있는 여유를 확보할 수 있습니다. 이러한 유형의 정보는 데이터베이스 호출과 같이 함수를 실행하는 데 시간이 걸리는 다른 상황에서도 유용할 수 있습니다. 이러한 유형의 정보에 대한 유비쿼터스 액세스를 지원하기 위해 Go는 표준 라이브러리에 컨텍스트 패키지를 포함했습니다.
	*/

	if err2 != nil{
		fmt.Println("암호화 실패")
		return
	}
	fmt.Println("AES CBC 모드 암호화 결과")
	fmt.Println(dst.Text)
	fmt.Println()
	p, err2 := client.DecryptAES(context.Background(), &pb.Input{
		Text: dst.Text,
		Key: key,
	}) 
	if err2 != nil{
		fmt.Println("복호화 실패")
		return
	}

	fmt.Println("AES CBC 모드 복호화 결과")
	fmt.Printf("%s\n\n",p.Text)
	fmt.Printf("Enter 눌러서 종료\n")
	fmt.Scanln() // wait for Enter Key
	
}

//마지막 퍼즐은 server.go에서 함수가 아니라 method를 사용해야한다는 점이다. (s *server) 을 grpc로 사용하려는 변수들에 붙여줘야한다. 설령 s를 사용하지 않더라도
// 공부를 위함으로 나중에 다시 정리
// https://grpc.io/docs/languages/go/basics/ 문서를 잘 읽어보자