# 통신이 어떻게 이루어지는지 학습
go의 net 패키지와 grpc 패키지를 사용하면 통신은 가능하다.  
하지만 grpc로써의 원하는 동작은 할 수 없다.  
예를 들면 golang으로 구현한 server 코드에서 c++로 작성된 client 코드에 존재하지 않는 함수 A를 가져와 사용하는 것이다.  
이를 위해 Protocol Buffer라는 것이 필요하다.  
공식사이트 - https://protobuf.dev/overview/


# Running
현재 코드는 localhost:8080 으로 통신하도록 고정되어 있음
```shell
PS> cd server
PS> go run server.go
```
다른 터미널 창에서
```shell
PS> go run client.go
```