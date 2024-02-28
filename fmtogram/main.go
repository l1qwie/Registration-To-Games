package fmtogram

import (
	"fmt"
	"log"
	"registrationtogames/bot"
	"registrationtogames/fmtogram/errors"
	"registrationtogames/fmtogram/executer"
	"registrationtogames/fmtogram/formatter"
	"registrationtogames/fmtogram/helper"
	"registrationtogames/fmtogram/types"
	"registrationtogames/tests"
	"registrationtogames/tests/othertests"
	preparationdata "registrationtogames/tests/preparationData"
	"time"
)

func pollResponse(output chan<- *formatter.Formatter, reg *executer.RegTable) {
	var (
		offset           int
		err              error
		telegramResponse *types.TelegramResponse
		index            int
		chatID           int
	)

	err = executer.RequestOffset(types.TelebotToken, &offset)
	for err != nil {
		err = executer.RequestOffset(types.TelebotToken, &offset)
	}
	for {
		telegramResponse = new(types.TelegramResponse)
		err = executer.Updates(&offset, telegramResponse)
		if len(telegramResponse.Result) != 0 && err == nil {
			chatID = helper.ReturnChatId(telegramResponse)
			index = reg.Seeker(chatID)
			if index != executer.None {
				reg.Reg[index].Chu <- telegramResponse
			} else {
				index = reg.NewIndex()
				reg.Reg[index].UserId = chatID
				reg.Reg[index].Chu = make(chan *types.TelegramResponse, 10)
				reg.Reg[index].Chu <- telegramResponse
			}
			go worker(reg.Reg[index].Chu, reg.Reg[index].Chb, output)
			offset = offset + 1
		} else if err != nil {
			log.Fatal(err)
		}
	}
}

func worker(input <-chan *types.TelegramResponse, mesoutput <-chan *types.MessageResponse, output chan<- *formatter.Formatter) {
	var (
		fm         *formatter.Formatter
		userstruct *types.TelegramResponse
		mes        *types.MessageResponse
	)
	for len(input) > 0 {
		userstruct = <-input
		if len(mesoutput) > 0 {
			mes = <-mesoutput
		} else {
			mes = &types.MessageResponse{
				Ok: false,
				Result: types.Message{
					MessageId: 0,
					Chat: types.Chat{
						Id: 0,
					},
					Photo: []types.Photo{{
						FileId: "",
					}},
				},
			}
		}
		fm = bot.Receiving(userstruct, mes)
		fm.Complete()
		output <- fm
	}
}
func pushRequest(requests <-chan *formatter.Formatter, reg *executer.RegTable) {
	for r := range requests {
		mes, err := r.Send()
		if err != nil {
			panic(err)
		}
		if mes.Ok {
			index := reg.Seeker(mes.Result.Chat.Id)
			reg.Reg[index].Chb = make(chan *types.MessageResponse, 10)
			reg.Reg[index].Chb <- mes
		}
	}
}

func StartWithTelegram() {
	var (
		requests chan *formatter.Formatter
		reg      *executer.RegTable
	)
	requests = make(chan *formatter.Formatter)
	reg = new(executer.RegTable)
	go pollResponse(requests, reg)
	go pushRequest(requests, reg)

	for {
		time.Sleep(time.Second)
	}
}

func StartTests() {
	var (
		counter   int
		responses chan *types.TelegramResponse
		requests  chan *formatter.Formatter
		mes       chan *types.MessageResponse
	)
	responses = make(chan *types.TelegramResponse, 8)
	requests = make(chan *formatter.Formatter, 8)

	defer errors.MakeIntestines()
	for counter < 8 {
		for i := 0; i < 3; i++ {
			if counter < 3 {
				preparationdata.WelcomeAct(counter, i, responses)
			} else if counter >= 3 && counter < 8 {
				preparationdata.RegToGamesAct(counter, i, responses)
			}
			tests.PreparationDatabase(counter)
			worker(responses, mes, requests)
			r := <-requests
			tests.AcceptanceOfResults(r, counter, i)
		}
		fmt.Printf("Global Test %d by counter: %d has been completed\n", counter+1, counter)
		counter++
	}
	fmt.Print("Global Test was alright!\n")
}

func RegToGames() {
	var (
		counter   int
		responses chan *types.TelegramResponse
		requests  chan *formatter.Formatter
		mes       chan *types.MessageResponse
	)
	responses = make(chan *types.TelegramResponse, 4)
	requests = make(chan *formatter.Formatter, 4)
	counter = 3
	for counter < 8 {
		for i := 0; i < 3; i++ {
			preparationdata.RegToGamesAct(counter, i, responses)
			tests.PreparationDatabaseForRegToGames(counter)
			worker(responses, mes, requests)
			r := <-requests
			tests.AcceptanceOfResOfRegToGames(r, counter, i)
		}
		fmt.Printf("RegToGamesTest %d by counter: %d has been completed\n", counter+1, counter)
		counter++
	}
	fmt.Print("RegToGamesTest was alright!\n")
}

func Welcome() {
	var (
		counter   int
		responses chan *types.TelegramResponse
		requests  chan *formatter.Formatter
		mes       chan *types.MessageResponse
	)
	responses = make(chan *types.TelegramResponse, 2)
	requests = make(chan *formatter.Formatter, 2)
	defer errors.MakeIntestines()
	for counter < 3 {
		for i := 0; i < 3; i++ {
			preparationdata.WelcomeAct(counter, i, responses)
			tests.PreparationDatabaseForWelcome(counter)
			worker(responses, mes, requests)
			r := <-requests
			tests.AcceptanceOfResOfWelcome(r, counter, i)
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
		mes       chan *types.MessageResponse
	)
	counter = 8
	responses = make(chan *types.TelegramResponse, 1)
	requests = make(chan *formatter.Formatter, 1)
	defer errors.MakeIntestines()
	preparationdata.SeeTheSchedule(counter, responses)
	tests.PreparationDatabaseForSchedule(counter)
	worker(responses, mes, requests)
	r := <-requests
	tests.AcceptanceOfResOfSchedule(r, counter)
	fmt.Printf("SeeTheSchedule %d by counter: %d has been completed\n", counter+1, counter)
	counter++
	fmt.Print("SeeTheSchedule was alright!\n")
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
