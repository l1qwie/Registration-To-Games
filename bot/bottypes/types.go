package bottypes

type RegToGames struct {
	GameId  int
	Seats   int
	Sport   string
	Payment string
}

type Media struct {
	Id        string
	Interval  string
	Direcrion string
	Limit     int
	DelGameId int
	Counter   int
}

type UserRec struct {
	Changeable     string
	GameId         int
	ActGame        string
	WillChangeable string
	NewPay         string
}
