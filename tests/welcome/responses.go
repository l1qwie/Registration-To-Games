package welcome

import "RegistrationToGames/fmtogram/types"

func defTR() *types.TelegramResponse {
	return &types.TelegramResponse{
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
				},
			},
		},
	}
}

func defMR() *types.MessageResponse {
	return &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 8888,
			Chat: types.Chat{
				Id: 456,
			},
			Photo: []types.Photo{},
		},
	}
}

func Trash() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "Hello World ma booooooooy"
	return tr
}

func Trash2() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "Owww, nooo, you are the best human in the world!"
	return tr
}

func Start() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "/start"
	return tr
}

func ShowRules() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "GoReg"
	return tr
}

func ToMainMenu() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "GoNext"
	return tr
}

func MesResp() *types.MessageResponse {
	return defMR()
}
