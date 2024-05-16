package client

import (
	pb "Registration/protos/out"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Upd struct {
	Id        int
	UserId    int
	GameId    int
	GameIdGWU int
	Sport     string
	Date      int
	Time      int
	Datestr   string
	Timestr   string
	Price     int
	Payment   string
	Statpay   int
	Currency  string
	Seats     int
	SeatsGWU  int
	StatusGWU int
	Status    int
	Action    string
	ActionGWU string
}

func (u *Upd) schedule() error {
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
		log.Print(`Called UpdSchedule:50053 from "Registration"`)
		_, err = client.UpdSchedule(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
	return err
}

func (u *Upd) setSchedule() error {
	conn, err := grpc.Dial("settings-app-1:50059", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewSettingsClient(conn)
		ctx := context.Background()
		request := &pb.SettingServRequestSch{
			Gameid:   int64(u.GameId),
			Sport:    u.Sport,
			Date:     int64(u.Date),
			Time:     int64(u.Time),
			Price:    int64(u.Price),
			Currency: u.Currency,
			Seats:    int64(u.Seats),
			Status:   int64(u.Status),
			Action:   u.Action,
		}
		log.Print(`Called UpdSettingSch:50059 from "Registration"`)
		_, err = client.UpdSettingSch(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
	return err
}

func (u *Upd) setGWU() error {
	conn, err := grpc.Dial("settings-app-1:50059", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewSettingsClient(conn)
		ctx := context.Background()
		request := &pb.SettingServRequestGWU{
			Id:      int64(u.Id),
			Userid:  int64(u.UserId),
			Gameid:  int64(u.GameIdGWU),
			Seats:   int64(u.SeatsGWU),
			Payment: u.Payment,
			Statpay: int64(u.Statpay),
			Status:  int64(u.StatusGWU),
			Action:  u.ActionGWU,
		}
		log.Print(`Called UpdSettingGWU:50059 from "Registration"`)
		_, err = client.UpdSettingGWU(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
	return err
}

func (u *Upd) settings() error {
	err := u.setSchedule()
	if err == nil {
		err = u.setGWU()
	}
	return err
}

func (u *Upd) media() error {
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
		log.Print(`Called UpdMediaSch:50055 from "Registration"`)
		_, err = client.UpdMediaSch(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
	return err
}

// Creates request to gRPC function and calls it
func Updates(u *Upd, dberr error) (err error) {
	if dberr == nil {
		err = u.schedule()
		if err == nil {
			err = u.settings()
			if err == nil {
				err = u.media()
			}
		}
	} else {
		log.Printf("You have an error %s", dberr)
	}
	return err
}
