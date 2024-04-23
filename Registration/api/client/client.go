package client

import (
	pb "Registraion/protos/out"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GWU struct {
	id      int
	gameid  int
	userid  int
	seats   int
	payment string
	status  int
}

type SSchedule struct {
	gameid   int
	sport    string
	date     string
	time     string
	seats    int
	price    int
	currency string
	action   string
}

type Mschedule struct {
}

type SetSchedule struct {
}

type updates struct {
	MGwu   GWU         //The postgres table GameWithUsers in a Media microservice
	Msch   Mschedule   //The postgres table Schedule in a Media microservice
	Ssch   SSchedule   //The redisDB table Schedule in a Schedule microservice
	SetSch SetSchedule //The posgres table Schedule in a Settings microservice
}

func (s *SSchedule) sch() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewRegistrationClient(conn)
		ctx := context.Background()
		request := &pb.ScheduleServRequest{
			Gameid:   int64(s.gameid),
			Sport:    s.sport,
			Date:     s.date,
			Time:     s.time,
			Seats:    int64(s.seats),
			Price:    int64(s.price),
			Currency: s.currency,
			Action:   s.action,
		}
		_, err = client.UpdSchedule(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
}

func (u *updates) schedule() {
	u.Ssch.sch()
}

func (u *GWU) gwu() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewRegistrationClient(conn)
		ctx := context.Background()
		request := &pb.GWURequest{
			Id:      int64(u.id),
			Userid:  int64(u.userid),
			Gameid:  int64(u.gameid),
			Seats:   int64(u.seats),
			Payment: u.payment,
			Statpay: 0,
			Status:  int64(u.status),
		}
		_, err = client.SetGamesWithUsers(ctx, request)
		if err != nil {
			log.Print(err)
		}
	}
}

func (u *updates) media() {
	//u.Msch.sch()
	u.MGwu.gwu()
}

func (u *updates) settings() {

}

// Creates request to gRPC function and calls it
func Updates(u *updates, dberr error) {
	if dberr != nil {
		u.schedule()
		u.media()
		u.settings()
	} else {
		log.Printf("You have an error %s", dberr)
	}
}
