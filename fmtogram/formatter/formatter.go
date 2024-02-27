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
		slicebyte     []byte
		finalBuffer   *bytes.Buffer
		interBuf      *bytes.Buffer
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
		if fm.Message.Photo == "" && fm.Message.Video == "" {
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
		} else if fm.Message.Video != "" || fm.Message.Photo != "" {
			if fm.Message.MessageId == 0 {
				marshalstatus = true
				if fm.mediatype == "photo" {
					function = "sendPhoto"
				} else if fm.mediatype == "video" {
					function = "sendVideo"
				}
			} else {
				function = "editMessageMedia"
			}
			if fm.kindofmedia == fromStorage {
				finalBuffer = bytes.NewBuffer(nil)
				if fm.Message.MessageId == 0 {
					fm.contenttype, err = fm.PrepareMedia(finalBuffer)
				} else {
					interBuf = bytes.NewBuffer(nil)
					fm.contenttype, err = fm.PrepareMediaForEdit(interBuf)
					if err != nil {
						panic(err)
					}
					slicebyte = interBuf.Bytes()
					jsonMessage, err = json.Marshal(fm.Message)
					if err != nil {
						panic(err)
					}
					finalBuffer.Write(jsonMessage)
					finalBuffer.Write(slicebyte)
				}
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
		mes = executer.Send(finalBuffer, function, fm.contenttype, marshalstatus)
	}
	fm.Reset()

	return mes, err
}
