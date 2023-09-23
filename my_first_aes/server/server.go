package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"log"
	"context"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"bytes"

	pb "practice/pb/pb_in"
)
// 예제의 server.go 57번째 줄 -> route_guide_grpc.pb.go의 218번째 줄 참고함 
// practice_crypto_grpc.pb.go 65번째 줄 참고, 아무것도 정의 안된 struct를 RegisterAESServer에 집어넣으려니 에러뜸
type server struct{pb.UnimplementedAESServer}

// golang은 외부함수를 가져다 사용하려면 함수이름이 대문자로 시작해야함
func (s *server) EncryptAES(ctx context.Context, blocks *pb.Input) (*pb.Output, error) { // 함수 인자로 왜 포인터를 사용하라는 워닝이 뜰까? -> 정답은 practice_crypto_grpc.pb.go에 있었다. AESServer 를 봐라 그 안에 EncryptAES가 있다
	//EncryptAES passes lock by value: practice/pb/pb_in.Input contains google.golang.org/protobuf/internal/impl.MessageState contains sync.Mutexgo-v
	/*
	`google.golang.org/protobuf/internal/impl.MessageState`는 Go 언어로 작성된 프로토콜 버퍼 라이브러리의 내부 구현 부분을 의미합니다. 이 라이브러리는 구조화된 데이터를 직렬화하고 역직렬화하는 데 사용되며, 주로 시스템 간 통신 및 데이터 저장에 활용됩니다.

이 특정 패키지인 `impl`에서는 프로토콜 버퍼의 내부 구현 세부 사항과 관련된 `MessageState` 타입이 있습니다. 이 타입은 여러 고루틴에 의한 동시 접근을 허용하는 공유 리소스에 대한 접근을 제어하기 위해 사용되는 Go의 동기화 도구인 `sync.Mutex`를 포함하고 있습니다.
server.go 예제에서 60번째 줄 보니까 "sync.Mutex // protects routeNotes " 이걸 쓰긴하는데, 136줄 보면 Lock()을 쓰고있다.
`sync.Mutex`의 존재는 `MessageState`에 일종의 내부 상태가 있어서 여러 고루틴에 의한 동시 접근으로 인한 경쟁 상태나 다른 동기화 문제를 방지하기 위해 보호해야 한다는 것을 시사합니다. `sync.Mutex`는 임계 영역을 생성하여 `MessageState`와 관련된 공유 리소스에 대한 액세스 또는 수정을 하나의 고루틴만 할 수 있도록 합니다.
1. 포인터로 전달: EncryptAES 함수에 변수의 포인터를 전달하도록 수정합니다.
2. EncryptAES 함수를 메서드로 변경하고(예제처럼), 해당 메서드를 수신자로 사용하도록 만듭니다. "sync" 패킼지 사용
*/

// 함수 인자의 ctx는 예제에서도 추가해주더라
	log.Printf("지금 서버에서 Encryption 실행중이고 받아온 평문은 %s입니다.\n\n",blocks.Text)
	block, err := aes.NewCipher(blocks.Key) // message안의 key -> Key로 써야되네
	
	if err != nil {
		return &pb.Output{Text : nil}, err
	}
	
	// 패딩된 블록 크기로 plaintext를 확장합니다.
	blockSize := block.BlockSize()
	plaintext := PKCS7Padding(blocks.Text, blockSize)

	// IV 생성
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return &pb.Output{Text : nil}, err
	}

	// 암호화
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return &pb.Output{Text : ciphertext}, nil
}

func (s *server) DecryptAES(ctx context.Context, blocks *pb.Input) (*pb.Output, error) {
	block, err := aes.NewCipher(blocks.Key)
	log.Printf("지금 서버에서 Decryption 실행중이고 받아온 암호문은 %v입니다.\n\n",blocks.Text)
	//var dst pb.Output
	if err != nil {
		
		return &pb.Output{Text : nil}, err
	}

	// IV 추출
	iv := blocks.Text[:aes.BlockSize]
	ciphertext := blocks.Text[aes.BlockSize:]

	// 복호화
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// 패딩 제거
	ciphertext = PKCS7Unpadding(ciphertext)

	return &pb.Output{Text : ciphertext}, nil
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func PKCS7Unpadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}


func main() {
	ln, err := net.Listen("tcp", ":8081")
	
	if err != nil {
		//err
		fmt.Println("fail server")
		return
	}
	// ln.Accept() -> gRPC 대신 net 패키지로 받아도 연결은 된다.
	
	// 예제의 routeguide/route_guide_grpc.pb.go 242와 server.go  240줄부터 참고
	grpcServer := grpc.NewServer()
	pb.RegisterAESServer(grpcServer, &server{})
	grpcServer.Serve(ln)
	
}