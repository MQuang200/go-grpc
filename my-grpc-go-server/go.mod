module github.com/MQuang200/my-grpc-go-server

go 1.22.2

replace github.com/MQuang200/my-grpc-proto/protogen/hello v0.0.0 => ../my-grpc-proto/protogen/hello

require (
	github.com/MQuang200/my-grpc-proto/protogen/hello v0.0.0
	google.golang.org/grpc v1.64.0
)

require (
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
