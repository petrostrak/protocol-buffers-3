# Protocol buffers 3 and gRPC in Golang

## Protocol buffers 3

Protocol Buffers (protobuf) is a fundamental data serialization format. It is leveraged by many top tech companies such as Google and enables micro-services to transfer data in a format that is safe and efficient. 

### Setup protoc
All the info regarding the protocol buffers and Go such as:

* how to install protoc 
* how to generate Go code from a `.proto` file
* and all the `.proto` file structure

can be found [here](https://developers.google.com/protocol-buffers/docs/reference/go-generated).

Find the correct protocol buffers version based on your Linux [distro](https://github.com/protocolbuffers/protobuf/releases).
Example with x64:
```
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
# Unzip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/
# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/
# Optional: change owner
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google
```

### Protobuf Scalar Types
There are 15 different scalar types we can use in protocol buffers. 

A list of the Protocol Buffers well-known types can be found [here](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#index).

| Type | Keyword | Default value |
| :---        |    :----:   |          ---: |
|Number|int32,int64,sint32,sint64,uint32,uint64,fixed32,fixed64,sfixed32,sfixed64,float,double|0|
|Boolean|bool|false|
|String|string|empty string|
|Bytes|bytes|empty byte slice|
|Enums|enum|first value of the enum|

<sub>Strings only accept UTF-8 encoded or 7-bit ASCII.</sub>

#### Integers
* uint32, uint 64: no negative value.
* int32, int64: accept negative value (but not efficient at serializing)
* sint32, sint64: accept negative value (but less efficient at serializing)
* fixed32, sfixed32: 4 bytes 
* fixed64, sfixes64: 8 bytes

<sub>sfixed are signed integers.</sub>

### Enums
Specifically for `Enums` the tag number starts at 0.
```
enum Weekdays {
    WEEKDAY_UNSPECIFIED = 0;
    WEEKDAY_MONDAY = 1;
    WEEKDAY_TUESDAY = 2;
    WEEKDAY_WEDNESDAY = 3;
    WEEKDAY_THURSDAY = 4;
    WEEKDAY_FRIDAY = 5;
    WEEKDAY_SATURDAY = 6;
    WEEKDAY_SUNDAY = 7;
}

message Schedule {
    WeekDays daysOfWeek = 1;
    // ...
}
```
<sub>For enums we generally keep the 0 tag for the first element of the enum which is unspecified.</sub>

### Repeated Fields
`repeated` fields represent lists

```
message Account {
    uint32 id = 1;
    Person person = 2;

    // repeated fields represent lists
    // and is equivalent of var phones []string
    repeated string phones = 3;
}
```

### Nested Messages
We can nest our messages in a `.proto` file as shown below:

```
message City {
    string name = 1;
    uint64 zip_code = 2;
    string country_name = 3;
    
    message Street {
        string street_name = 1;
        City city = 2;

        message Building {
            string building_name = 1;
            uint32 building_number = 2;
            Street street = 3;
        }
    }
}

message Address {
    City city = 1;
    City.Street street = 2;
    City.Street.Building building = 3;
}
```

### Import Messages
We can import messages from different `.proto` files:

```
import "city.proto";
import "street.proto";
import "building.proto";

message Address {
    City city = 1;
    Street street = 2;
    Building building = 3;
}
```

### Import Messages from package
Or we can import messages from different packages:

```
package city;

message City {
    string name = 1;
    uint64 zip_code = 2;
    string country_name = 3;
}
```

```
package street;

import "city.proto";

message Street {
    string street_name = 1;
    city.City city = 2;
}
```

```
package building;

import "street.proto";

message Building {
    string building_name = 1;
    uint32 building_number = 2;
    street.Street street = 3;
}
```

```
package address;

import "city.proto";
import "street.proto";
import "building.proto";

message Address {
    city.City city = 1;
    street.Street street = 2;
    building.Building building = 3;
}
```

## Golang Programming with protobuf
### A simple proto struct example
```
# .proto
message Simple {
    uint32 id = 1;
    bool is_simple = 2;
    string name = 3;
    repeated int32 sample_list = 4;
}

# .go
func getSimple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   true,
		Name:       "Petros Trak",
		SampleList: []int32{1, 2, 3, 4, 5, 6},
	}
}
```

### A complex proto struct example
```
# .proto
message Dummy {
    int32 id = 1;
    string name = 2;
}

message Complex {
    Dummy one_dummy = 1;
    repeated Dummy multi_dummy = 2;
}

# .go
func getComplex() *pb.Complex {
	return &pb.Complex{
		OnDummy: &pb.Dummy{
			Id:   1,
			Name: "Petros Trak",
		},
		MultiDummy: []*pb.Dummy{
			{Id: 2, Name: "Eirini Tour"},
			{Id: 3, Name: "Deppy Bou"},
			{Id: 4, Name: "Giannis Lio"},
		},
	}
}
```

### An enum proto struct example
```
# .proto
enum EyeColor {
    EYE_COLOR_UNSPECIFIED = 0;
    EYE_COLOR_GREEN = 1;
    EYE_COLOR_BLUE = 2;
    EYE_COLOR_BROWN = 3;
}

message Enumeration {
    EyeColor eye_color = 1;
}

# .go
func getEnum() *pb.Enumeration {
	return &pb.Enumeration{
		// EyeColor: pb.EyeColor_EYE_COLOR_BROWN,
		EyeColor: 3,
	}
}
```

### OneOfs
```
# .proto

// oneof either takes a message that is string, 
// or a message that is int32.
message Result {
    oneof result {
        string message = 1;
        uint32 id = 2;
    }
}

# .go
func getOneOf(msg any) {
	switch x := msg.(type) {
	case *pb.Result_Id:
		fmt.Println(msg.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(msg.(*pb.Result_Message).Message)
	default:
		fmt.Printf("msg has unexpected type %v\n", x)
	}
}
```
<sub>oneof cannot be repeated.</sub>

### Maps
```
# .proto

// IdWrapper is the value of the key-value pair in a map.
// So the type of the wrapper will be the type of the value
// of the map.
message IdWrapper {
    uint32 id = 1;
}

// MapExample is a map[string]uint32
message MapExample {
    map<string, IdWrapper> ids = 1;
}

# .go
func getMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"key1": {Id: 1},
			"key2": {Id: 2},
			"key3": {Id: 4},
			"key4": {Id: 5},
		},
	}
}
```
<sub>Maps cannot be repeated and cannot use float, double or enums for keys.</sub>

### Reading and Writing to Disk
#### Write to file with proto
```
func writeToFile(filename string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("cannot serialize to bytes", err)
		return
	}

	err = ioutil.WriteFile(filename, out, 0644)
	if err != nil {
		log.Fatalln("cannot write to file", err)
		return
	}
}
```

#### Read from file with proto
```
func readFromFile(filename string, pb proto.Message) {
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("cannot read file", err)
		return
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("cannot deserialize data", err)
		return
	}
}
```

### Reading and Writing to JSON
#### Encode to JSON
```
func toJSON(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	out, err := option.Marshal(pb)
	if err != nil {
		log.Fatalln("cannot encode to JSON", err)
		return ""
	}
	return string(out)
}
```

#### Decode from JSON
```
func fromJSON(in string, pb proto.Message) {
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	err := option.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatalln("cannot decode JSON to pb", err)
		return
	}
}

func pbfromJSON(json string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(json, message)
	return message
}

{
    // Example, decode JSON to pb.Simple struct.
    msg := pbfromJSON(json, reflect.TypeOf(pb.Simple{}))
    // ...
}
```

### Data Evolution with Protobuf
The most important rule when it comes to updating a `.proto` file is **not to change the number tags** of the existing fields. 
A more detailed guide about updating a message type can be found [here](https://developers.google.com/protocol-buffers/docs/proto3?hl=en#updating).

#### Renaming Fields
We can rename fields freely when we want to update a `.proto` file. Remember that the only thing that matters when serializing or deserializing a proto message is the tags, not the name.

#### Removing Fields
If we want to remove an existing field for our newer `.proto` version, we use the keyword `reserved` before the number tag and optionally before the name. In this way, we prevent the tag and name for future use.
```
message Example {
	reserved 2,3,9 to 11; // 9 and 11 including
	reserved "first_name", "last_name";
	uint32 id = 1;
}
```
### --decode_raw option
The `--decode_raw` command reads an arbitrary protocol message from standard input and writes the raw tag/value pairs in text format to standard output.
```
# .proto
message Simple {
    uint32 id = 1;
    bool is_simple = 2;
    string name = 3;
    repeated int32 sample_list = 4;
}

# .bin
*My name"

$ cat simple.bin | protoc --decode_raw
1: 42
2: 1
3: "My name"
4: "\001\002\003\004\005\006"
```

### --decode option
The `--decode` command reads a binary message of the given type from standard input and writes it in text format to standard output.  The message type must be defined in PROTO_FILES or their imports.. 
```
# .proto
message Simple {
    uint32 id = 1;
    bool is_simple = 2;
    string name = 3;
    repeated int32 sample_list = 4;
}

# .bin
*My name"

$ cat simple.bin | protoc --decode=Simple simple.proto
id: 42
is_simple: true
name: "My name"
sample_list: 1
sample_list: 2
sample_list: 3
sample_list: 4
sample_list: 5
sample_list: 6
```
<sub>In case the `.proto` file is defined in a package e.g. simple, then the --decode parameter changes. E.g. `cat simple.bin | protoc --decode=simple.Simple simple.proto`</sub>

### --encode
The `--encode` command reads a text-format message of the given type from standard input and writes it in binary to standard output.
```
$ cat simple.bin | protoc --decode=Simple simple.proto > simple.txt
id: 42
is_simple: true
name: "My name"
sample_list: 1
sample_list: 2
sample_list: 3
sample_list: 4
sample_list: 5
sample_list: 6

$ cat simple.bin | protoc --decode=Simple simple.proto > simple.txt
$ cat simple.txt | protoc --encode=Simple simple.proto > simple.pb

# simple.pb
*My name"

# To check that the encoded message is the same with the source
$ diff simple.bin simple.pb
```

### Protocol Buffers Options
Protocol Buffers Options are defined in the [descriptor.proto](https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/descriptor.proto) file.

The `descriptor.proto` defines the metadata for the proto files, messages, enums et cetera.

### Services
A `service` is quite generic in protocol buffers and it's not designed for serialization or deserialization. Instead it is designed for *communication*.

Services are a set of endpoints that are defining an API. A contract for RPC framework. Then with this framework we send serialized messages and these messages will automatically  be deserialized into objects that we can use on the server side or on the client side.
```
message GetSomethingRequest {
	//...
}
message GetSomethingResponse {
	//...
}
message ListSomethingRequest {
	//...
}
message ListSomethingResponse {
	//...
}

service FooService {
	rpc GetSomething(GetSomethingRequest) returns (GetSomethingResponse);
	rpc ListSomething(ListSomethingRequest) returns (ListSomethingResponse);
}
```

## gRPC
- An RPC is a Remote Procedure Call.
- In your CLIENT code, it looks like you're just calling a function directly on the SERVER.

![alt text](https://github.com/petrostrak/protocol-buffers-3-and-gRPC-in-Go/blob/main/imgs/grpc.svg)

### How to get started?
- At the core of gRPC, we need to define the messages and services using [Protocol Buffers](https://developers.google.com/protocol-buffers).
- The rest of the gRPC code will be generated for us and we'll have to provide an implementation for it. 
- One `.proto` file works for over 12 programming languages (server and client), and allows us to use a framework that scales to millions of RPC per seconds.

### Why Protocol Buffers?
- Protocol Buffers are language agnostic.
- Code can be generated for pretty much any language.
- Data is binary and efficiently serialized (small payloads).
- Very convenient for transporting a lot of data.
- Allows for easy API evolution using rules.

### Protocol Buffers role in gRPC
- Protocol Buffers is used to define the:
    * Messages (data, Response and Request).
    * Service (service name and RPC endpoints).

### What is HTTP/2?
- gRPC leverages HTTP/2 as a backbone for communications.
- HTTP/2 is the newer standard for internet communcations that address common pitfall of HTTP/1.1 on modern web pages.

### How HTTP/1.1 works
- HTTP/1.1 opens a new TCP connection to a server at each request.
- It does not compress headers (which are plaintext).
- It only works with Response / Request mechanism (no server push).
- Was originally omposed of two commands:
    * GET to ask for content.
    * POST to send content.

### How HTTP 2 works
- HTTP 2 supports multiplexing
    * The client and server can push messages in parallel over the same TCP connection.
    * This greatly reduces latency.
- HTTP 2 supports server push
    * Servers can push streams (multiple messages) for one request from the client.
    * This saves round trips (latency).
- HTTP 2 supports headers compression.
- HTTP 2 is binary.
- HTTP 2 is secure (SSL is not required but recommended by default).

### 4 Types of API in gRPC

![alt text](https://github.com/petrostrak/protocol-buffers-3-and-gRPC-in-Go/blob/main/imgs/types-of-grpc.png)

- Unary is what a traditional API looks like (HTTP REST).
- HTTP 2 as we've seen, enables APIs to now have streaming capabilities.
- The server and the client can push multiple messages as part of one request.
- In gRPC it's very easy to define these APIs.

### What is an Unary API?
- Unary RPC calls are the basic Response / Request.
- The client sends one message to the server and recieves one response from the server.
- Unary calls are very well suited for small data.
- Start with Unary when writing APIs and use streaming API if performance is an issue.

### What is a Server Streaming API?
- Server Streaming RPC API is a new kind of API enabled thanks to HTTP 2.
- The client will send one message to the server and will receive many responses from the server.
- Streaming Servers are well suited
   * When the server needs to send a lot of data (big data).
   * When the server needs to `PUSH` data to the client continuously without having the client request all over again.
- In gPRC, Server Streaming Calls are defined using the keyword `stream`.
- As for each RPC call, we have to define a Request message and a Response message..

### What is a Client Streaming API?
- Server Streaming RPC API is a new kind of API enabled thanks to HTTP 2.
- The client will send many messages to the server and will receive one response from the server.
- Streaming Clients are well suited
   * When the client needs to send a lot of data (big data).
   * When the server processing is expensive and should happen as the client sends data.
   * When the client needs to `PUSH` data to the server without really expecting a response.
- In gPRC, Client Streaming Calls are defined using the keyword `stream`.
- As for each RPC call, we have to define a Request message and a Response message.

### What is a Bi Directional Streaming API?
- Bi Directional Streaming RPC API is a new kind of API enabled thanks to HTTP 2.
- The client will send many messages to the server and will receive many responses from the server.
- The number of requests and responses does not have to much.
- Bi Directional Streaming RPC is well suited
   * When the client and the server needs to send a lot of data asynchronously.
   * `Chat` protocol.
   * Long running connections.
- In gPRC, Bi Directional Streaming API Calls are defined using the keyword `stream`, twice.
- As for each RPC call, we have to define a Request message and a Response message.

### Scalability in gRPC
- gRPC Servers are asynchronous by default. This means they do not block threads on request. Therefore each gRPC server can serve millions of requests in parallel.
- gRPC Clients can be asynchronous or synchronous (blocking). The client decides which model works best for the performance needs.
- gRPC Clients can perform client side load balancing.

### Security in gRPC
- By default gRPC strongly advocates for the use of SSL in any API.
- Each language will provide an API to load gRPC with the required certificates and provide encryption capability out of the box.
- Additionally using Interceptors, we can also provide authntication.

### gRPC vs REST

| gRPC | REST  |
|---|---|
| Protocol Buffers - smaller, faster  | JSON - text based, slower, bigger  |
| HTTP 2 (low latency)  | HTTP1.1 (higher latency)  |
| Bidirectional & Async  | Client => Server requests only  |
| Stream Support  | Response / Request support only  |
| API Oriented (no constraints - free design)  | CRUD Oriented / POST GET PUT DELETE  |
| Code Generation through Protocol Buffers in any language  | Code Generation though OpenAPI / Swagger (add-on)  |
| RPC based - gRPC does the plumbing  | HTTP verb based - we have to write the plumbing or use 3rd party library  |

### Why use gRPC
- Easy code definition.
- Uses a modern, low latency HTTP 2 transport mechanism.
- SSL security is built in.
- Support for streaming APIs for maximum performance.
- gRPC is API oriented, instead of Resource Oriented like REST.

### Generate Go code with `dummy.proto` file for chapter 07 greet project.
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  greet/proto/dummy.proto
```
Breaking down the above command:

* `protoc` calls the protocol buffer compiler executable.
* `-I` helps protoc to find the import.
* `--go_out=` is where we want to generate the code.
* `--go-grpc_out=` is where we want to generate the gRPC code.
* `----go_opt=paths=source_relative=` enables the gRPC plugin, then code will be generated to support gRPC.
* `--go-grpc_opt=paths=source_relative=` enables the gRPC plugin, then code will be generated to support gRPC.
* `greet/proto/dummy.proto` the path to the `.proto` file.

### Create a minimal gRPC server and client
#### Server
```
var addr = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(l); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
```
#### Client
```
var addr = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect: %v\n", err)
	}
	defer conn.Close()
}
```

### Unary API 
#### Server Implementation
```
func main() {
	...
	pb.RegisterGreetServiceServer(s, &Server{})
	...
}

// Implementation of the Greet API.
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
```

#### Client Implementation
```
func main() {
	...
	c := pb.NewGreetServiceClient(conn)

	greet(c)
}

func greet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Petros",
	})
	if err != nil {
		log.Printf("Could not greet: %v\n", err)
	}

	log.Printf("Greetin: %s", res.Result)
}
```

### Server Streaming API
#### Server Implementation
```
# in .proto file add the streaming server API
service GreetService {
	...
    rpc GreetManyTimes (GreetRequest) returns (stream GreetResponse);
}
```
```
// Implementation of the Streaming Server API.
func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
```
#### Client Implementation
```
func main() {
	...
	doGreetManyTimes(c)
}

func doGreetManyTimes(c pb.GreetServiceClient) {
	req := &pb.GreetRequest{
		FirstName: "Petros",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Printf("error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("error while reading stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
```
### Client Streaming API
#### Server Implementation
```
# in .proto file add the streaming client API
service GreetService {
	...
    rpc LongGreet (stream GreetRequest) returns (GreetResponse);
}
```
```
// Client Streaming Server Implementation
func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		} else if err != nil {
			log.Printf("error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
```
#### Client Implementation
```
func main() {
	...
	doLongGreet(c)
}

func doLongGreet(c pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest{
		{FirstName: "Petros"},
		{FirstName: "Eirini"},
		{FirstName: "Maggie"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Printf("error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)

		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
```
### Bi-Directional Streaming API
#### Server Implementation
```
# in .proto file add the Bi-Directional API
service GreetService {
	...
    rpc GreetEveryone (stream GreetRequest) returns (stream GreetResponse);
}
```
```
func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			log.Printf("error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Printf("error while sending data to client: %v\n", err)
		}
	}
}
```
#### Client Implementation
```
func main() {
	...
	doGreetEveryone(c)
}

func doGreetEveryone(c pb.GreetServiceClient) {
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Printf("error while creating steam: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Petros"},
		{FirstName: "Eirini"},
		{FirstName: "Maggie"},
	}

	waitChan := make(chan struct{})

	// This goroutine sends the requests to the server
	go func() {
		for _, req := range reqs {
			log.Printf("send request: %v\n", req)
			err := stream.Send(req)
			if err != nil {
				log.Printf("error sending request: %v\n", err)
			}
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	// This goroutine waits for the response from the server and
	// ultimately closes the wait-channel.
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Printf("error while receiving: %v\n", err)
			}

			log.Printf("received: %v\n", res.Result)
		}
		close(waitChan)
	}()

	<-waitChan
}
```
### Error Handling in gRPC
To return an gRPC error we make use of the `status` package from `google.golang.org/grpc/status`. This package provides us with many helpful methods, one of which is called `.Errorf(c codes.Code, format string, a ...any)`. 

For this method, we need to provide the error code, e.g. Aborted, AlreadyExists, InvalidArgument using the `codes` package from `google.golang.org/grpc/codes` and a string with the given message of the error.

- [Error Handling 1](https://avi.im/grpc-errors/)
- [Error Handling 2](https://grpc.io/docs/guides/error/)

#### Server-side Error Handling
```
func (s *Server) CalculateSquareRoot(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	number := in.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negatice number: %d", number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(in.Number)),
	}, nil
}
```
#### Client-side Error Handling
In client-side, we can check the errors coming from the server with the use of `.FromError(e errors)` from the `status` package from `google.golang.org/grpc/status`. This method returns the status of the error (code and message that were defined in the server previously) and a bool that tells us whether the error is an gRPC error or not. 
```
func calculateSqrt(c pb.CalculatorServiceClient, n int32) {
	res, err := c.CalculateSquareRoot(context.Background(), &pb.SqrtRequest{
		Number: n,
	})
	if err != nil {
		// ok tells us if it is a gRPC error
		e, ok := status.FromError(err)
		if ok {
			// gRPC error
			log.Printf("error message from server: %s\n", e.Message())
			log.Printf("error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("We probable sent a negative number!")
				return
			}
		} else {
			log.Printf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Result)
}
```
### gRPC Deadlines
- Deadlines allow gRPC clients to specify how long they are willing to wait for an RPC to complete before the RPC is returned with the error DEADLINE_EXCEEDED.
- The gRPC documentation recommends you set a deadline for all client RPC calls.
- The server should check if the deadling has exceeded and calcel the work it is doing.
- Deadlines are propagated across if gRPC calls are chained. 
   * A => B => C  (Deadline for A is passed to B and then passed to C)
- More info on [gRPC Deadlines](https://grpc.io/blog/deadlines/)
#### Server-side with deadlines
```
func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request!")
		}

		// Simulates delay
		time.Sleep(time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName + " !",
	}, nil
}
```
#### Client-side with deadlines
```
func greetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "Petros!"}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			// handle error
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadlline exceeded!")
			} else {
				log.Printf("unexpected gRPC err: %v\n", err)
			}
		} else {
			log.Printf("A non gRPC err: %v\n", err)
		}
	}

	log.Printf("greetWithDeadline: %s\n", res.Result)
}
```
### SSL Encryption in gRPC
- In production, gRPC calls should be running with encryption enabled.
- This is done by generating SSL certificates.
- SSL allows communication to be secure end-to-end and ensuring no Man in the middle attack can be performed.
![alt text](https://github.com/petrostrak/protocol-buffers-3-and-gRPC-in-Go/blob/main/imgs/SSL-in-gRPC.png)

#### Hands on SSL Encryption in gRPC with Go
- Setup a certificate authority.
- Setup a server certificate.
   * Setup the server to use TLS.
- Sign a server certificate.
   * Setup the client to connect securely over TLS.
- More info on [ssl with gRPC](https://grpc.io/docs/guides/auth/)

#### How SSL works
- When you communicate over the internet, your data is visible by all the servers that transfer your packet.
- Any router in the middle can view the packets you're sending using PLAINTEXT.
![alt text](https://github.com/petrostrak/protocol-buffers-3-and-gRPC-in-Go/blob/main/imgs/http.png)

- SSL allows clients and servers to encrypt packets.
![alt text](https://github.com/petrostrak/protocol-buffers-3-and-gRPC-in-Go/blob/main/imgs/https.png)
#### Server Configuration with SSL Encryption
```
func main() {
	
	...
	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		certFile := "ssl/server.ctr"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Printf("failed to load certs")
			os.Exit(1)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	...

}
```
#### Client Configuration with SSL Encryption
```
func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"

		// We are on localhost so we don't need the second parameter
		// serverNameOverride therefore we enter "".
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Printf("error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
}
```