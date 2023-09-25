# grpc 통신을 통한 AES 암복호화 서비스

## AES 암복호화 테스트 코드

```bash
$ cd AES-test
$ go run aes.go
```

## 모듈 사용법
현재 모듈을 설정해두지 않았기에 모듈 설정이 필수다. -> 한개의 프로젝트로 인식하도록 하기위함
터미널이 `my_first_aes` 디렉터리 위치라고 가정한다. 
```bash
$ go mod init practice # practice 대신 쓰고싶은 이름을 쓰고 client.go에서 import 부분을 바꿔줘도된다.
$ cd server
$ go mod tidy
```