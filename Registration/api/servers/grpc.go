package servers

import (
	"context"
	"fmt"
	"log"
	"net"

	"Registraion/app/handler"
	apptype "Registraion/apptype"
	"Registraion/fmtogram/types"

	pb "Registraion/protos/out"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedRegistrationServer
}

// Create values (date, time, stat [int], act [bool] and g [*apptype.Game]).
// Directs to a function which connects to database and does the stuff
func (s *server) UpdReg(ctx context.Context, req *pb.RegServRequest) (*pb.EmptyResponse, error) {
	date := int(req.GetDate())
	time := int(req.GetTime())
	stat := int(req.GetStatus())
	act := req.GetAction()
	g := new(apptype.Game)
	g.Id = int(req.GetGameid())
	g.Sport = req.GetSport()
	g.Seats = int(req.GetSeats())
	g.Price = int(req.GetPrice())
	g.Currency = req.GetCurrency()
	g.Address = req.GetAddress()
	g.Lattitude = req.GetLatitude()
	g.Longitude = req.GetLongitude()
	types.Db = apptype.ConnectToDatabase(false)
	err := handler.UpdateTheSchedule(date, time, stat, g, act)
	return nil, err
}

// Starts a gRPC server
func Start() {
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterRegistrationServer(s, &server{})
	reflection.Register(s)

	log.Println("Server started on port 50050")
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
