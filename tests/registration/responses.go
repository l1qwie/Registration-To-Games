package registration

import "RegistrationToGames/fmtogram/types"

func defTR() *types.TelegramResponse {
	return &types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   477,
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
}

func defMR() *types.MessageResponse {
	return &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 8883,
			Chat: types.Chat{
				Id: 477,
			},
			Photo: []types.Photo{},
		},
	}
}

func Trash() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "O, I want to registrait myself to game!"
	return tr
}

func Trash2() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "!#!@#QWESASED#!#!#$%#3123e1"
	return tr
}

func PresScheduele() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "I don't want to type anything"
	return tr
}

func ChGame() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "2"
	return tr
}

func ChSeats() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "2"
	return tr
}

func ChPayment() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "card"
	return tr
}

func BWishTR() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "Next"
	return tr
}

func MesResp() *types.MessageResponse {
	return defMR()
}

func BwishMR() *types.MessageResponse {
	mr := defMR()
	mr.Result.Photo = []types.Photo{{FileId: "12(!@(U#IOQWEIOPIOJ!@#oijk!@#LIOMASKL:DKLASasdaASDHJKEHUIO$!@#d))"}}
	return mr
}
