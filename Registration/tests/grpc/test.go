package main

import (
	apptype "Registraion/apptype"
	"Registraion/fmtogram/types"

	pb "Registraion/protos/out"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	update bool = false
	insert bool = true
)

type server struct {
	pb.UnimplementedRegistrationServer
}

func checkNewGame() bool {
	var count int
	err := types.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		gameId = $1 AND sport = $2 AND date = $3 AND time = $4 AND 
		seats = $5 AND latitude = $6 AND longitude = $7 AND address = $8 AND 
		price = $9 AND currency = $10 AND status = $11`,
		99, "volleyball", 20240225, 1500, 34, 123.1223232,
		55.112399992, "Tel-Aviv Yafo", 100, "USDT", 1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func checkChangeGame() bool {
	var count int
	err := types.Db.QueryRow(`SELECT COUNT(*) FROM Schedule WHERE 
		gameId = $1 AND sport = $2 AND date = $3 AND time = $4 AND 
		seats = $5 AND latitude = $6 AND longitude = $7 AND address = $8 AND 
		price = $9 AND currency = $10 AND status = $11`,
		99, "volleyball", 20240225, 1500, 34, 123.1223232,
		55.112399992, "Tel-Aviv Yafo", 100, "USDT", -1).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func deleteGame() {
	_, err := types.Db.Exec("DELETE FROM Schedule WHERE gameId = $1", 99)
	if err != nil {
		panic(err)
	}
}

func createGame() {
	_, err := types.Db.Exec(`INSERT INTO Schedule (gameId, sport, date, time, seats, latitude, longitude, address, price, currency, status) 
		VALUES (99, 'volleyball', 20240225,1500, 34, 123.1223232, 55.112399992, 'Tel-Aviv Yafo', 100, 'USDT', 1)`)
	if err != nil {
		panic(err)
	}
}

func new(client pb.RegistrationClient, ctx context.Context) {
	defer deleteGame()
	request := &pb.RegServRequest{
		Gameid:    99,
		Sport:     "volleyball",
		Date:      20240225,
		Time:      1500,
		Seats:     34,
		Latitude:  123.1223232,
		Longitude: 55.112399992,
		Address:   "Tel-Aviv Yafo",
		Price:     100,
		Currency:  "USDT",
		Status:    1,
		Action:    insert,
	}
	_, err := client.UpdReg(ctx, request)
	if err != nil {
		panic(err)
	}
	if !checkNewGame() {
		panic("The game wasn't added into the database")
	}
}

func upd(client pb.RegistrationClient, ctx context.Context) {
	createGame()
	defer deleteGame()
	request := &pb.RegServRequest{
		Gameid:    99,
		Sport:     "volleyball",
		Date:      20240225,
		Time:      1500,
		Seats:     34,
		Latitude:  123.1223232,
		Longitude: 55.112399992,
		Address:   "Tel-Aviv Yafo",
		Price:     100,
		Currency:  "USDT",
		Status:    -1,
		Action:    update,
	}
	_, err := client.UpdReg(ctx, request)
	if err != nil {
		panic(err)
	}
	if !checkChangeGame() {
		panic("The game wasn't updated in the database")
	}
}

func testClient() {
	types.Db = apptype.ConnectToDatabase(false)
	conn, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	ctx := context.Background()
	new(client, ctx)
	upd(client, ctx)
	log.Print("All client tests were OK")
}

func startServer() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterRegistrationServer(s, &server{})
	reflection.Register(s)

	log.Println("Server started on port 50052")
	if err := s.Serve(lis); err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}

func (s *server) SetGamesWithUsers(ctx context.Context, req *pb.GWURequest) (*pb.EmptyResponse, error) {
	if req.GetId() != 4 {
		panic((fmt.Sprintf("Id != 1 because Id = %d", req.Id)))
	}
	if req.GetUserid() != 899 {
		panic(fmt.Sprintf("Userid != 899 because Userid = %d", req.GetUserid()))
	}
	if req.Gameid != 23 {
		panic(fmt.Sprintf("Gameid != 23 because Gameid = %d", req.GetGameid()))
	}
	if req.Seats != 66 {
		panic(fmt.Sprintf("Seats != 66 because Seats = %d", req.GetSeats()))
	}
	if req.Status != 1 {
		panic(fmt.Sprintf("Status != 1 because Status = %d", req.GetStatus()))
	}
	if req.Payment != "card" {
		panic(fmt.Sprintf(`Payment != "card" because Payment = %s`, req.GetPayment()))
	}
	if req.Statpay != 0 {
		panic(fmt.Sprintf("req.Statpay != 0 because req.Statpay = %d", req.GetStatpay()))
	}
	return &pb.EmptyResponse{}, nil
}

func testServer() {
	go startServer()
	//client.Updates(4, 899, 23, 66, 1, "card", nil)
	log.Print("The test of server was OK")
}

func main() {
	testClient()
	//testServer()
}
