package schedule

import "registrationtogames/fmtogram/types"

func QueryForSeeSchedule(responses chan<- *types.TelegramResponse) {
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
}
