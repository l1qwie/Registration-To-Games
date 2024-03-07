package welcome

import "RegistrationToGames/fmtogram/types"

func JustTrash(responses chan<- *types.TelegramResponse, output chan<- *types.MessageResponse) {
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
					Text: "Hello World ma booooooooy",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 8888,
			Chat: types.Chat{
				Id: 456,
			},
			Photo: []types.Photo{{
				FileId: "",
			}},
		},
	}
}

func JustTrash2(responses chan<- *types.TelegramResponse, output chan<- *types.MessageResponse) {
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
					Text: "Owww, nooo, you are the best human in the world!",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 8888,
			Chat: types.Chat{
				Id: 456,
			},
			Photo: []types.Photo{{
				FileId: "",
			}},
		},
	}
}

func Start(responses chan<- *types.TelegramResponse, output chan<- *types.MessageResponse) {
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
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 8888,
			Chat: types.Chat{
				Id: 456,
			},
			Photo: []types.Photo{{
				FileId: "",
			}},
		},
	}
}

func QueryForShowRules(responses chan<- *types.TelegramResponse, output chan<- *types.MessageResponse) {
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
					Text: "GoReg",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 8888,
			Chat: types.Chat{
				Id: 456,
			},
			Photo: []types.Photo{{
				FileId: "",
			}},
		},
	}
}

func QueryForWelcomeToMainMenu(responses chan<- *types.TelegramResponse, output chan<- *types.MessageResponse) {
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
					Text: "GoNext",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 8888,
			Chat: types.Chat{
				Id: 456,
			},
			Photo: []types.Photo{{
				FileId: "",
			}},
		},
	}
}
