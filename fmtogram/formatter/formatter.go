package formatter

import (
	"RegistrationToGames/fmtogram/errors"
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

func (fm *Formatter) Error(err error) {
	fm.err = err

}

func (fm *Formatter) Complete() error {
	if fm.err != nil {
		errors.MadeMisstake(fm.err)
	}
	return fm.err
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
		if len(fm.Message.InputMedia) == 0 && fm.Message.Photo == "" && fm.Message.Video == "" {
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
		} else if len(fm.Message.InputMedia) != 0 || fm.Message.Photo != "" || fm.Message.Video != "" {
			marshalstatus = true
			if len(fm.Message.InputMedia) > 1 {
				function = "sendMediaGroup"
				//finalBuffer = bytes.NewBuffer(nil)

				jsonMessage, err = json.Marshal(fm.Message)
				if err == nil {
					fm.contenttype = "application/json"
					finalBuffer = bytes.NewBuffer(jsonMessage)
				}

				//fm.contenttype, err = fm.createMediaGroup(finalBuffer)
				//fmt.Println(err)
				//_ = executer.Send(finalBuffer, "sendMediaGroup", fm.contenttype, marshalstatus)
			} else {
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
			}
		}

	}
	if err == nil {
		fmt.Println(finalBuffer.String())
		fmt.Println()
		fmt.Println()
		mes = executer.Send(finalBuffer, function, fm.contenttype, marshalstatus)
	}

	return mes, err
}
