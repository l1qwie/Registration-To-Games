package helper

import (
	"RegistrationToGames/fmtogram/types"
	"fmt"
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
	if message.Ok {
		messageId = message.Result.MessageId
	} else {
		err = fmt.Errorf("we don't have a message id")
	}
	return messageId, err
}

func ReturnPhotosFileIdout(message *types.MessageResponse) ([]string, error) {
	var (
		err     error
		fileIds []string
	)
	fileIds = make([]string, len(message.Result.Photo))
	if len(message.Result.Photo) > 0 {
		for i := 0; i < len(message.Result.Photo); i++ {
			fileIds[i] = message.Result.Photo[i].FileId
		}
	} else {
		err = fmt.Errorf("we don't have a Photo fileId")
	}
	return fileIds, err
}

func ReturnVideosFileIdout(message *types.MessageResponse) ([]string, error) {
	var (
		err     error
		fileIds []string
	)
	fileIds = make([]string, len(message.Result.Video))
	if len(message.Result.Video) > 0 {
		for i := 0; i < len(message.Result.Video); i++ {
			fileIds[i] = message.Result.Video[i].FileId
		}
	} else {
		err = fmt.Errorf("we don't have a Video fileId")
	}
	return fileIds, err
}

func ReturnPhotosFileIdfrom(tr *types.TelegramResponse) ([]string, error) {
	var (
		err     error
		fileIds []string
	)
	j := 0
	if len(tr.Result[0].Message.Photo) > 0 {
		l := len(tr.Result[0].Message.Photo) / 4
		fileIds = make([]string, l)
		for i := 0; i < l; i++ {
			fileIds[i] = tr.Result[0].Message.Photo[j].FileId
			j = j + 4
		}
	} else {
		err = fmt.Errorf("we don't have a Photo fileId")
	}
	return fileIds, err
}

func ReturnVideosFileIdfrom(tr *types.TelegramResponse) ([]string, error) {
	var (
		err     error
		fileIds []string
	)
	fileIds = make([]string, len(tr.Result[0].Message.Video))
	if len(tr.Result[0].Message.Video) > 0 {
		for i := 0; i < len(tr.Result[0].Message.Video); i++ {
			fileIds[i] = tr.Result[0].Message.Video[i].FileId
		}
	} else {
		err = fmt.Errorf("we don't have a Video fileId")
	}
	return fileIds, err
}
