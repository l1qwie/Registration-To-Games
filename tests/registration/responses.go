package registration

import "registrationtogames/fmtogram/types"

func JustTrash(responses chan *types.TelegramResponse) {
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
}

func JustTrash2(responses chan *types.TelegramResponse) {
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
}

func QueryForPresentationScheduele(responses chan *types.TelegramResponse) {
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
}

func QueryForChooseGame(responses chan *types.TelegramResponse) {
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
}

func QueryForChooseSeats(responses chan *types.TelegramResponse) {
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
}

func QueryForChoosePayment(responses chan *types.TelegramResponse) {
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
}

func QueryForBestWishes(responses chan *types.TelegramResponse) {
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
}
