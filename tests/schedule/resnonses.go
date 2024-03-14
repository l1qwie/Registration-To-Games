package schedule

import (
	"RegistrationToGames/fmtogram/types"
)

func QueryForSeeSchedule(responses chan<- *types.TelegramResponse, output chan<- *types.MessageResponse) {
	responses <- &types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   488,
						IsBot:    false,
						Name:     "John",
						LastName: "Doe",
						Username: "johndoe",
						Language: "ru",
					},
					Text: "Show me the schedule",
				},
			},
		},
	}
	output <- &types.MessageResponse{
		Ok: true,
		Result: types.Message{
			MessageId: 9999,
			Chat: types.Chat{
				Id: 488,
			},
			Photo: []types.Photo{},
		},
	}
}
