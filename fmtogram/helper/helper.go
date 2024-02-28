package helper

import (
	"fmt"
	"registrationtogames/fmtogram/types"
)

func ReturnText(telegramResponse *types.TelegramResponse) (text string) {
	if telegramResponse.Result[0].Message.Text != "" {
		text = telegramResponse.Result[0].Message.Text
	} else if telegramResponse.Result[0].Query.Data != "" {
		text = telegramResponse.Result[0].Query.Data
	}
	return text
}

func ReturnChatId(telegramResponse *types.TelegramResponse) (chatID int) {
	if telegramResponse.Result[0].Message.TypeFrom.UserID != 0 {
		chatID = telegramResponse.Result[0].Message.TypeFrom.UserID
	} else if telegramResponse.Result[0].Query.TypeFrom.UserID != 0 {
		chatID = telegramResponse.Result[0].Query.TypeFrom.UserID
	}
	return chatID
}

func ReturnName(telegramResponse *types.TelegramResponse) (name string) {
	if telegramResponse.Result[0].Message.TypeFrom.Name != "" {
		name = telegramResponse.Result[0].Message.TypeFrom.Name
	} else if telegramResponse.Result[0].Query.TypeFrom.Name != "" {
		name = telegramResponse.Result[0].Query.TypeFrom.Name
	}
	return name
}

func ReturnLastName(telegramResponse *types.TelegramResponse) (lastname string) {
	if telegramResponse.Result[0].Message.TypeFrom.LastName != "" {
		lastname = telegramResponse.Result[0].Message.TypeFrom.LastName
	} else if telegramResponse.Result[0].Query.TypeFrom.LastName != "" {
		lastname = telegramResponse.Result[0].Query.TypeFrom.LastName
	}
	return lastname
}

func ReturnUsername(telegramResponse *types.TelegramResponse) (username string) {
	if telegramResponse.Result[0].Message.TypeFrom.Username != "" {
		username = telegramResponse.Result[0].Message.TypeFrom.Username
	} else if telegramResponse.Result[0].Query.TypeFrom.Username != "" {
		username = telegramResponse.Result[0].Query.TypeFrom.Username
	}
	return username
}

func ReturnLanguage(telegramResponse *types.TelegramResponse) (language string) {
	if telegramResponse.Result[0].Message.TypeFrom.Language != "" {
		language = telegramResponse.Result[0].Message.TypeFrom.Language
	} else if telegramResponse.Result[0].Query.TypeFrom.Language != "" {
		language = telegramResponse.Result[0].Query.TypeFrom.Language
	}
	return language
}

func ReturnBotStatus(telegramResponse *types.TelegramResponse) (botstatus bool) {

	m_isbot := telegramResponse.Result[0].Message.TypeFrom.IsBot
	cl_isbot := telegramResponse.Result[0].Query.TypeFrom.IsBot

	if !m_isbot && !cl_isbot {
		botstatus = false
	} else if m_isbot && !cl_isbot || !m_isbot && cl_isbot {
		botstatus = true
	}
	return botstatus
}

func ReturnMessageId(message *types.MessageResponse) (messageId int, err error) {
	/*if len(message.Result) != 0 {
		messageId = message.Result[0].Chat.Id
	} else {
		err = fmt.Errorf("Don't have message Id")
	}*/
	if message.Ok {
		messageId = message.Result.MessageId
	} else {
		err = fmt.Errorf("we don't have a message id")
	}
	return messageId, err
}

func ReturnPhotoFileId(message *types.MessageResponse) (fileId string, err error) {
	if len(message.Result.Photo) > 0 {
		fileId = message.Result.Photo[0].FileId
	} else {
		err = fmt.Errorf("we don't have a Photo fileId")
	}
	return fileId, err
}

func ReturnVideo(message *types.MessageResponse) (fileId string, err error) {
	if len(message.Result.Video) > 0 {
		fileId = message.Result.Video[0].FileId
	} else {
		err = fmt.Errorf("we don't have a Video fileId")
	}
	return fileId, err
}
