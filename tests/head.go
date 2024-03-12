package tests

import (
	"RegistrationToGames/fmtogram"
	"RegistrationToGames/fmtogram/errors"
	"RegistrationToGames/fmtogram/formatter"
	"RegistrationToGames/fmtogram/types"
	"RegistrationToGames/tests/othertests"
	preparationdata "RegistrationToGames/tests/preparationData"
	"fmt"
)

func GlobalTest() {
	var (
		counter   int
		responses chan *types.TelegramResponse
		requests  chan *formatter.Formatter
		output    chan *types.MessageResponse
	)
	responses = make(chan *types.TelegramResponse, 8)
	requests = make(chan *formatter.Formatter, 8)
	output = make(chan *types.MessageResponse, 8)
	defer errors.MakeIntestines()
	for counter < 8 {
		for i := 0; i < 3; i++ {
			if counter < 3 {
				preparationdata.WelcomeAct(counter, i, responses, output)
			} else if counter >= 3 && counter < 8 {
				preparationdata.RegToGamesAct(counter, i, responses, output)
			} else if counter >= 8 && counter < 9 {
				preparationdata.SeeTheSchedule(counter, responses, output)
			}
			PreparationDatabase(counter)
			fmtogram.Worker(responses, output, requests)
			r := <-requests
			AcceptanceOfResults(r, counter, i)
		}
		fmt.Printf("Global Test %d by counter: %d has been completed\n", counter+1, counter)
		counter++
	}
	fmt.Print("Global Test was alright!\n")
}

func RegToGamesTest() {
	var (
		counter   int
		responses chan *types.TelegramResponse
		requests  chan *formatter.Formatter
		output    chan *types.MessageResponse
	)
	responses = make(chan *types.TelegramResponse, 4)
	requests = make(chan *formatter.Formatter, 4)
	output = make(chan *types.MessageResponse, 4)
	counter = 3
	for counter < 8 {
		for i := 0; i < 3; i++ {
			preparationdata.RegToGamesAct(counter, i, responses, output)
			PreparationDatabaseForRegToGames(counter)
			fmtogram.Worker(responses, output, requests)
			r := <-requests
			AcceptanceOfResOfRegToGames(r, counter, i)
		}
		fmt.Printf("RegToGamesTest %d by counter: %d has been completed\n", counter+1, counter)
		counter++
	}
	fmt.Print("RegToGamesTest was alright!\n")
}

func WelcomeTest() {
	var (
		counter   int
		responses chan *types.TelegramResponse
		requests  chan *formatter.Formatter
		output    chan *types.MessageResponse
	)
	responses = make(chan *types.TelegramResponse, 2)
	requests = make(chan *formatter.Formatter, 2)
	output = make(chan *types.MessageResponse, 2)
	defer errors.MakeIntestines()
	for counter < 3 {
		for i := 0; i < 3; i++ {
			preparationdata.WelcomeAct(counter, i, responses, output)
			PreparationDatabaseForWelcome(counter)
			fmtogram.Worker(responses, output, requests)
			r := <-requests
			AcceptanceOfResOfWelcome(r, counter, i)
		}
		fmt.Printf("WelcomeTest %d by counter: %d has been completed\n", counter+1, counter)
		counter++
	}
	fmt.Print("WelcomeTest was alright!\n")
}

func SeeTheSchedule() {
	var (
		counter   int
		responses chan *types.TelegramResponse
		requests  chan *formatter.Formatter
		output    chan *types.MessageResponse
	)
	counter = 8
	responses = make(chan *types.TelegramResponse, 1)
	requests = make(chan *formatter.Formatter, 1)
	output = make(chan *types.MessageResponse, 1)
	defer errors.MakeIntestines()
	preparationdata.SeeTheSchedule(counter, responses, output)
	PreparationDatabaseForSchedule(counter)
	fmtogram.Worker(responses, output, requests)
	r := <-requests
	AcceptanceOfResOfSchedule(r, counter)
	fmt.Printf("SeeTheSchedule %d by counter: %d has been completed\n", counter+1, counter)
	counter++
	fmt.Print("SeeTheSchedule was alright!\n")
}

func MediaTest() {
	var (
		responses chan *types.TelegramResponse
		requests  chan *formatter.Formatter
		output    chan *types.MessageResponse
	)
	responses = make(chan *types.TelegramResponse, 10)
	requests = make(chan *formatter.Formatter, 10)
	output = make(chan *types.MessageResponse, 10)
	defer errors.MakeIntestines()
	/*counter := 0
	for counter < 3 {
		for j := 0; j < 3; j++ {
			preparationdata.MediaUnload(counter, j, responses, output)
			preparationDatabaseForMediaUnload(counter)
			fmtogram.Worker(responses, output, requests)
			acceptanceOfResOfUnloadMedia(<-requests, counter, j)
		}
		fmt.Printf("UploadTest %d has been complete\n", counter)
		counter++
	}
	*/
	UnloadOne(responses, requests, output) //with two games (to unload and upload)
	UploadOne(responses, requests, output) //with one games (to upload)
	//UnloadALot()
	//UploadALot()
	//UploadNotAll()
	fmt.Print("MediaTest was alright!\n")
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
