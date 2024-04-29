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
	pb.UnimplementedUserServer
}

func (s *server) UpdUserUsers(ctx context.Context, req *pb.UpdUsersRequest) (*pb.EmptyResponse, error) {
	log.Print("The server UpdUserUsers:50061 was called by someone")
	u := new(apptype.User)
	u.Id = int(req.GetUserid())
	u.Language = req.GetLanguage()
	custlang := req.GetCustomlang()
	err := handler.UpdateTheUser(u, custlang)
	log.Print("The server UpdUserUsers:50061 ended its job")
	return nil, err
}

// Starts a gRPC server
func Start() {
	lis, err := net.Listen("tcp", ":50061")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	reflection.Register(s)

	log.Print("Server started on port 50061")
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
