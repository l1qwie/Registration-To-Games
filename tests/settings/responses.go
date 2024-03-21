package settings

import "RegistrationToGames/fmtogram/types"

func defTR() *types.TelegramResponse {
	return &types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   899,
						IsBot:    false,
						Name:     "Nataniel",
						LastName: "Sahar",
						Username: "natansah",
						Language: "ru",
					},
				},
			},
		},
	}
	//return tr
}

func defMR() *types.MessageResponse {
	return &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 1111,
			Chat: types.Chat{
				Id: 899,
			},
			Photo: []types.Photo{},
		},
	}
	//return mr
}

func Trash() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "()#!JKASKLDKLAJSLKJSALKD"
	return tr
}

func Trash2() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "PRIVET GUYS!"
	return tr
}

func ChOpt() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "YA HOCHU CHEGO-NIBUD'"
	return tr
}

func WhatOpt() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "records"
	return tr
}

func ChLang() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "en"
	return tr
}

func ChGame() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "1"
	return tr
}

func ChDelGame() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "2"
	return tr
}

func Del() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "del"
	return tr
}

func Change() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "change"
	return tr
}

func Payment() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "payment"
	return tr
}

func Paymethod() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "card"
	return tr
}

func MesResp() *types.MessageResponse {
	return defMR()
}
