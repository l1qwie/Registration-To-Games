package dictionary

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["WelcomeToBot"] = "Добро пожаловать в нашего бота! Этот бот предназначен для регистрации на спортивные игры в Стамбуле, но для начала вам нужно зарегистрироваться у нас!"
	dict["reg"] = "Зарегистрироваться"
}

func en(dict map[string]string) {
	dict["WelcomeToBot"] = "Welcome to our bot! This bot is designed for registration for sports games in Istanbul, but first, you need to register with us!"
	dict["reg"] = "Register"
}

func tur(dict map[string]string) {
	dict["WelcomeToBot"] = "Botumuza hoş geldiniz! Bu bot, İstanbul'daki spor oyunlarına kaydolmak için tasarlanmıştır, ancak önce bizimle kayıt olmanız gerekiyor!"
	dict["reg"] = "Kayıt Ol"
}

func init() {
	Dictionary = make(map[string]map[string]string)
	Dictionary["ru"] = make(map[string]string)
	Dictionary["en"] = make(map[string]string)
	Dictionary["tur"] = make(map[string]string)

	ru(Dictionary["ru"])
	en(Dictionary["en"])
	tur(Dictionary["tur"])
}
