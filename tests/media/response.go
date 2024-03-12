package media

import "RegistrationToGames/fmtogram/types"

func JustTrash(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "O, I want to see some photos or videos from some games!",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 66666,
			Chat: types.Chat{
				Id: 499,
			},
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
						UserID:   499,
						IsBot:    false,
						Name:     "Bogdan",
						LastName: "Dmitriev",
						Username: "Bogdy",
						Language: "ru",
					},
					Text: "Please give me photos!",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 66666,
			Chat: types.Chat{
				Id: 499,
			},
		},
	}
}

func QueryForChooseDirection(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "OK!",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 66666,
			Chat: types.Chat{
				Id: 499,
			},
		},
	}
}

func QueryForChooseMediaGameUnload(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "unload",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 66666,
			Chat: types.Chat{
				Id: 499,
			},
		},
	}
}

func QueryForChooseMediaGameUpload(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "upload",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 66666,
			Chat: types.Chat{
				Id: 499,
			},
		},
	}
}

func QueryForWaitingYourMedia(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "10",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 66666,
			Chat: types.Chat{
				Id: 499,
			},
		},
	}
}

func QueryForUnload(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "10",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 66666,
			Chat: types.Chat{
				Id: 499,
			},
		},
	}
}

func QueryForUpload(responses chan *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
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
					Text: "Take my photo",
				},
				Photo: []types.Photo{{
					FileId: "!@#UIO!@#IOJJKLASEDKLKL#IO!JKLASJKL13419",
				}},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 66666,
			Chat: types.Chat{
				Id: 499,
			},
		},
	}
}
