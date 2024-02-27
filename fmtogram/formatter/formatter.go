package formatter

import (
	"bytes"
	"encoding/json"
	"registrationtogames/fmtogram/executer"
	"registrationtogames/fmtogram/types"
)

func (fm *Formatter) Reset() {
	*fm = Formatter{
		Message: types.SendMessagePayload{
			ChatID:      0,
			Text:        "",
			ReplyMarkup: "",
			Photo:       "",
			Video:       "",
		},
		Keyboard: InlineKeyboard{
			Keyboard: nil,
			x:        0,
			y:        0,
		},
		contenttype: "",
		kindofmedia: 0,
		mediatype:   "",
	}
}

func (fm *Formatter) WriteString(lineoftext string) {
	fm.Message.Text = lineoftext
}

func (fm *Formatter) WriteChatId(chatID int) {
	fm.Message.ChatID = chatID
}

func (fm *Formatter) WriteParseMode(mode string) {
	fm.Message.ParseMode = mode
}

func (fm *Formatter) Complete() {

}

func (fm *Formatter) Send() (mes *types.MessageResponse, err error) {
	var (
		jsonMessage  []byte
		jsonKeyboard []byte
		finalBuffer  *bytes.Buffer
		function     string
	)

	if fm.Keyboard.Keyboard != nil {
		jsonKeyboard, err = json.Marshal(fm.Keyboard)
		if err == nil {
			fm.Message.ReplyMarkup = string(jsonKeyboard)
		}
	}

	if err == nil {
		if fm.Message.Photo == "" && fm.Message.Video == "" {
			function = "sendMessage"
			jsonMessage, err = json.Marshal(fm.Message)
			if err == nil {
				fm.contenttype = "application/json"
				finalBuffer = bytes.NewBuffer(jsonMessage)
			}
		} else if fm.Message.Video != "" || fm.Message.Photo != "" {
			if fm.mediatype == "photo" {
				function = "sendPhoto"
			} else if fm.mediatype == "video" {
				function = "sendVideo"
			}

			if fm.kindofmedia == fromStorage {
				finalBuffer = bytes.NewBuffer(nil)
				fm.contenttype, err = fm.PrepareMedia(finalBuffer)
			} else {
				jsonMessage, err = json.Marshal(fm.Message)
				if err == nil {
					fm.contenttype = "application/json"
					finalBuffer = bytes.NewBuffer(jsonMessage)
				}
			}
		}
	}
	if err == nil {
		mes = executer.Send(finalBuffer, function, fm.contenttype)
	}
	fm.Reset()

	return mes, err
}
