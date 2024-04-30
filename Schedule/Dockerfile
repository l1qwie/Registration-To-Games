FROM golang:1.22.2-bullseye as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN --mount=type=cache,target="/root/.cache/go-build" go build -o bin .

FROM builder as downloader

COPY wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

FROM downloader as final

CMD ["/app/bin"]
