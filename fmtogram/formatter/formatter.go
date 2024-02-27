package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (fm *Formatter) WriteDeleteMesId(mesId int) {
	fm.DeleteMessage.MessageId = mesId
}

func (fm *Formatter) Complete() {

}

func (fm *Formatter) CheckDelete() (err error) {
	var (
		function    string
		jsonMessage []byte
		finalBuffer *bytes.Buffer
	)
	fmt.Println(fm.DeleteMessage.MessageId, "HELLO BEFORE")
	if fm.DeleteMessage.MessageId != 0 {
		fmt.Println("HELLO!")
		function = "deleteMessage"
		fm.DeleteMessage.ChatId = fm.Message.ChatID
		jsonMessage, err = json.Marshal(fm.DeleteMessage)
		if err == nil {
			fm.contenttype = "application/json"
			finalBuffer = bytes.NewBuffer(jsonMessage)
		}
		if err == nil {
			_ = executer.Send(finalBuffer, function, fm.contenttype, false)
		}
	}
	return err
}

func (fm *Formatter) Send() (mes *types.MessageResponse, err error) {
	var (
		jsonMessage  []byte
		jsonKeyboard []byte
		finalBuffer  *bytes.Buffer
		function     string
	)

	err = fm.CheckDelete()
	if err == nil {
		if fm.Keyboard.Keyboard != nil {
			jsonKeyboard, err = json.Marshal(fm.Keyboard)
			if err == nil {
				fm.Message.ReplyMarkup = string(jsonKeyboard)
			}
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
		mes = executer.Send(finalBuffer, function, fm.contenttype, true)
	}
	fm.Reset()

	return mes, err
}
