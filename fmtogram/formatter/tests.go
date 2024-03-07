package formatter

import (
	"RegistrationToGames/fmtogram/errors"
	"fmt"
)

func (tfm *Formatter) AssertPhoto(path string, condition bool) (err error) {
	var function string
	if tfm.Message.Photo != path {
		if tfm.kindofmedia == fromStorage {
			function = "AddPhotoFromStorage"
		} else if tfm.kindofmedia == fromInternet {
			function = "AddPhotoFromInternet"
		} else if tfm.kindofmedia == fromTelegram {
			function = "AddPhotoFromTG"
		}
		err = errors.AssertTest(tfm.Message.Photo, function, path, "AssertPhoto")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (tfm *Formatter) AssertVideo(path string, condition bool) (err error) {
	var function string
	if tfm.Message.Video != path {
		if tfm.kindofmedia == fromStorage {
			function = "AddVideoFromStorage"
		} else if tfm.kindofmedia == fromInternet {
			function = "AddVideoFromInternet"
		} else if tfm.kindofmedia == fromTelegram {
			function = "AddVideoFromTG"
		}
		err = errors.AssertTest(tfm.Message.Video, function, path, "AssertVideo")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (tfm *Formatter) AssertInlineKeyboard(testdim []int, kbNames, kbDatas, typeofbuttons []string, condition bool) (err error) {
	var (
		dim     []int
		counter int
	)

	for i := 0; i < len(tfm.Keyboard.Keyboard); i++ {
		dim = append(dim, len(tfm.Keyboard.Keyboard[i]))
	}
	if len(testdim) == len(dim) {
		for i := 0; i < len(dim); i++ {
			if testdim[i] != dim[i] {
				err = errors.AssertTest(fmt.Sprint(dim), "SetIkbdDim", fmt.Sprint(testdim), "AssertInlineKeyboard")
				if condition {
					panic(err)
				}
			}
		}
		if len(kbNames) == len(kbDatas) && len(kbNames) == len(typeofbuttons) && err == nil {
			for i := 0; i < len(testdim); i++ {
				for j := 0; j < testdim[i]; j++ {
					if tfm.Keyboard.Keyboard[i][j].Label != kbNames[counter] {
						err = errors.AssertTest(fmt.Sprint("name of buttons is ", tfm.Keyboard.Keyboard[i][j].Label), "WriteInlineButtonUrl/WriteInlineButtonCmd", fmt.Sprint("name of buttons is ", kbNames[counter]), "AssertInlineKeyboard")
						if condition {
							panic(err)
						}
					} else if typeofbuttons[i] == "url" && tfm.Keyboard.Keyboard[i][j].Url != kbDatas[counter] {
						err = errors.AssertTest(fmt.Sprint("url of button is ", tfm.Keyboard.Keyboard[i][j].Url), "WriteInlineButtonUrl", fmt.Sprint("url of button is ", kbDatas[counter]), "AssertInlineKeyboard")
						if condition {
							panic(err)
						}
					} else if typeofbuttons[i] == "cmd" && tfm.Keyboard.Keyboard[i][j].Cmd != kbDatas[counter] {
						err = errors.AssertTest(fmt.Sprint("cmd of button is ", tfm.Keyboard.Keyboard[i][j].Cmd), "WriteInlineButtonCmd", fmt.Sprint("cmd of button is ", kbDatas[counter]), "AssertInlineKeyboard")
						if condition {
							panic(err)
						}
					}
					counter++
				}
			}
		} else if err == nil {
			err = errors.JustError()
			if condition {
				panic(err)
			}
		}
	} else {
		err = errors.AssertTest(fmt.Sprint("length of slice is ", len(dim)), "SetIkbdDim", fmt.Sprint("length of slice is ", len(testdim)), "AssertInlineKeyboard")
		if condition {
			panic(err)
		}
	}

	return err
}

func (tfm *Formatter) AssertString(lineoftext string, condition bool) (err error) {
	if tfm.Message.Text != lineoftext {
		err = errors.AssertTest(fmt.Sprint(tfm.Message.Text), "WriteString", fmt.Sprint(lineoftext), "AssertString")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}

	return err
}

func (tfm *Formatter) AssertChatId(chatID int, condition bool) (err error) {
	if tfm.Message.ChatID != chatID {
		err = errors.AssertTest(fmt.Sprint(tfm.Message.ChatID), "WriteChatId", fmt.Sprint(chatID), "AssertChatId")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (tfm *Formatter) AssertParseMode(parseMode string, condition bool) (err error) {
	if tfm.Message.ParseMode != parseMode {
		err = errors.AssertTest(fmt.Sprintf(tfm.Message.ParseMode), "WriteParseMode", fmt.Sprint(parseMode), "AssertParseMode")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (tfm *Formatter) AssertEditMessageId(messageId int, condition bool) (err error) {
	if tfm.Message.MessageId != messageId {
		err = errors.AssertTest(fmt.Sprint(tfm.Message.MessageId), "WriteEditMesId", fmt.Sprint(messageId), "AssertEditMessageId")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (tfm *Formatter) AssertDeleteMessageId(messageId int, condition bool) (err error) {
	if tfm.DeleteMessage.MessageId != messageId {
		err = errors.AssertTest(fmt.Sprint(tfm.DeleteMessage.MessageId), "WriteDeleteMesId", fmt.Sprint(messageId), "AssertDeleteMessageId")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}
