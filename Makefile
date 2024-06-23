build:
	protoc -I ./Protocol --go_out=./Protocol --go-grpc_out=./Protocol ./Protocol/common.proto
path:
	export PATH="$PATH:$(go env GOPATH)/bin"