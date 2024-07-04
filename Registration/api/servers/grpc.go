package servers

import (
	"context"
	"fmt"
	"log"
	"net"

	"Registration/app/handler"
	"Registration/apptype"

	pb "github.com/l1qwie/Proto-RTG/result"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedRegistrationServer
}

// Create values (date, time, stat [int], act [bool] and g [*apptype.Game]).
// Directs to a function which connects to database and does the stuff
func (s *server) UpdReg(ctx context.Context, req *pb.RegServRequest) (*pb.EmptyResponse, error) {
	var err error
	log.Print("The server UpdReg:50054 was called by someone")
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
	g.Link = req.GetLink()
	con := new(handler.Conn)
	con.Db = apptype.ConnectToDatabase()
	if act == "change" || act == "del" {
		err = con.UpdateTheSchedule(date, time, stat, g)
	} else if act == "new" {
		err = con.AddNewGame(date, time, stat, g)
	}
	log.Print("The server UpdReg:50054 ended its job")
	return nil, err
}

// Starts a gRPC server
func Start() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterRegistrationServer(s, &server{})
	reflection.Register(s)

	log.Println("Server started on port 50054")
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
