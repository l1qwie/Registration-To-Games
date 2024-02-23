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
				reg.Reg[index].Ch <- telegramResponse
			} else {
				index = reg.NewIndex()
				reg.Reg[index].UserId = chatID
				reg.Reg[index].Ch = make(chan *types.TelegramResponse, 10)
				fmt.Println("help1", index)
				reg.Reg[index].Ch <- telegramResponse
				fmt.Println("help2")
				go worker(reg.Reg[index].Ch, output)
			}
			offset = offset + 1
		} else if err != nil {
			log.Fatal(err)
		}
	}
}

func worker(input <-chan *types.TelegramResponse, output chan<- *formatter.Formatter) {
	for len(input) > 0 {
		fm := bot.Receiving(<-input)
		fm.Complete()
		output <- fm
	}
	//закрой за мной дверь, я ухожу
}

func pushRequest(requests <-chan *formatter.Formatter) {
	for r := range requests {
		err := r.Send()
		if err != nil {
			log.Fatal(err)
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
	go pushRequest(requests)

	for {
		time.Sleep(time.Second)
	}
}

func StartTests() {
	var (
		counter   int
		responses chan *types.TelegramResponse
		requests  chan *formatter.Formatter
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
			worker(responses, requests)
			r := <-requests
			tests.AcceptanceOfResults(r, counter, i)
		}
		fmt.Printf("Test %d by counter: %d has been completed\n", counter+1, counter)
		counter++
	}
	fmt.Print("All was alright!")
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
