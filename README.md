# gRPC-golang-practice
golang 학습과 이를 응용한 gRPC 학습  
Windows 11 환경에서 학습  
os는 상관없다  

## Protocol Buffer
```shell
PS> protoc --proto_path=.\ --go_out=.\ --go-grpc_out=.\ .\practice_crypto.proto
```
