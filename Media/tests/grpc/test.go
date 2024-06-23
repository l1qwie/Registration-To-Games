package grpc

import (
	"Media/apptype"
	pb "Media/protos/out"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func add(ctx context.Context, cl pb.MediaClient) {
	defer deleteGame(55)
	request := &pb.MediaServRequestSch{
		Gameid: 55,
		Sport:  "football",
		Date:   20201212,
		Time:   1200,
		Action: "new",
		Status: 1,
	}
	_, err := cl.UpdMediaSch(ctx, request)
	if err != nil {
		log.Print(err)
	}
	if !checkAddedGame() {
		panic("The game wasn't added into the Schedule table")
	}
}

func change(ctx context.Context, cl pb.MediaClient) {
	defer deleteGame(55)
	createGame(55)
	request := &pb.MediaServRequestSch{
		Gameid: 55,
		Sport:  "football",
		Date:   20201212,
		Time:   1200,
		Action: "change",
		Status: -1,
	}
	_, err := cl.UpdMediaSch(ctx, request)
	if err != nil {
		log.Print(err)
	}
	if !checkChangeGame() {
		panic("The game wasn't changed")
	}
}

func testServer() {
	apptype.Db = apptype.ConnectToDatabase()
	conn, err := grpc.NewClient("localhost:50057", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewMediaClient(conn)
		ctx := context.Background()
		add(ctx, client)
		change(ctx, client)
	}
	log.Print("Server's tests were successful conpleted")
}

func Start() {
	testServer()
}
