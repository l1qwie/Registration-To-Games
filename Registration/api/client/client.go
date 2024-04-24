package client

import (
	pb "Registraion/protos/out"
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

func (u *Upd) schedule() {
	conn, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewRegistrationClient(conn)
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
		_, err = client.UpdSchedule(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
}

func (u *Upd) setSchedule() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewRegistrationClient(conn)
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
		_, err = client.UpdSettingSch(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
}

func (u *Upd) setGWU() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewRegistrationClient(conn)
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
		_, err = client.UpdSettingGWU(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
}

func (u *Upd) settings() {
	u.setSchedule()
	u.setGWU()

}

func (u *Upd) media() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewRegistrationClient(conn)
		ctx := context.Background()
		request := &pb.MediaServRequestSch{
			Gameid: int64(u.GameId),
			Sport:  u.Sport,
			Date:   int64(u.Date),
			Time:   int64(u.Time),
			Status: int64(u.Status),
			Action: u.Action,
		}
		_, err = client.UpdMediaSch(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
}

// Creates request to gRPC function and calls it
func Updates(u *Upd, dberr error) {
	if dberr == nil {
		u.schedule()
		u.settings()
		u.media()
	} else {
		log.Printf("You have an error %s", dberr)
	}
}
