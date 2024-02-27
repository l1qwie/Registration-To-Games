package types

import "fmt"

const HttpsRequest = "https://api.telegram.org/"

const (
	Markdown string = "Markdown"
	HTML     string = "HTML"
)

type InfMessage struct {
	TypeFrom User   `json:"from"`
	Text     string `json:"text"`
}

type User struct {
	UserID   int    `json:"id"`
	IsBot    bool   `json:"is_bot"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Language string `json:"language_code"`
}

type TelegramUpdate struct {
	UpdateID  int        `json:"update_id"`
	MessageId int        `json:"message_id"`
	Message   InfMessage `json:"message"`
	Query     Callback   `json:"callback_query"`
}

type Callback struct {
	TypeFrom User   `json:"from"`
	Data     string `json:"data"`
}

type TelegramError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TelegramResponse struct {
	Ok     bool             `json:"ok"`
	Result []TelegramUpdate `json:"result"`
	Error  *TelegramError   `json:"error,omitempty"`
}

type JustForUpdate struct {
	Ok     bool            `json:"ok"`
	Result []StorageOfJson `json:"result"`
}

type StorageOfJson struct {
	ID int `json:"update_id"`
}

type SendMessagePayload struct {
	ChatID      int    `json:"chat_id"`
	Text        string `json:"text"`
	ReplyMarkup string `json:"reply_markup"`
	Photo       string `json:"photo"`
	Video       string `json:"video"`
	ParseMode   string `json:"parse_mode"`
}

type DelMessage struct {
	ChatId    int `json:"chat_id"`
	MessageId int `json:"message_id"`
}

type Chat struct {
	Id int `json:"id"`
}

type Message struct {
	MessageId int  `json:"message_id"`
	Chat      Chat `json:"chat"`
}

type MessageResponse struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type FMTRS interface {
	WriteString(string)
	WriteChatId(int)
	WriteDeleteMesId(int)
	WriteEditMesId(int)
	AddPhotoFromStorage(string)
	AddPhotoFromTG(string)
	AddPhotoFromInternet(string)
	AddVideoFromStorage(string)
	AddVideoFromTG(string)
	AddVideoFromInternet(string)
	SetIkbdDim([]int)
	WriteInlineButtonCmd(string, string)
	WriteInlineButtonUrl(string, string)
	Send() error
}

type Responser interface {
	RequestOffset(string, *int) error
	Updates(string, *int, *TelegramResponse) error
}

func ConnectTo() (body string) {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", username, password, dbname, sslmode)
}
