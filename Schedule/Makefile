build:
	protoc -I ./protos --go_out=./protos --go-grpc_out=./protos ./protos/sch.proto
path:
	export PATH="$PATH:$(go env GOPATH)/bin"
rm:
	docker compose stop \
	&& docker compose rm \
	
up:
	docker compose -f docker-compose.yml up --force-recreate
rb:
	docker build . -t schedule-app
r:
	docker compose stop \
	&& docker compose rm \
	docker compose -f docker-compose.yml up --force-recreate
rb: