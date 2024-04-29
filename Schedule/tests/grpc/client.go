package main

import (
	apptype "Schedule/apptype"
	pb "Schedule/protos/out"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func panics(g *apptype.Game, answr *pb.ScheduleServRequest, act string) {
	if g.Id != int(answr.Gameid) {
		panic(fmt.Sprintf(`g.Id != int(answr.Gameid) because g.Id = %d and int(answr.Gameid) = %d`, g.Id, int(answr.Gameid)))
	}
	if g.Sport != answr.Sport {
		panic(fmt.Sprintf(`g.Sport != answr.Sport because g.Sport = %s and answr.Sport = %s`, g.Sport, answr.Sport))
	}
	if g.Date != answr.Date {
		panic(fmt.Sprintf(`g.Date != answr.Date because g.Date = %s and answr.Date = %s`, g.Date, answr.Date))
	}
	if g.Time != answr.Time {
		panic(fmt.Sprintf(`g.Time != answr.Time because g.Time = %s and answr.Time = %s`, g.Time, answr.Time))
	}
	if g.Seats != int(answr.Seats) {
		panic(fmt.Sprintf(`g.Seats != int(answr.Seats) because g.Seats = %d and int(answr.Seats) = %d`, g.Seats, int(answr.Seats)))
	}
	if g.Currency != answr.Currency {
		panic(fmt.Sprintf(`g.Currency != answr.Currency because g.Currency = %s and answr.Currency = %s`, g.Currency, answr.Currency))
	}
	if act != answr.Action {
		panic(fmt.Sprintf(`act != answr.Action because act = %s and answr.Action = %s`, act, answr.Action))
	}

}

func delSch() {
	err := apptype.Client.FlushAll(context.Background()).Err()
	if err != nil {
		panic(err)
	}
}

func createGame() {
	g := new(apptype.Game)
	g.Id = 5
	g.Sport = "volleyball"
	g.Date = "14-06-2025"
	g.Time = "18:00"
	g.Seats = 22
	g.Price = 122
	g.Currency = "USD"
	key := fmt.Sprintf("gameid:%d", g.Id)
	jsonbyt, err := json.Marshal(g)
	if err != nil {
		log.Print(err)
	} else {
		err = apptype.Client.Set(context.Background(), key, jsonbyt, 0).Err()
		if err != nil {
			log.Print(err)
		}
	}
}

func createGame2() {
	g := new(apptype.Game)
	g.Id = 1
	g.Sport = "football"
	g.Date = "14-06-2025"
	g.Time = "18:00"
	g.Seats = 22
	g.Price = 122
	g.Currency = "USD"
	key := fmt.Sprintf("gameid:%d", g.Id)
	jsonbyt, err := json.Marshal(g)
	if err != nil {
		log.Print(err)
	} else {
		err = apptype.Client.Set(context.Background(), key, jsonbyt, 0).Err()
		if err != nil {
			log.Print(err)
		}
	}
}

func ckeckNewGame(answr *pb.ScheduleServRequest) {
	g := new(apptype.Game)
	jsonstr, err := apptype.Client.Get(context.Background(), fmt.Sprintf("gameid:%d", 4)).Result()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(jsonstr), g)
	if err != nil {
		panic(err)
	}
	panics(g, answr, "new")
}

func ckeckDelGame() bool {
	_, err := apptype.Client.Get(context.Background(), fmt.Sprintf("gameid:%d", 5)).Result()
	return err == redis.Nil
}

func ckeckChangeGame(answr *pb.ScheduleServRequest) {
	g := new(apptype.Game)
	jsonstr, err := apptype.Client.Get(context.Background(), fmt.Sprintf("gameid:%d", 1)).Result()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(jsonstr), g)
	if err != nil {
		panic(err)
	}
	panics(g, answr, "change")
}

func newG(client pb.ScheduleClient, ctx context.Context) {
	defer delSch()
	request := &pb.ScheduleServRequest{
		Gameid:   4,
		Sport:    "football",
		Date:     "14-06-2025",
		Time:     "18:00",
		Seats:    22,
		Price:    122,
		Currency: "RUB",
		Action:   "new",
	}
	_, err := client.UpdSchedule(ctx, request)
	if err != nil {
		panic(err)
	}
	ckeckNewGame(request)
}

func del(client pb.ScheduleClient, ctx context.Context) {
	createGame()
	request := &pb.ScheduleServRequest{
		Gameid:   5,
		Sport:    "volleyball",
		Date:     "14-06-2025",
		Time:     "18:00",
		Seats:    22,
		Price:    122,
		Currency: "USD",
		Action:   "del",
	}
	_, err := client.UpdSchedule(ctx, request)
	if err != nil {
		panic(err)
	}
	if !ckeckDelGame() {
		panic("The game wasn't deleted from the redis DB")
	}
}

func upd(client pb.ScheduleClient, ctx context.Context) {
	defer delSch()
	createGame2()
	request := &pb.ScheduleServRequest{
		Gameid:   1,
		Sport:    "football",
		Date:     "14-06-2025",
		Time:     "17:00",
		Seats:    22,
		Price:    122,
		Currency: "RUB",
		Action:   "change",
	}
	_, err := client.UpdSchedule(ctx, request)
	if err != nil {
		panic(err)
	}
	ckeckChangeGame(request)
}

func Client() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewScheduleClient(conn)
	ctx := context.Background()
	apptype.Client = apptype.AddCleint()
	newG(client, ctx)
	del(client, ctx)
	upd(client, ctx)
	log.Print("All client tests were OK")
}

func main() {
	Client()
}
