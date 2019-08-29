protoc-gen-gohttp
=================

[![CircleCI](https://circleci.com/gh/nametake/protoc-gen-gohttp.svg?style=svg)](https://circleci.com/gh/nametake/protoc-gen-gohttp)

protoc-gen-gohttp is a plugin for converting Server's interface generated by protoc-gen-go's gRPC plugin to http.Handler.

In addition to this plugin, you need the protoc command, the proto-gen-go plugin and Google gRPC package.

The code generated by this plugin imports only the standard library, `github.com/golang/protobuf` and `google.golang.org/grpc`.

The converted http.Handler checks Content-Type Header, and changes Marshal/Unmarshal packages. The correspondence table is as follows.

| Content-Type           | package                           |
|------------------------|-----------------------------------|
| application/json       | github.com/golang/protobuf/jsonpb |
| application/protobuf   | github.com/golang/protobuf/proto  |
| application/x-protobuf | github.com/golang/protobuf/proto  |

Install
-------

```console
go get -u github.com/nametake/protoc-gen-gohttp
```

And install dependent tools. (e.g. macOS)

```console
brew install protobuf
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc
```

How to use
----------

```console
protoc --go_out=plugins=grpc:. --gohttp_out=. *.proto
```

Example
-------

### Run

You can execute examples with the following command.

```console
make run_examples
```

You can confirm the operation with the following command.

```console
curl -H "Content-Type: application/json" localhost:8080/sayhello -d '{"name": "john"}'
curl -H "Content-Type: application/json" localhost:8080/greeter/sayhello -d '{"name": "john"}'
```

### Description

Define greeter.proto.

```proto
syntax = "proto3";

package helloworld;

option go_package = "main";

service Greeter {
  rpc SayHello(HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

From greeter.proto you defined, use the following command to generate greeter.pb.go and greeter.http.go.

```console
protoc --go_out=plugins=grpc:. --gohttp_out=. examples/greeter.proto
```

Using the generated Go file, implement as follows.

```go
// EchoGreeterServer has implemented the GreeterServer interface that created from the service in proto file.
type EchoGreeterServer struct {
}

// SayHello implements the GreeterServer interface method.
// SayHello returns a greeting to the name sent.
func (s *EchoGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return &HelloReply{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}, nil
}

func main() {
	// Create the GreeterServer.
	srv := &EchoGreeterServer{}

	// Create the GreeterHTTPConverter generated by protoc-gen-gohttp.
	// This converter converts the GreeterServer interface that created from the service in proto to http.HandlerFunc.
	conv := NewGreeterHTTPConverter(srv)

	// Register SayHello HandlerFunc to the server.
	// If you do not need a callback, pass nil as argument.
	http.Handle("/sayhello", conv.SayHello(logCallback))
	// If you want to create a path from Proto's service name and method name, use the SayHelloWithName method.
	// In this case, the strings 'Greeter' and 'SayHello' are returned.
	http.Handle(restPath(conv.SayHelloWithName(logCallback)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// logCallback is called when exiting ServeHTTP
// and receives Context, ResponseWriter, Request, service argument, service return value and error.
func logCallback(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error) {
	log.Printf("INFO: call %s: arg: {%v}, ret: {%s}", r.RequestURI, arg, ret)
	// YOU MUST HANDLE ERROR
	if err != nil {
		log.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		p := status.New(codes.Unknown, err.Error()).Proto()
		switch r.Header.Get("Content-Type") {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(p)
			if err != nil {
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				return
			}
		case "application/json":
			if err := json.NewEncoder(w).Encode(p); err != nil {
				return
			}
		default:
		}
	}
}

func restPath(service, method string, hf http.HandlerFunc) (string, http.HandlerFunc) {
	return fmt.Sprintf("/%s/%s", strings.ToLower(service), strings.ToLower(method)), hf
}
```

#### HTTPRule

protoc-gen-gohttp supports [google.api.HttpRule](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#httprule) option.

When the Service is defined using HttpRule, Converter implements the `{RpcName}HTTPRule` method. `{RpcName}HTTPRule` method returns Request Method, Path and http.HandlerFunc.

In the following example, Converter implements `GetMessageHTTPRule`. `GetMessageHTTPRule` returns `http.MethodGet`, `"/v1/messages/{message_id}"` and http.HandlerFunc.

```proto
syntax = "proto3";

package example;

option go_package = "main";

import "google/api/annotations.proto";

service Messaging {
  rpc GetMessage(GetMessageRequest) returns (GetMessageResponse) {
    option (google.api.http).get = "/v1/messages/{message_id}";
  }
}

message GetMessageRequest {
  string message_id = 1;
  repeated string tags = 2;
}

message GetMessageResponse {
  string message_id = 1;
  string message = 2;
  repeated string tags = 4;
}
```

`{RpcName}HTTPRule` method is intended for use with HTTP libraries like [go-chi/chi](https://github.com/go-chi/chi) and [gorilla/mux](https://github.com/gorilla/mux) as follows:

```go
type Messaging struct{}

func (m *Messaging) GetMessage(ctx context.Context, req *GetMessageRequest) (*GetMessageResponse, error) {
	return &GetMessageResponse{
		MessageId: req.MessageId,
		Message:   req.Message,
		Tags:      req.Tags,
	}, nil
}

func main() {
	conv := NewMessagingHTTPConverter(&Messaging{})
	r := chi.NewRouter()

	r.Method(conv.GetMessageHTTPRule(nil))

	log.Fatal(http.ListenAndServe(":8080", r))
}
```

protoc-gen-gohttp parses Get Method according to [google.api.HttpRule](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#httprule) option. Therefore, you can pass values to the server in the above example with query string like `/v1/messages/abc1234?message=hello&tags=a&tags=b`.

When you actually execute the above server and execute `curl -H "Content-Type: application/json" "localhost:8080/v1/messages/abc1234?message=hello&tags=a&tags=b"`, the following JOSN is returned.

```json
{
  "messageId": "abc1234",
  "message": "hello",
  "tags": ["a", "b"]
}
```

Callback
--------

Callback is called when the end of the generated code is reached without error or when an error occurs.

Callback is passed HTTP context and http.ResponseWriter and http.Request, RPC arguments and return values, and error.

RPC arguments and return values, and errors may be nil. Here's when nil is passed:

| Timing                                  | RPC argument | RPC return value | error |
|-----------------------------------------|--------------|------------------|-------|
| When an error occurs after calling RPC  | nil          | nil              | err   |
| When RPC returns an error               | arg          | nil              | err   |
| When an error occurs before calling RPC | arg          | ret              | err   |
| When no error occurred                  | arg          | ret              | nil   |

You **MUST HANDLE ERROR** in the callback. If you do not handle it, the error is ignored.

If nil is set, errors are always handled as InternalServerError.

NOT SUPPORTED
-------------

-	Streaming API
	-	Not create a convert method.
-	HttpRule field below
	-	[selector](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule.FIELDS.string.google.api.HttpRule.selector)
	-	[additional_bindings](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule.FIELDS.repeated.google.api.HttpRule.google.api.HttpRule.additional_bindings)
	-	[custom](https://cloud.google.com/endpoints/docs/grpc-service-config/reference/rpc/google.api#google.api.HttpRule.FIELDS.google.api.CustomHttpPattern.google.api.HttpRule.custom)
-	`map` type query string
