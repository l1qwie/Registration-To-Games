build:
	protoc -I ./protos --go_out=./protos --go-grpc_out=./protos ./protos/media.proto
path:
	export PATH="$PATH:$(go env GOPATH)/bin"
rm:
	docker compose stop \
	&& docker compose rm \
	&& sudo rm -rf pgdata/
up:
	docker compose -f docker-compose.yml up --force-recreate
rb:
	docker build . -t media-media-app

r:
	docker compose stop \
	&& docker compose rm \
	&& sudo rm -rf pgdata/
	docker compose -f docker-compose.yml up --force-recreate