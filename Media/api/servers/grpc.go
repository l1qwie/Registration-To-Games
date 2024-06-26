package servers

import (
	"Media/app/handler"
	"Media/apptype"

	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/l1qwie/Proto-RTG/result"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedMediaServer
}

func (s *server) UpdMediaSch(ctx context.Context, req *pb.MediaServRequestSch) (*pb.EmptyResponse, error) {
	log.Print("The server UpdMediaSch:50055 was called by someone")
	g := new(apptype.Game)
	date := int(req.GetDate())
	time := int(req.GetTime())
	act := req.GetAction()
	stat := int(req.GetStatus())
	g.Id = int(req.GetGameid())
	g.Sport = req.GetSport()
	apptype.Db = apptype.ConnectToDatabase()
	err := handler.UpdateTheSchedule(date, time, stat, g, act)
	log.Print("The server UpdMediaSch:50055 ended its job")
	return nil, err
}

// Starts a gRPC server
func Start() {
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterMediaServer(s, &server{})
	reflection.Register(s)

	log.Println("Server started on port 50055")
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
