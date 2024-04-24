build:
	protoc -I ./protos --go_out=./protos --go-grpc_out=./protos ./protos/set.proto
path:
	export PATH="$PATH:/bin"
	export PATH="$PATH:/snap/bin"
	export PATH="$PATH:$(go env GOPATH)/bin"