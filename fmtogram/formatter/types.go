package formatter

import "registrationtogames/fmtogram/types"

type Formatter struct {
	Message       types.SendMessagePayload
	Keyboard      InlineKeyboard
	DeleteMessage types.DelMessage
	contenttype   string
	kindofmedia   int
	mediatype     string
}

type InlineKeyboard struct {
	Keyboard [][]btn `json:"inline_keyboard"`
	x        int
	y        int
}

type btnKind int

const (
	bCmd btnKind = 1
	bUrl btnKind = 2

	fromStorage  int = 0
	fromTelegram int = 1
	fromInternet int = 2
)

type btn struct {
	Label string `json:"text"`
	what  btnKind
	Cmd   string `json:"callback_data"`
	Url   string `json:"url"`
}
