package servers

import (
	"Media/app/handler"
	"Media/apptype"

	pb "Media/protos/out"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedMediaServer
}

func (s *server) UpdMediaSch(ctx context.Context, req *pb.MediaServRequestSch) (*pb.EmptyResponse, error) {
	g := new(apptype.Game)
	date := int(req.GetDate())
	time := int(req.GetTime())
	act := req.GetAction()
	stat := int(req.GetStatus())
	g.Id = int(req.GetGameid())
	g.Sport = req.GetSport()
	err := handler.UpdateTheSchedule(date, time, stat, g, act)
	return nil, err
}

// Starts a gRPC server
func Start() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterMediaServer(s, &server{})
	reflection.Register(s)

	log.Println("Server started on port 50052")
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
