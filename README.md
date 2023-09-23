# gRPC-golang-practice
golang 학습과 이를 응용한 gRPC 학습  

Windows 11 환경에서 학습  

os는 상관없다  
## 공부 방법
2023년 기준으로 golang과 Protocol Buffer는 꾸준히 기능이 업데이트 되고 있기에 시간이 지날수록 새로운 기능이나 기존의 기능이 없어지는 경우가 많다. 정보를 정확히 알수있는 곳은 공식 문서를 보는 것이다. 따라서 url을 따라가서 학습하자.  

golang - https://go.dev/doc/  
Protocol Buffer - https://protobuf.dev/  
grpc - https://grpc.io/docs/what-is-grpc/ , https://grpc.io/docs/languages/go/  

## grpc 예제 설명서
`route_guide` 디렉터리는 grpc 공식 홈페이지에서 제공하는 go언어 기반 예제이다. 이에 대한 설명을 md파일로 정리하였다.  

[grpc 예제 설명](https://github.com/cryptogus/gRPC-golang-practice/blob/main/route_guide/description.md)
## Protocol Buffer
```shell
PS> protoc --proto_path=.\ --go_out=.\ --go-grpc_out=.\ .\practice_crypto.proto
```
