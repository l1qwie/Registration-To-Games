package server

import (
	"User/app/handler"
	"User/apptype"
	pb "User/protos/out"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedUsersServer
}

func (s *server) UpdUserUsers(ctx context.Context, req *pb.UpdUsersRequest) (*pb.EmptyResponse, error) {
	u := new(apptype.User)
	u.Id = int(req.GetUserid())
	u.Language = req.GetLanguage()
	custlang := req.GetCustomlang()
	err := handler.UpdateTheUser(u, custlang)
	return nil, err
}

// Starts a gRPC server
func Start() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterUsersServer(s, &server{})
	reflection.Register(s)

	log.Print("Server started on port 50054")
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
