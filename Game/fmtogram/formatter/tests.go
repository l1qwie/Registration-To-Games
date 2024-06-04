package formatter

import (
	"Game/fmtogram/errors"
	"fmt"
)

func (fm *Formatter) AssertPhoto(path string, condition bool) (err error) {
	var function string
	if len(fm.Message.Photo) > 0 {
		if fm.Message.Photo != path {
			if fm.kindofmedia[0] == fromStorage {
				function = "AddPhotoFromStorage"
			} else if fm.kindofmedia[0] == fromInternet {
				function = "AddPhotoFromInternet"
			} else if fm.kindofmedia[0] == fromTelegram {
				function = "AddPhotoFromTG"
			}
			err = errors.AssertTest(fm.Message.Photo, function, path, "AssertPhoto")
		}
	} else {
		err = errors.AssertTest(fmt.Sprint(fm.Message.Photo), function, path, "AssertPhoto")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (fm *Formatter) AssertVideo(path string, condition bool) (err error) {
	var function string
	if len(fm.Message.Video) > 0 {
		if fm.Message.Video != path {
			if fm.kindofmedia[0] == fromStorage {
				function = "AddVideoFromStorage"
			} else if fm.kindofmedia[0] == fromInternet {
				function = "AddVideoFromInternet"
			} else if fm.kindofmedia[0] == fromTelegram {
				function = "AddVideoFromTG"
			}
			err = errors.AssertTest(fm.Message.Video, function, path, "AssertVideo")
		}
	} else {
		err = errors.AssertTest(fmt.Sprint(fm.Message.Video), function, path, "AssertPhoto")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (fm *Formatter) AssertInlineKeyboard(testdim []int, kbNames, kbDatas, typeofbuttons []string, condition bool) (err error) {
	var (
		dim     []int
		counter int
	)

	for i := 0; i < len(fm.Keyboard.Keyboard); i++ {
		dim = append(dim, len(fm.Keyboard.Keyboard[i]))
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
					if fm.Keyboard.Keyboard[i][j].Label != kbNames[counter] {
						err = errors.AssertTest(fmt.Sprint("name of buttons is ", fm.Keyboard.Keyboard[i][j].Label), "WriteInlineButtonUrl/WriteInlineButtonCmd", fmt.Sprint("name of buttons is ", kbNames[counter]), "AssertInlineKeyboard")
						if condition {
							panic(err)
						}
					} else if typeofbuttons[i] == "url" && fm.Keyboard.Keyboard[i][j].Url != kbDatas[counter] {
						err = errors.AssertTest(fmt.Sprint("url of button is ", fm.Keyboard.Keyboard[i][j].Url), "WriteInlineButtonUrl", fmt.Sprint("url of button is ", kbDatas[counter]), "AssertInlineKeyboard")
						if condition {
							panic(err)
						}
					} else if typeofbuttons[i] == "cmd" && fm.Keyboard.Keyboard[i][j].Cmd != kbDatas[counter] {
						err = errors.AssertTest(fmt.Sprint("cmd of button is ", fm.Keyboard.Keyboard[i][j].Cmd), "WriteInlineButtonCmd", fmt.Sprint("cmd of button is ", kbDatas[counter]), "AssertInlineKeyboard")
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

func (fm *Formatter) AssertString(lineoftext string, condition bool) (err error) {
	if fm.Message.Text != lineoftext {
		err = errors.AssertTest(fmt.Sprint(fm.Message.Text), "WriteString", fmt.Sprint(lineoftext), "AssertString")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}

	return err
}

func (fm *Formatter) AssertChatId(chatID int, condition bool) (err error) {
	if fm.Message.ChatID != chatID {
		err = errors.AssertTest(fmt.Sprint(fm.Message.ChatID), "WriteChatId", fmt.Sprint(chatID), "AssertChatId")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (fm *Formatter) AssertParseMode(parseMode string, condition bool) (err error) {
	if fm.Message.ParseMode != parseMode {
		err = errors.AssertTest(fmt.Sprintf(fm.Message.ParseMode), "WriteParseMode", fmt.Sprint(parseMode), "AssertParseMode")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (fm *Formatter) AssertEditMessageId(messageId int, condition bool) (err error) {
	if fm.Message.MessageId != messageId {
		err = errors.AssertTest(fmt.Sprint(fm.Message.MessageId), "WriteEditMesId", fmt.Sprint(messageId), "AssertEditMessageId")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (fm *Formatter) AssertDeleteMessageId(messageId int, condition bool) (err error) {
	if fm.DeleteMessage.MessageId != messageId {
		err = errors.AssertTest(fmt.Sprint(fm.DeleteMessage.MessageId), "WriteDeleteMesId", fmt.Sprint(messageId), "AssertDeleteMessageId")
	}
	if condition {
		if err != nil {
			panic(err)
		}
	}
	return err
}

func (fm *Formatter) AssertMapOfMedia(group map[string]string, con bool) (err error) {
	if len(fm.Message.InputMedia) != len(group) {
		err = errors.AssertTest(fmt.Sprint("length of map is ", (len(fm.Message.Photo)+len(fm.Message.Video))), "AddMapOfMedia", fmt.Sprint("length of map is ", len(group)), "AssertMapOfMedia")
	} else {
		for key := range group {
			found := false
			_, ok := group[key]
			if ok {
				for i := 0; i < len(fm.Message.InputMedia) && !found; i++ {
					if fm.Message.InputMedia[i].Media == key {
						found = true
					}
				}
			}
			if !found {
				err = errors.AssertTest(fmt.Sprint(fm.Message.InputMedia), "AddMapOfMedia", fmt.Sprint(group), "AssertMapOfMedia")
			}
		}
	}
	if con {
		if err != nil {
			panic(err)
		}
	}
	return err
}
