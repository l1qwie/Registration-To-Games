package grpc

import (
	"Registration/apptype"
	"context"
	"log"

	"github.com/l1qwie/Fmtogram/types"
	pb "github.com/l1qwie/Proto-RTG/result"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func new(client pb.RegistrationClient, ctx context.Context) {
	defer deleteGame()
	request := &pb.RegServRequest{
		Gameid:   99,
		Sport:    "volleyball",
		Date:     20240225,
		Time:     1500,
		Seats:    34,
		Price:    100,
		Currency: "USDT",
		Status:   1,
		Action:   "new",
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
		Gameid:   99,
		Sport:    "volleyball",
		Date:     20240225,
		Time:     1500,
		Seats:    34,
		Price:    100,
		Currency: "USDT",
		Status:   -1,
		Action:   "change",
	}
	_, err := client.UpdReg(ctx, request)
	if err != nil {
		panic(err)
	}
	if !checkChangeGame() {
		panic("The game wasn't updated in the database")
	}
}

func testServer() {
	types.Db = apptype.ConnectToDatabase()
	conn, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

func Start() {
	testServer()
}
