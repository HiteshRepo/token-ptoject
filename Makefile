gen:
	protoc --go_out=. --go-grpc_out=. internal/pkg/proto/*.proto