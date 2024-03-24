package formatter

import (
	"RegistrationToGames/fmtogram/executer"
	"RegistrationToGames/fmtogram/types"
	"bytes"
	"encoding/json"
	"fmt"
)

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

func (fm *Formatter) WriteEditMesId(mesId int) {
	fm.Message.MessageId = mesId
}

func (fm *Formatter) Complete() {

}

func (fm *Formatter) CheckDelete() (err error) {
	var (
		function    string
		jsonMessage []byte
		finalBuffer *bytes.Buffer
	)
	if fm.DeleteMessage.MessageId != 0 {
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
		jsonMessage   []byte
		jsonKeyboard  []byte
		finalBuffer   *bytes.Buffer
		function      string
		marshalstatus bool
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
		if len(fm.Message.Photo) == 0 && len(fm.Message.Video) == 0 {
			if fm.Message.MessageId == 0 {
				marshalstatus = true
				function = "sendMessage"
			} else {
				function = "editMessageText"
			}
			jsonMessage, err = json.Marshal(fm.Message)
			if err == nil {
				fm.contenttype = "application/json"
				finalBuffer = bytes.NewBuffer(jsonMessage)
			}
		} else if len(fm.Message.Photo) != 0 || len(fm.Message.Video) != 0 {
			marshalstatus = true
			/*
				if len(fm.Message.Photo) > 1 || len(fm.Message.Video) > 1 {
					finalBuffer = bytes.NewBuffer(nil)
					fm.contenttype, err = fm.createMediaGroup(finalBuffer)
				} else {
			*/
			if fm.mediatype[0] == "photo" {
				function = "sendPhoto"
			} else if fm.mediatype[0] == "video" {
				function = "sendVideo"
			}
			if fm.kindofmedia[0] == fromStorage {
				finalBuffer = bytes.NewBuffer(nil)
				fm.contenttype, err = fm.PrepareMedia(finalBuffer)
			} else {
				jsonMessage, err = json.Marshal(fm.Message)
				if err == nil {
					fm.contenttype = "application/json"
					finalBuffer = bytes.NewBuffer(jsonMessage)
				}
			}
			//}
		}

	}
	if err == nil {
		fmt.Println(function, fm.Message.Photo)
		mes = executer.Send(finalBuffer, function, fm.contenttype, marshalstatus)
	}

	return mes, err
}
