package grpc

import (
	"Settings/apptype"
	pb "Settings/protos/out"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func aNewOneSch(ctx context.Context, cl pb.SettingsClient) {
	defer deleteGame(199)
	request := &pb.SettingServRequestSch{
		Gameid:   199,
		Sport:    "volleyball",
		Date:     20240909,
		Time:     1800,
		Price:    33,
		Currency: "USD",
		Seats:    55,
		Status:   1,
		Action:   "new",
	}
	_, err := cl.UpdSettingSch(ctx, request)
	if err != nil {
		panic(err)
	}
	if !checkANewGameSch() {
		panic("The game wasn't added to the Schedule")
	}
}

func changeSch(ctx context.Context, cl pb.SettingsClient) {
	defer deleteGame(99)
	createGame(99)
	request := &pb.SettingServRequestSch{
		Gameid:   99,
		Sport:    "football",
		Date:     20240909,
		Time:     1900,
		Price:    33,
		Currency: "RUB",
		Seats:    55,
		Status:   -1,
		Action:   "change",
	}
	_, err := cl.UpdSettingSch(ctx, request)
	if err != nil {
		log.Print(err)
	}
	if !checkChangeGameSch() {
		panic("The games hasn't been changed yet")
	}
}

func schedule() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewSettingsClient(conn)
		ctx := context.Background()
		aNewOneSch(ctx, client)
		changeSch(ctx, client)

	}
}

func addGWU(ctx context.Context, cl pb.SettingsClient) {
	defer deleteGWUGame(3)
	request := &pb.SettingServRequestGWU{
		Id:      44,
		Userid:  233,
		Gameid:  3,
		Seats:   5,
		Payment: "card",
		Statpay: 0,
		Status:  1,
		Action:  "new",
	}
	_, err := cl.UpdSettingGWU(ctx, request)
	if err != nil {
		log.Print(err)
	}
	if !ckeckAddGWU() {
		panic("The user's game wasn't added in the GamesWithUsers table")
	}
}

func changeGWU(ctx context.Context, cl pb.SettingsClient) {
	defer deleteGWUGame(8)
	createGWUGame(8)
	request := &pb.SettingServRequestGWU{
		Id:      884,
		Userid:  33333,
		Gameid:  8,
		Seats:   5,
		Payment: "cash",
		Statpay: 1,
		Status:  -1,
		Action:  "change",
	}
	_, err := cl.UpdSettingGWU(ctx, request)
	if err != nil {
		log.Print(err)
	}
	if !ckeckChangeGWU() {
		panic("The user's game wasn't changed")
	}
}

func gamesWithUsers() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewSettingsClient(conn)
		ctx := context.Background()
		addGWU(ctx, client)
		changeGWU(ctx, client)
	}
}

func testServer() {
	apptype.Db = apptype.ConnectToDatabase(false)
	schedule()
	gamesWithUsers()
	log.Print("Server's tests were successfuly completed")
}
