package api

import (
	"context"
	"log"

	pb "github.com/l1qwie/Proto-RTG/result"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Update(userId int, language string) error {
	conn, err := grpc.Dial("settings-app-1:50059", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewSettingsClient(conn)
		ctx := context.Background()
		request := &pb.SettingServRequestUser{
			Userid:     int64(userId),
			Language:   language,
			Customlang: false,
		}
		log.Print(`Called UpdSettingUser:50059 from "Welcome"`)
		_, err = client.UpdSettingUser(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
	return err
}
