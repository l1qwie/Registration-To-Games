package fmtogram

import (
	"fmt"
	"log"
	"registrationtogames/bot"
	"registrationtogames/fmtogram/executer"
	"registrationtogames/fmtogram/formatter"
	"registrationtogames/fmtogram/helper"
	"registrationtogames/fmtogram/types"
	"registrationtogames/tests"
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
	fmt.Println("In worker")
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
	var responses chan *types.TelegramResponse
	var requests chan *formatter.Formatter

	responses = make(chan *types.TelegramResponse, 1)
	requests = make(chan *formatter.Formatter, 1)
	responses <- &types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   456,
						IsBot:    false,
						Name:     "John",
						LastName: "Doe",
						Username: "johndoe",
						Language: "ru",
					},
					Text: "/start",
				},
			},
		},
	}
	tests.DataPreparation()
	worker(responses, requests)
	r := <-requests
	tests.AcceptanceOfResults(r)
	fmt.Print("All was alright!")
}
