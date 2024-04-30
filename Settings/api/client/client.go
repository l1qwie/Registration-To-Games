package client

import (
	pb "Settings/protos/out"

	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Upd struct {
	GameId     int
	Sport      string
	Date       int
	Time       int
	Datestr    string
	Timestr    string
	Price      int
	Currency   string
	Seats      int
	Latitude   float64
	Longitude  float64
	Address    string
	Status     int
	Action     string
	UserId     int
	Language   string
	Customlang bool
}

func (u *Upd) schedule() {
	conn, err := grpc.Dial("schedule-app-1:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewScheduleClient(conn)
		ctx := context.Background()
		request := &pb.ScheduleServRequest{
			Gameid:   int64(u.GameId),
			Sport:    u.Sport,
			Date:     u.Datestr,
			Time:     u.Timestr,
			Price:    int64(u.Price),
			Currency: u.Currency,
			Seats:    int64(u.Seats),
			Action:   u.Action,
		}
		log.Print(`Called UpdSchedule:50053 from "Settings"`)
		_, err = client.UpdSchedule(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
}

func (u *Upd) media() {
	conn, err := grpc.Dial("media-app-1:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewMediaClient(conn)
		ctx := context.Background()
		request := &pb.MediaServRequestSch{
			Gameid: int64(u.GameId),
			Sport:  u.Sport,
			Date:   int64(u.Date),
			Time:   int64(u.Time),
			Status: int64(u.Status),
			Action: u.Action,
		}
		log.Print(`Called UpdMediaSch:50055 from "Settings"`)
		_, err = client.UpdMediaSch(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
}

func (u *Upd) reg() {
	conn, err := grpc.Dial("registration-app-1:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewRegistrationClient(conn)
		ctx := context.Background()
		request := &pb.RegServRequest{
			Gameid:   int64(u.GameId),
			Sport:    u.Sport,
			Date:     int64(u.Date),
			Time:     int64(u.Time),
			Seats:    int64(u.Seats),
			Price:    int64(u.Price),
			Currency: u.Currency,
			Status:   int64(u.Status),
			Action:   u.Action,
		}
		log.Print(`Called UpdReg:50054 from "Settings"`)
		_, err = client.UpdReg(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
}

func (u *Upd) user() {
	conn, err := grpc.Dial("user-app-1:50061", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewUserClient(conn)
		ctx := context.Background()
		request := &pb.UpdUsersRequest{
			Userid:     int64(u.UserId),
			Language:   u.Language,
			Customlang: u.Customlang,
		}
		log.Print(`Called UpdUserUsers:50061 from "Settings"`)
		_, err = client.UpdUserUsers(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
}

func Updates(u *Upd, act string, dberr error) {
	if dberr == nil {
		if act == "user" {
			u.user()
		} else {
			u.schedule()
			u.media()
			u.reg()
		}
	} else {
		log.Printf("You have an error %s", dberr)
	}
}
