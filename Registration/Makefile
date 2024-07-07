build:
	protoc -I ./protos --go_out=./protos --go-grpc_out=./protos ./protos/reg.proto
path:
	export PATH="$PATH:$(go env GOPATH)/bin"
rm:
	docker compose stop \
	&& docker compose rm \
	&& sudo rm -rf pgdata/
up:
	docker compose -f docker-compose.yml up --force-recreate
rb:
	docker build . -t registration-reg-app
r:
	docker compose stop \
	&& docker compose rm \
	&& sudo rm -rf pgdata/
	docker compose -f docker-compose.yml up --force-recreate