# go를 이용한 gRPC Basic tutorial
모든 정보는 공식 문서를 읽으며 내가 이해하기 편하게 바꾸었음   
공식 사이트 - https://grpc.io/docs/languages/go/basics/  
`.proto` 파일은 client, server 모두 가지고 있어야한다.
## gRPC를 위한 protocol buffer 작성 예제
`routeguide/route_guide.proto` 파일을 살펴보자. 참고로 주석등 필요없는 것들은 삭제하고 코드를 가져왔다.  

- message
    ```proto
    message Point {
        int32 latitude = 1;
        int32 longitude = 2;
    }

    message Feature {
        string name = 1;
        Point location = 2;
    }
    ```
    `message` 키워드는 간단하게 생각하면 gRPC로 함수를 이용할때 사용하는 변수의 모음집(구조체)이라고 생각하면 된다.  

    `Point`라는 위치정보는 위도(`latitude`), 경도(`longitude`) 정보를 나타낸다. 이를 구조체처럼 적어준 것이다.  
    
    자료형과 변수에 할당된 값에 대해서는 설명을 생략한다. (찾아보면 별거아니다)
- service
    ```proto
    service RouteGuide {
        
        rpc GetFeature(Point) returns (Feature) {}
       
        rpc ListFeatures(Rectangle) returns (stream Feature) {}

        rpc RecordRoute(stream Point) returns (RouteSummary) {}

        rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}

    }
    ```
    RPC 프로토콜은 내가 구현하지 않은 함수의 기능을 다른 곳에 구현되어있는 함수를 불러와 사용하는 개념이다. (아마도)  

    따라서 gRPC로 기능을 제공할 함수를 알려줘야한다. 이 예제에서는 service할 함수가 4개다.  
    아까 위에서 `message`는 변수 모음집(구조체)라고 했다.  

    그럼 맨 첫번째 줄을 보자.  
    `GetFeature`라는 이름의 함수가 있고 인자로 `Point`라는 구조체 변수 받아오며 `Feature`라는 구조체 결과 값을 반환해준다는 의미라는 것을 프로그래밍을 하신 분들이라면 쉽게 이해 가능하다.
    근데 `rpc`는 무슨 뜻일까? `rpc`는 stub을 사용하여 함수 기능을 요청하고 결과 값을 받을 수 있도록한다.      
    
    아래는 공식문서의 설명이다.  
    Then you define rpc methods inside your service definition, specifying their request and response types. gRPC lets you define four kinds of service method, all of which are used in the RouteGuide service:

    - A simple RPC where the client sends a request to the server using the stub and waits for a response to come back, just like a normal function call.

    아직 한 가지 의문점이 남아있다. 바로 `stream`이라는 키워드이다. 공식 설명을 보는 것을 추천한다. 내가 이해한대로면 배열 등의 연속한 값들을 사용하는 변수들일 경우 `stream` 키워드를 붙여주는 것 같다.
- build

    From the examples/route_guide directory, run the following command:

    ```shell
    PS> protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    routeguide/route_guide.proto
    ```

    `route_guide.pb.go`와 `route_guide_grpc.pb.go` 파일이 생성된다. 하지만 `route_guide_grpc.pb.go` 파일만 설명한다.    

    공식 문서를 그대로 변역하자면 "클라이언트가 RouteGuide 서비스에 정의된 메서드를 호출할 수 있도록 해주는 인터페이스 기능(또는 stub<스텁>)"과 "서버가 RouteGuide 서비스에 정의된 메서드를 제공할 수 있도록 하는 인터페이스 기능(함수)"가 포함되어 있다.

## Creating the server
server에 protocol buffer에서 적어놓았던 함수들을 구현해야한다. 즉, 해야할 일이 두 가지이다. (번역해서 가져옴)
1. .proto 서비스 정의에서 생성된 서비스 인터페이스 구현: 서비스의 실제 '작업'을 수행합니다.
2. 클라이언트의 요청을 수신하고 올바른 서비스를 전송하기 위해 gRPC 서버를 실행합니다.

- `server.go`
    ```go
    import pb "google.golang.org/grpc/examples/route_guide/routeguide"

    type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
	savedFeatures []*pb.Feature // read-only after initialized

	mu         sync.Mutex // protects routeNotes
	routeNotes map[string][]*pb.RouteNote
    }
    ```
    여기서 `routeGuideServer` 구조체가 정의되어있는데 잘 보면 필드에 `pb.UnimplementedRouteGuideServer` 라는게 있다. 애는 `route_guide_grpc.pb.go` 파일에서 찾아볼 수 있다. 나머지 pb.<변수>는 이와 같다.
    ```go
    // UnimplementedRouteGuideServer must be embedded to have forward compatible implementations.
    type UnimplementedRouteGuideServer struct {
    }
    ```
    그냥 공식 문서 그대로 가져오면  
    our server has a `routeGuideServer` struct type that implements the generated `RouteGuideServer` interface 라고 한다.
    좀 더 친절하게 `route_guide_grpc.pb.go` 파일로 가보면 `RouteGuideServer` 인터페이스가 있다.
    ```go
    // RouteGuideServer is the server API for RouteGuide service.
    // All implementations must embed UnimplementedRouteGuideServer
    // for forward compatibility
    type RouteGuideServer interface {
    	// A simple RPC.
    	//
    	// Obtains the feature at a given position.
    	//
    	// A feature with an empty name is returned if there's no feature at the given
    	// position.
    	GetFeature(context.Context, *Point) (*Feature, error)
    	// A server-to-client streaming RPC.
    	//
    	// Obtains the Features available within the given Rectangle.  Results are
    	// streamed rather than returned at once (e.g. in a response message with a
    	// repeated field), as the rectangle may cover a large area and contain a
    	// huge number of features.
    	ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
    	// A client-to-server streaming RPC.
    	//
    	// Accepts a stream of Points on a route being traversed, returning a
    	// RouteSummary when traversal is completed.
    	RecordRoute(RouteGuide_RecordRouteServer) error
    	// A Bidirectional streaming RPC.
    	//
    	// Accepts a stream of RouteNotes sent while a route is being traversed,
    	// while receiving other RouteNotes (e.g. from other users).
    	RouteChat(RouteGuide_RouteChatServer) error
    	mustEmbedUnimplementedRouteGuideServer()
    }
    ```
    근데 이게 정말 중요한게 method 형식으로 기능을 구현할때 꼭 필요하다. 이것도 공식 문서를 그대로 가져오겠다.  
    The `routeGuideServer` implements all our service methods. Let’s look at the simplest type first, `GetFeature`, which just gets a `Point` from the client and returns the corresponding feature information from its database in a `Feature`.

    그래서 위에 `.proto` 파일에 작성했던 `GetFeature`의 기능을 server에 구현할 때 다음과 같은 형식을 띈다. (simple request and response objects)
    ```go
    func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
      for _, feature := range s.savedFeatures {
        if proto.Equal(feature.Location, point) {
          return feature, nil
        }
      }
      // No feature was found, return an unnamed feature
      return &pb.Feature{Location: point}, nil
    }
    ```
    `(s *routeGuideServer)` 가 보이는가? `s`를 사용하던 안하던 이것을 정의해줘야 한다. 이유는 추가예정
    context 패키지를 사용한 인자도 받아준다. 이는 갑자기 클라이언트에서 기능을 request하고 respone을 받지 않고 통신을 종료할 수 있어서 이다. 사실(아마) 생략해도 되고 clinet에서 사용할 때는 반드시 인자로 들어가야한다. 그냥 서버에서 `func (s *routeGuideServer) GetFeature(point *pb.Point) (*pb.Feature, error)` 이렇게 써도 된다는 설명일뿐 별 의미는 없다.  

    공식문서 설명:
    The method is passed a context object for the RPC and the client’s `Point` protocol buffer request. It returns a Feature protocol buffer object with the response information and an `error`. In the method we populate the `Feature` with the appropriate information, and then return it along with a `nil` error to tell gRPC that we’ve finished dealing with the RPC and that the `Feature` can be returned to the client.  

    아까 `.proto`에서 `stream`에 대해 말 했었다. 그럼 streaming RPC에 대한 예제를 보자. 이 예제는 생략되어있지만 client에서 `ListFeatures` 를 사용할 때 context object 와 pb.Rectangle 변수가 인자로 들어간다.

    ```go
    func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
      for _, feature := range s.savedFeatures {
        if inRange(feature.Location, rect) {
          if err := stream.Send(feature); err != nil {
            return err
          }
        }
      }
      return nil
    }
    ```
    근데 의문이 든다. 그렇다면 `stream pb.RouteGuide_ListFeaturesServer` 이 것은 무엇일까?
    `route_guide_grpc.pb.go` 파일에는 다음과 같이 정의되어 있다.
    ```go
    type RouteGuide_ListFeaturesServer interface {
    	Send(*Feature) error
    	grpc.ServerStream
    }
    ```
    즉, `RouteGuide_ListFeaturesServer` object가 `Send`를 이용하여 response을 해주는 것이다.  
    In the method, we populate as many `Feature` objects as we need to return, writing them to the `RouteGuide_ListFeaturesServer` using its `Send()` method.

    여기까지가 공식문서의 `Simple RPC`,`Server-side streaming RPC`에 대한 설명이고 `Client-side streaming RPC`, `Bidirectional streaming RPC`에 대한 설명은 생략한다. 이정도 했으면 공식문서보고 어느정도 감이 올 것이다.

    이제 server.go의 main 함수이다.
    ```go
    func main() {
    	flag.Parse()
    	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
    	if err != nil {
    		log.Fatalf("failed to listen: %v", err)
    	}
    	var opts []grpc.ServerOption
    	if *tls {
    		if *certFile == "" {
    			*certFile = data.Path("x509/server_cert.pem")
    		}
    		if *keyFile == "" {
    			*keyFile = data.Path("x509/server_key.pem")
    		}
    		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
    		if err != nil {
    			log.Fatalf("Failed to generate credentials: %v", err)
    		}
    		opts = []grpc.ServerOption{grpc.Creds(creds)}
    	}
    	grpcServer := grpc.NewServer(opts...)
    	pb.RegisterRouteGuideServer(grpcServer, newServer())
    	grpcServer.Serve(lis)
    }
    ```
    To build and start a server, we:

    1. Specify the port we want to use to listen for client requests using: `lis, err := net.Listen(...)`. → net 패키지를 사용해한다.
    2. Create an instance of the gRPC server using `grpc.NewServer(...)`. → 여기서 opts는 tls를 사용할때 사용하는 옵션들이고 옵션 사용
    3. Register our service implementation with the gRPC server. → `pb.RegisterRouteGuideServer`의 인자로 위에서 정의했던 `routeGuideServer` 구조체의 주소가 들어간다.
    4. Call `Serve()` on the server with our port details to do a blocking wait until the process is killed or `Stop()` is called.

## `route_guide_grpc.pb.go` 에 정의된 함수들 이름의 정체
이쯤 되면 눈치 챘을테지만 `.proto` 파일을 컴파일 하게되면 자동으로 함수들이 생성된다. 근데 이 함수들의 이름에 규칙이 있다는 생각이 들지 않는가?  예를 들면  `RegisterRouteGuideServer`함수라던가 `UnimplementedRouteGuideServer` 라던가 분명 중복적으로 들어가는게 있다. 바로 `RouteGuide`다.  
이유는 단순한데 `.proto` 파일의 `service` 명을 `RouteGuide` 로 지었기 때문이다. 위에 protocol buffer를 다시한번 보면 알 수 있다. 그럼 만약 `AES`라고 `service` 명을 지으면?? `RouteGuide` 대신 모든 생성되는 것들에 `AES`가 들어갈 것이다. 즉, 규칙은 `Register<서비스 이름>Server` 등 이런식으로 통일 되어있다는 것이다.

##  Creating the client

윗 부분을 이해했다면 여기부터는 공식 설명을 보아도 이해가 충분히 가능하다.

https://grpc.io/docs/languages/go/basics/  
To call service methods, we first need to create a gRPC channel to communicate with the server. We create this by passing the server address and port number to grpc.Dial() as follows:  
```go
func main()
{
    var opts []grpc.DialOption
    ...
    conn, err := grpc.Dial(*serverAddr, opts...)
    if err != nil {
      ...
    }
    defer conn.Close()
}
```
이하 생략
---
이 예시틀이 거의 대부분 활용하게 될 기능들이다.