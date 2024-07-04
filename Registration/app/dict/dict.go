package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["volleyball"] = "Волейбол"
	dict["football"] = "Футбол"
	dict["freeSpace"] = "Свободный мест осталось:"
	dict["MainMenu"] = "Главное Меню"
	dict["ChooseAnyGame"] = "Выберите интересующую вас игру"
	dict["NoGames"] = "Игра пока что нет\n\n"
	dict["1"] = "Я буду один"
	dict["2"] = "Нас будет двое"
	dict["3"] = "Нас будет трое"
	dict["ChooseSeats"] = "Выберите или введите желаемое количество мест. Прошу обратить ваше внимание на то, что цена за одно место на эту игру составляет %d %s"
	dict["NoMoreSeats"] = "К сожалению, вы ввели либо слишком большое количество свободных мест для игры, которых не существует"
	dict["payByCard"] = "Онлайн оплата"
	dict["payByCash"] = "Наличкой администратору"
	dict["ChoosePaymethod"] = "Выберите способ оплаты"
	dict["pay"] = "Оплатить"
	dict["GoNext"] = "Дальше"
	dict["RegistrationCompleted"] = `Прекрасно! Теперь вы зарегестрированны на игру\n
	1. Вид спорта: <b>%s</b>\n
	2. Дата: <b>%s</b>\n
	3. Время: <b>%s</b>\n
	4. Вы записали <b>%d</b> персон на эту игру\n
	5. Оплата: <b>by %s</b>\n
	6. Ваша стоимость посещения: <b>%d %s</b>\n\n
	
	***Вы можете изменить некоторые данные Вашей записи\n
	или же удалить ее в Главном Меню нажав на <b>"Настройки | Мои игры"</b>***\n\n
	
	❤️❤️❤️Ждем вас по адресу: %s\n
	%s`
	dict["SeatsAreFull"] = "Упс! Похоже, места, на которые вы рассчитывали, уже заняты! Мне очень жаль. У нас осталось мест: %s.\n"
	dict["Review"] = "Хотите пересмотреть свои планы?"
	dict["go-ahead"] = "Да, хочу"
	dict["no-deleteAll"] = "Нет, не хочу"
	dict["WaitForYourMoney"] = `Вам нужно перевести %d %s. Выше я выслал вам QR-code, по которому вы можете перейти и оплатить или же нажмите на кнопку ниже "Оплатить"`
	dict["first"] = "Просмотр расписания"
	dict["second"] = "Регистрация на игру"
	dict["third"] = "Наши фото и видео"
	dict["fourth"] = "Настройки | Мои игры"
	dict["Oops!"] = "Ой! Кажется, что-то внутри сломалось! Повторите попытку позже. Если проблема не исчезнет, обратитесь к администратору!"
}

func en(dict map[string]string) {
	dict["volleyball"] = "Volleyball"
	dict["football"] = "Football"
	dict["freeSpace"] = "Available seats left:"
	dict["MainMenu"] = "Main Menu"
	dict["ChooseAnyGame"] = "Choose the game you are interested in"
	dict["NoGames"] = "The game is not available yet\n\n"
	dict["1"] = "I will be alone"
	dict["2"] = "There will be two of us"
	dict["3"] = "There will be three of us"
	dict["ChooseSeats"] = "Choose or enter the desired number of seats. Please note that the price for one seat for this game is %d %s"
	dict["NoMoreSeats"] = "Unfortunately, you entered either too many available seats for the game that do not exist"
	dict["payByCard"] = "Pay by card"
	dict["payByCash"] = "Pay by cash"
	dict["ChoosePaymethod"] = "Choose a payment method"
	dict["pay"] = "Pay"
	dict["GoNext"] = "Next"
	dict["RegistrationCompleted"] = `Great! Now you are registered for the game\n
	1. Sport type: <b>%s</b>\n
	2. Date: <b>%s</b>\n
	3. Time: <b>%s</b>\n
	4. You have registered <b>%d</b> persons for this game\n
	5. Payment: <b>%s</b>\n
	6. Your visit cost: <b>%d %s</b>\n\n
	
	You can modify some details of your registration\n
	or delete it in the Main Menu by clicking <b>"Settings | My games"</b>\n\n
	
	❤️❤️❤️We look forward to seeing you at: %s\n
	%s`
	dict["SeatsAreFull"] = "Oops! It seems the seats you were counting on are already taken! I'm sorry. We have seats left: %s.\n"
	dict["Review"] = "Would you like to reconsider your plans?"
	dict["go-ahead"] = "Yes, I want"
	dict["no-deleteAll"] = "No, I don't want"
	dict["WaitForYourMoney"] = `You need to transfer %d %s. Above, I sent you a QR code, which you can use to proceed and make the payment, or click the "Pay" button below`
	dict["first"] = "View schedule"
	dict["second"] = "Game registration"
	dict["third"] = "Our photos and videos"
	dict["fourth"] = "Settings | My games"
	dict["Oops!"] = "Oops! It looks like something broke inside! Please try again later. If the problem persists, contact the administrator!"
}

func tur(dict map[string]string) {
	dict["volleyball"] = "Voleybol"
	dict["football"] = "Futbol"
	dict["freeSpace"] = "Kalan boş koltuklar:"
	dict["MainMenu"] = "Ana Menü"
	dict["ChooseAnyGame"] = "İlgilendiğiniz oyunu seçin"
	dict["NoGames"] = "Oyun henüz mevcut değil\n\n"
	dict["1"] = "Ben yalnız olacağım"
	dict["2"] = "İkimiz olacağız"
	dict["3"] = "Üç kişi olacağız"
	dict["ChooseSeats"] = "İstenilen koltuk sayısını seçin veya girin. Lütfen unutmayın ki bu oyun için bir koltuk fiyatı %d %s'dir"
	dict["NoMoreSeats"] = "Maalesef, var olmayan oyun için fazla sayıda boş koltuk girdiniz"
	dict["payByCard"] = "Online ödeme"
	dict["payByCash"] = "Nakit, yöneticiye"
	dict["ChoosePaymethod"] = "Ödeme yöntemi seçin"
	dict["pay"] = "Ödeme Yap"
	dict["GoNext"] = "İleri"
	dict["RegistrationCompleted"] = `Harika! Şimdi oyuna kayıt oldunuz\n
	1. Spor türü: <b>%s</b>\n
	2. Tarih: <b>%s</b>\n
	3. Saat: <b>%s</b>\n
	4. Bu oyuna <b>%d</b> kişiyi kaydettiniz\n
	5. Ödeme: <b>%s</b>\n
	6. Ziyaret maliyetiniz: <b>%d %s</b>\n\n
	
	Kaydınızın bazı detaylarını değiştirebilirsiniz\n
	veya <b>"Ayarlar | Oyunlarım"</b> seçeneğine tıklayarak silebilirsiniz\n\n
	
	❤️❤️❤️Sizi bekliyoruz: %s\n
	%s`
	dict["SeatsAreFull"] = "Hoop! Hesapladığınız koltuklar maalesef zaten dolu! Üzgünüm. Kalan yerlerimiz: %s.\n"
	dict["Review"] = "Planlarınızı tekrar düşünmek ister misiniz?"
	dict["go-ahead"] = "Evet, istiyorum"
	dict["no-deleteAll"] = "Hayır, istemiyorum"
	dict["WaitForYourMoney"] = `%d %s göndermeniz gerekiyor. Yukarıda size bir QR kodu gönderdim, bunu kullanarak ödeme yapabilir veya aşağıdaki "Ödeme Yap" düğmesine tıklayabilirsiniz`
	dict["first"] = "Programı görüntüle"
	dict["second"] = "Oyun kaydı"
	dict["third"] = "Bizim fotoğraflar ve videolar"
	dict["fourth"] = "Ayarlar | Oyunlarım"
	dict["Oops!"] = "Aman! İçeride bir şey kırıldı gibi görünüyor! Lütfen daha sonra tekrar deneyin. Sorun devam ederse, yöneticiye başvurun!"
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
