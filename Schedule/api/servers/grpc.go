package servers

import (
	"Schedule/app/handler"
	"Schedule/apptype"

	pb "Schedule/protos/out"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedScheduleServer
}

// Create values (date, time [int], and g [*apptype.Game]).
// Directs to a function which connects to database and does the stuff
func (s *server) UpdSchedule(ctx context.Context, req *pb.ScheduleRequest) (*pb.EmptyResponse, error) {
	g := new(apptype.Game)
	g.Date = req.GetDate()
	g.Time = req.GetTime()
	g.Id = int(req.GetGameid())
	g.Sport = req.GetSport()
	g.Action = req.GetAction()
	g.Seats = int(req.GetSeats())
	g.Price = int(req.GetPrice())
	g.Currency = req.GetCurrency()
	err := handler.UpdateTheSchedule(g)
	return nil, err
}

// Starts a gRPC server
func Start() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterScheduleServer(s, &server{})
	reflection.Register(s)

	log.Println("Server started on port 50052")
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
