package grpc

import (
	"fmt"
	"log"
	"net"

	pb "github.com/l1qwie/Proto-RTG/result"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedSettingsServer
}

func testClient() {
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterSettingsServer(s, &server{})
	reflection.Register(s)

	log.Println("Server started on port 50055")
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
