package grpc

import (
	"User/fmtogram/types"
	pb "User/protos/out"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testServer() {
	defer deleteUser(999)
	createUser(999)
	conn, err := grpc.Dial("localhost:50059", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewUsersClient(conn)
		ctx := context.Background()
		request := &pb.UpdUsersRequest{
			Userid:     999,
			Language:   "en",
			Customlang: true,
		}
		_, err := client.UpdUserUsers(ctx, request)
		if err != nil {
			panic(err)
		}
		if !checkChangeLanguage() {
			panic("The language wasn't changed")
		}
	}
	log.Print("Server's tests were successful conpleted")
}

func Start() {
	types.Db = types.ConnectToDatabase(false)
	testServer()
}
