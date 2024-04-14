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
	dict["Changelang"] = "Изменить язык"
	dict["ChangRec"] = "Изменить бронь на игру"
	dict["YouHaveGames"] = "Вот игры на которые у вас есть запись:\n\n"
	dict["WhatUCanDo"] = "Вы можете изменить информацию по вашей записи или же изменить язык"
	dict["UserSch"] = "<b>%d.</b> <em>Спорт:</em> <b>%s</b>, <em>Дата:</em> <b>%s</b>, <em>Время:</em> <b>%s</b>, <em>Включая вас с вами будет:</em> <b>%d</b>, <em>Цена за одно место:</em> <b>%d %s</b>, <em>Способ оплаты:</em> <b>%s</b>, <em>Статус оплаты:</em> <b>%s</b>\n\n\n"
	dict["volleyball"] = "Волейбол"
	dict["football"] = "Футбол"
	dict["ChooseGame"] = "Выберите вашу игру\n\n"
	dict["ChangeOrDel"] = "Вы хотите изменить или удалить?"
	dict["Change"] = "Изменить игру"
	dict["DelGame"] = "Удалить игру"
	dict["GameDeleted"] = "Ваша бронь на игру успешно удалена\n\n"
	dict["Paid"] = "Оплачено"
	dict["WaitForPaid"] = "Не оплачено"
	dict["WhatUWhantToCh"] = "Что вы хотите изменить?"
	dict["Payment"] = "Способ оплаты"
	dict["Seats"] = "Количество человек со мной"
	dict["GameDeleted"] = "Ваша бронь на игру успешно удалена\n\n"
	dict["ChooseSeat"] = "Выберите или напишите мне количество мест на игру, которые вы хоите занять. На эту игру есть свободных мест %d. У вас уже есть бронь на %d. После ввода ваша бронь на %d мест будет снята. Всего мест, если не учитывать вашу бронь %d"
	dict["ThxForChange"] = "Все успешно изменено\n\n"
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
	dict["Changelang"] = "Change language"
	dict["ChangRec"] = "Change reservation for the game"
	dict["YouHaveGames"] = "Here are the games you have recordings of\n\n"
	dict["WhatUCanDo"] = "You can modify the information about your recording or change the language"
	dict["UserSch"] = "<b>%d.</b> <em>Sport:</em> <b>%s</b>, <em>Date:</em> <b>%s</b>, <em>Time:</em> <b>%s</b>, <em>Including you, there will be:</em> <b>%d</b>, <em>Price per spot:</em> <b>%d %s</b>, <em>Payment method:</em> <b>%s</b>, <em>Payment status:</em> <b>%s</b>\n\n\n"
	dict["volleyball"] = "Volleyball"
	dict["football"] = "Football"
	dict["ChooseGame"] = "Choose your game\n\n"
	dict["ChangeOrDel"] = "Do you want to change or delete?"
	dict["Change"] = "Change the game"
	dict["DelGame"] = "Delete the game"
	dict["GameDeleted"] = "Your game reservation has been successfully deleted\n\n"
	dict["Paid"] = "Paid"
	dict["WaitForPaid"] = "Not paid"
	dict["WhatUWhantToCh"] = "What do you want to change?"
	dict["Payment"] = "Payment method"
	dict["Seats"] = "The number of people with me"
	dict["GameDeleted"] = "Your game reservation has been successfully deleted\n\n"
	dict["ChooseSeat"] = "Choose or write to me the number of seats for the game you want to occupy. There are %d available seats for this game. You already have a reservation for %d. After entering, your reservation for %d seats will be canceled. Total seats, excluding your reservation: %d"
	dict["ThxForChange"] = "Everything has been successfully changed\n\n"
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
	dict["Changelang"] = "Dil değiştir"
	dict["ChangRec"] = "Oyun için rezervasyonu değiştir"
	dict["YouHaveGames"] = "İşte kayıtlarınızın olduğu oyunlar\n\n"
	dict["WhatUCanDo"] = "Kaydınızla ilgili bilgileri değiştirebilir veya dilini değiştirebilirsiniz"
	dict["UserSch"] = "<b>%d.</b> <em>Spor:</em> <b>%s</b>, <em>Tarih:</em> <b>%s</b>, <em>Saat:</em> <b>%s</b>, <em>Sizinle birlikte toplamda olacak kişi sayısı:</em> <b>%d</b>, <em>Her bir yer için fiyat:</em> <b>%d %s</b>, <em>Ödeme yöntemi:</em> <b>%s</b>, <em>Ödeme durumu:</em> <b>%s</b>\n\n\n"
	dict["volleyball"] = "Voleybol"
	dict["football"] = "Futbol"
	dict["ChooseGame"] = "Oyununuzu seçin\n\n"
	dict["WhatUWhantToCh"] = "Ne değiştirmek istiyorsunuz?"
	dict["Payment"] = "Ödeme yöntemi"
	dict["Seats"] = "Benimle birlikteki kişi sayısı"
	dict["GameDeleted"] = "Oyun rezervasyonunuz başarıyla silindi\n\n"
	dict["Paid"] = "Ödendi"
	dict["WaitForPaid"] = "Ödenmedi"
	dict["WhatUWhantToCh"] = "Ne değiştirmek istiyorsunuz?"
	dict["Payment"] = "Ödeme yöntemi"
	dict["Seats"] = "Benimle birlikteki kişi sayısı"
	dict["GameDeleted"] = "Oyun rezervasyonunuz başarıyla silindi\n\n"
	dict["ChooseSeat"] = "Oyunda kapmak istediğiniz koltuk sayısını seçin veya yazın. Bu oyunda %d boş koltuk var. Zaten %d için bir rezervasyonunuz var. Girdikten sonra, %d koltuk için rezervasyonunuz iptal edilecek. Toplam koltuk sayısı, rezervasyonunuz hariç: %d"
	dict["ThxForChange"] = "Hepsi başarıyla değiştirildi\n\n"
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
