package servers

import (
	"Settings/app/handler"
	"Settings/apptype"
	"Settings/fmtogram/types"
	pb "Settings/protos/out"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedSettingsServer
}

func (s *server) UpdSettingSch(ctx context.Context, req *pb.SettingServRequestSch) (*pb.EmptyResponse, error) {
	g := new(apptype.Game)
	g.Id = int(req.GetGameid())
	g.Sport = req.GetSport()
	date := int(req.GetDate())
	time := int(req.GetTime())
	stat := int(req.GetStatus())
	act := req.GetAction()
	g.Price = int(req.GetPrice())
	g.Currency = req.GetCurrency()
	g.Seats = int(req.GetSeats())
	apptype.Db = apptype.ConnectToDatabase(false)
	err := handler.UpdateTheSchedule(date, time, stat, g, act)
	return nil, err
}

func (s *server) UpdSettingGWU(ctx context.Context, req *pb.SettingServRequestGWU) (*pb.EmptyResponse, error) {
	g := new(apptype.GWU)
	g.Id = int(req.GetId())
	g.GameId = int(req.GetGameid())
	g.UserId = int(req.GetUserid())
	g.Seats = int(req.GetSeats())
	g.Payment = req.GetPayment()
	g.Statpay = int(req.GetStatpay())
	g.Status = int(req.GetStatus())
	act := req.GetAction()
	types.Db = apptype.ConnectToDatabase(false)
	err := handler.UpdateGWU(g, act)
	return nil, err
}

// Starts a gRPC server
func Start() {
	lis, err := net.Listen("tcp", ":50051")
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
