#FROM golang:latest

#WORKDIR /app

#COPY go.mod .
#COPY ./ .

#RUN go get
#RUN go build -o bin .

#ENTRYPOINT [ "/app/bin" ]

FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY . .

#RUN go mod download
RUN go get github.com/lib/pq
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]