package servers

import (
	"Settings/app/handler"
	"Settings/apptype"
	"context"
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

func (s *server) UpdSettingSch(ctx context.Context, req *pb.SettingServRequestSch) (*pb.EmptyResponse, error) {
	log.Print("The server UpdSettingSch:50059 was called by someone")
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
	apptype.Db = apptype.ConnectToDatabase()
	err := handler.UpdateTheSchedule(date, time, stat, g, act)
	log.Print("The server UpdSettingSch:50059 ended its job")
	return nil, err
}

func (s *server) UpdSettingGWU(ctx context.Context, req *pb.SettingServRequestGWU) (*pb.EmptyResponse, error) {
	log.Print("The server UpdSettingGWU:50059 was called by someone")
	g := new(apptype.GWU)
	g.Id = int(req.GetId())
	g.GameId = int(req.GetGameid())
	g.UserId = int(req.GetUserid())
	g.Seats = int(req.GetSeats())
	g.Payment = req.GetPayment()
	g.Statpay = int(req.GetStatpay())
	g.Status = int(req.GetStatus())
	act := req.GetAction()
	apptype.Db = apptype.ConnectToDatabase()
	err := handler.UpdateGWU(g, act)
	log.Print("The server UpdSettingGWU:50059 ended its job")
	return nil, err
}

func (s *server) UpdSettingUser(ctx context.Context, req *pb.SettingServRequestUser) (*pb.EmptyResponse, error) {
	log.Print("The server UpdSettingGWU:50059 was called by someone")
	userId := int(req.GetUserid())
	language := req.GetLanguage()
	customlang := req.GetCustomlang()
	apptype.Db = apptype.ConnectToDatabase()
	err := handler.UpdateUsers(userId, language, customlang)
	log.Print("The server UpdSettingGWU:50059 ended its job")
	return nil, err
}

// Starts a gRPC server
func Start() {
	lis, err := net.Listen("tcp", ":50059")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterSettingsServer(s, &server{})
	reflection.Register(s)

	log.Println("Server started on port 50059")
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
