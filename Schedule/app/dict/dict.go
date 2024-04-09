package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["Schedule"] = "<b>%d.</b> <em>Спорт:</em> <b>%s</b>, <em>Дата:</em> <b>%s</b>, <em>Время:</em> <b>%s</b>, <em>Осталось свободных мест:</em> <b>%d</b>, <em>Цена за одно место:</em> <b>%d %s</b>\n\n\n"
	dict["NoGames"] = "Игра пока что нет\n\n"
	dict["first"] = "Просмотр расписания"
	dict["second"] = "Регистрация на игру"
	dict["third"] = "Наши фото и видео"
	dict["fourth"] = "Настройки | Мои игры"
	dict["volleyball"] = "Волейбол"
	dict["football"] = "Футбол"
	dict["MainMenu"] = "Главное Меню"
}

func en(dict map[string]string) {
	dict["Schedule"] = "<b>%d.</b> <em>Sport:</em> <b>%s</b>, <em>Date:</em> <b>%s</b>, <em>Time:</em> <b>%s</b>, <em>Remaining seats:</em> <b>%d</b>, <em>Price per seat:</em> <b>%d %s</b>\n\n\n"
	dict["NoGames"] = "The game is not available yet\n\n"
	dict["first"] = "View schedule"
	dict["second"] = "Game registration"
	dict["third"] = "Our photos and videos"
	dict["fourth"] = "Settings | My games"
	dict["volleyball"] = "Volleyball"
	dict["football"] = "Football"
	dict["MainMenu"] = "Main Menu"
}

func tur(dict map[string]string) {
	dict["Schedule"] = "<b>%d.</b> <em>Spor:</em> <b>%s</b>, <em>Tarih:</em> <b>%s</b>, <em>Saat:</em> <b>%s</b>, <em>Kalan boş yerler:</em> <b>%d</b>, <em>Bir koltuk için fiyat:</em> <b>%d %s</b>\n\n\n"
	dict["NoGames"] = "Oyun henüz mevcut değil\n\n"
	dict["first"] = "Programı görüntüle"
	dict["second"] = "Oyun kaydı"
	dict["third"] = "Bizim fotoğraflar ve videolar"
	dict["fourth"] = "Ayarlar | Oyunlarım"
	dict["volleyball"] = "Voleybol"
	dict["football"] = "Futbol"
	dict["MainMenu"] = "Ana Menü"
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
