build:
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc -I ./protos --go_out=./protos --go-grpc_out=./protos ./protos/sch.proto