package registration

import "RegistrationToGames/fmtogram/types"

func JustTrash(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "O, I want to registrait myself to game!",
				},
			},
		},
	}
	output <- &types.MessageResponse{
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

func JustTrash2(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "!#!@#QWESASED#!#!#$%#3123e1",
				},
			},
		},
	}
	output <- &types.MessageResponse{
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

func QueryForPresentationScheduele(responses chan<- *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "I don't want to type anything",
				},
			},
		},
	}
	output <- &types.MessageResponse{
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

func QueryForChooseGame(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "2",
				},
			},
		},
	}
	output <- &types.MessageResponse{
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

func QueryForChooseSeats(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "2",
				},
			},
		},
	}
	output <- &types.MessageResponse{
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

func QueryForChoosePayment(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "card",
				},
			},
		},
	}
	output <- &types.MessageResponse{
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

func QueryForBestWishes(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "Next",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 8883,
			Chat: types.Chat{
				Id: 477,
			},
			Photo: []types.Photo{{
				FileId: "12(!@(U#IOQWEIOPIOJ!@#oijk!@#LIOMASKL:DKLASasdaASDHJKEHUIO$!@#d))",
			}},
		},
	}
}
