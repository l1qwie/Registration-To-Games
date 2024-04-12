package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["NoGames"] = "У вас пока нет ни одной записи об игре! Вам доступна только смена языка! "
	dict["ChangeLang"] = "Пожалуйста, выберите предпочтительный язык"
	dict["en"] = "Английский язык"
	dict["ru"] = "Русский язык"
	dict["tur"] = "Турецкий язык"
	dict["Lanchanged"] = "Язык бота успешно изменен"
	dict["MainMenu"] = "Главное Меню"
	dict["first"] = "Просмотр расписания"
	dict["second"] = "Регистрация на игру"
	dict["third"] = "Наши фото и видео"
	dict["fourth"] = "Настройки | Мои игры"
}

func en(dict map[string]string) {
	dict["NoGames"] = "You don't have any game records yet! You only have access to language change! "
	dict["ChangeLang"] = "Choose your preferred language"
	dict["en"] = "English language"
	dict["ru"] = "Russian language"
	dict["tur"] = "Turkish language"
	dict["Lanchanged"] = "The bot language has been successfully changed"
	dict["MainMenu"] = "Main Menu"
	dict["first"] = "View schedule"
	dict["second"] = "Game registration"
	dict["third"] = "Our photos and videos"
	dict["fourth"] = "Settings | My games"
}

func tur(dict map[string]string) {
	dict["NoGames"] = "Henüz hiç oyun kaydınız yok! Sadece dil değiştirme seçeneğiniz var! "
	dict["ChangeLang"] = "Tercih ettiğiniz dili seçin"
	dict["en"] = "İngilizce dil"
	dict["ru"] = "Rusça dil"
	dict["tur"] = "Türkçe dil"
	dict["Lanchanged"] = "Bot dil başarıyla değiştirildi"
	dict["MainMenu"] = "Ana Menü"
	dict["first"] = "Programı görüntüle"
	dict["second"] = "Oyun kaydı"
	dict["third"] = "Bizim fotoğraflar ve videolar"
	dict["fourth"] = "Ayarlar | Oyunlarım"
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
