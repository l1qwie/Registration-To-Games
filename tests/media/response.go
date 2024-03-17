package media

import "RegistrationToGames/fmtogram/types"

func defTR() *types.TelegramResponse {
	return &types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   499,
						IsBot:    false,
						Name:     "Bogdan",
						LastName: "Dmitriev",
						Username: "Bogdy",
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
			MessageId: 66666,
			Chat: types.Chat{
				Id: 499,
			},
		},
	}
}

func Trash() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "O, I want to see some photos or videos from some games!"
	return tr
}

func Trash2() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "Please give me photos!"
	return tr
}

func ChDir() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "OK!"
	return tr
}

func ChMUnload() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "unload"
	return tr
}

func ChMUpload() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "upload"
	return tr
}

func ChWaitMediaOne() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "10"
	return tr
}

func ChWaitMediaAfew() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "1"
	return tr
}

func UnlOne() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "10"
	return tr
}

func UnlAfew() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "1"
	return tr
}

func UplOne() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "Take my photo"
	tr.Result[0].Photo[0].FileId = "!@#UIO!@#IOJJKLASEDKLKL#IO!JKLASJKL13419"
	return tr
}

func UplAfew() *types.TelegramResponse {
	tr := defTR()
	tr.Result[0].Message.Text = "Take my photo"
	tr.Result[0].Photo[0].FileId = "!@#UIO!@#IOJJKLASEDKLKL#IO!JKLASJKL13419"
	tr.Result[0].Photo[1].FileId = "!@#UIO!@#IOJJKLIO!JKLASJKL13419"
	tr.Result[0].Photo[2].FileId = "IJ!#JJKLASERJKLIOPEIO*()%*()IOPSDKL:ASDOPK#I!#~!@31313"
	tr.Result[0].Photo[3].FileId = "H!UIO@#HUI!@HASJKLDIOJKL#*()!@()_IOASDEUIO%()_)_"
	return tr
}

func MesResp() *types.MessageResponse {
	return defMR()
}
