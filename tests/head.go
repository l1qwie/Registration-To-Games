package tests

import (
	"RegistrationToGames/fmtogram"
	"RegistrationToGames/fmtogram/errors"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/types"
	"RegistrationToGames/tests/database"
	"RegistrationToGames/tests/othertests"
	preparationdata "RegistrationToGames/tests/preparationData"
	"fmt"
)

func RegToGamesTest() {
	response := make(chan *types.TelegramResponse, 1)
	request := make(chan *formatter.Formatter, 1)
	output := make(chan *types.MessageResponse, 1)
	defer errors.MakeIntestines()
	RegToOneGame(response, request, output) //with one game in database
	fmt.Print("RegToGamesTest was alright!\n")
}

func WelcomeTest() {
	response := make(chan *types.TelegramResponse, 1)
	request := make(chan *formatter.Formatter, 1)
	output := make(chan *types.MessageResponse, 1)
	defer errors.MakeIntestines()
	JustWelcome(response, request, output)
	fmt.Print("WelcomeTest was alright!\n")
}

func SeeTheSchedule() {
	var (
		responses chan *types.TelegramResponse
		requests  chan *formatter.Formatter
		output    chan *types.MessageResponse
	)
	responses = make(chan *types.TelegramResponse, 1)
	requests = make(chan *formatter.Formatter, 1)
	output = make(chan *types.MessageResponse, 1)
	defer errors.MakeIntestines()
	defer database.DeleteUser(userIdT3)
	preparationdata.SeeTheSchedule(responses, output)
	PreparationDatabaseForSchedule()
	fmtogram.Worker(responses, output, requests)
	r := <-requests
	AcceptanceOfResOfSchedule(r)
	fmt.Print("SeeTheSchedule was alright!\n")
}

func MediaTest() {
	responses := make(chan *types.TelegramResponse, 1)
	requests := make(chan *formatter.Formatter, 1)
	output := make(chan *types.MessageResponse, 1)
	defer errors.MakeIntestines()
	UnloadOne(responses, requests, output)  //with two games (to unload and upload)
	UploadOne(responses, requests, output)  //with one games (to upload)
	UnloadAfew(responses, requests, output) //with four games (only to unload)
	UploadAfew(responses, requests, output) //with four games (only to upload)
	fmt.Print("MediaTest was alright!\n")
}

func SettingsTest() {
	responses := make(chan *types.TelegramResponse, 1)
	requests := make(chan *formatter.Formatter, 1)
	output := make(chan *types.MessageResponse, 1)
	defer errors.MakeIntestines()
	ChangeLanguage(responses, requests, output) //test "change language" functionality
	//DeleteGame(responses, requests, output)     //test "delete user game" functionality
	//ChangeSeats()                               //test "change seats on user games" functionality
	//ChagePayment()                              //test "change payment on user games" functionality
}

func JustOtherTests() {
	defer errors.MakeIntestines()
	if !othertests.TestFromIntToStrDate() {
		panic("Date isn't a date")
	}
	fmt.Println("Date is correct")
	if !othertests.TestFromIntToStrTime() {
		panic("Time isn't a time")
	}
	fmt.Println("Time is correct")
}
