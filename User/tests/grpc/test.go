package grpc

import (
	"User/apptype"
	"User/fmtogram/formatter"
	"User/fmtogram/types"
	"context"
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/l1qwie/Proto-RTG/result"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func testServer() {
	defer deleteUser(999)
	createUser(999)
	conn, err := grpc.NewClient("localhost:50059", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Could not connect: %v", err)
	} else {
		defer conn.Close()
		client := pb.NewUserClient(conn)
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

func DelEnv() {
	DeleteSchedule()
	DeleteEmptyMediaGame(10)
}

func CreateEnv() {
	DelEnv()
	CreateSchedule()
	CreateEmptyMediaGame(10)
}

// Adds data for schedule in array
// Returns this data
func addData() []*apptype.Game {
	lsG := make([]*apptype.Game, 4)
	for k := 0; k < 4; k++ {
		lsG[k] = &apptype.Game{}
		lsG[k].Id = k
		if k == 0 {
			lsG[k].Sport = "volleyball"
			lsG[k].Date = "12-02-2025"
			lsG[k].Time = "12:00"
			lsG[k].Seats = 44
			lsG[k].Price = 100
			lsG[k].Currency = "POUNDS"
		} else if k == 1 {
			lsG[k].Sport = "volleyball"
			lsG[k].Date = "12-06-2026"
			lsG[k].Time = "11:00"
			lsG[k].Seats = 34
			lsG[k].Price = 10
			lsG[k].Currency = "POUNDS"
		} else if k == 2 {
			lsG[k].Sport = "football"
			lsG[k].Date = "12-04-2025"
			lsG[k].Time = "18:00"
			lsG[k].Seats = 14
			lsG[k].Price = 1000
			lsG[k].Currency = "USD"
		} else if k == 3 {
			lsG[k].Sport = "volleyball"
			lsG[k].Date = "02-02-2025"
			lsG[k].Time = "08:00"
			lsG[k].Seats = 77
			lsG[k].Price = 100
			lsG[k].Currency = "RUB"
		}
	}
	return lsG
}

// Put data into Redis DB
func putInRedis(lsg []*apptype.Game, cl *redis.Client) {
	for k := 0; k < 4; k++ {
		key := fmt.Sprintf("gameid:%d", lsg[k].Id)
		jsonbyt, err := json.Marshal(lsg[k])
		if err != nil {
			panic(err)
		}
		err = cl.Set(context.Background(), key, jsonbyt, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

// The main function
// Main logic is here
func CreateSchRedis(cl *redis.Client) {
	lsG := addData()
	putInRedis(lsG, cl)
}

// deletes all schedule from Redis DB
func DelSchRedis(cl *redis.Client) {
	err := cl.FlushAll(context.Background()).Err()
	if err != nil {
		panic(err)
	}
}

func SendMediaGroup() {
	fm := new(formatter.Formatter)
	mg := make([]types.Media, 3)

	mg[0].Type = "photo"
	mg[0].Media = "AgACAgIAAxkDAAIB22YugWU9or1vJw-8aOiJK-anFgPOAAJM2zEbYGpwSYXL0xbw8UhYAQADAgADcwADNAQ"
	mg[1].Type = "photo"
	mg[1].Media = "AgACAgIAAxkDAAIB22YugWU9or1vJw-8aOiJK-anFgPOAAJM2zEbYGpwSYXL0xbw8UhYAQADAgADeQADNAQ"
	fm.AddMapOfMedia(mg)
	//fm.WriteString("HI")
	fm.WriteChatId(738070596)
	_, err := fm.Send()
	if err != nil {
		panic(err)
	}
}
