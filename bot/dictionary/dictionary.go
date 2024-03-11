package dictionary

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["volleyball"] = "Волейбол"
	dict["football"] = "Футбол"
	dict["WelcomeToBot"] = "Добро пожаловать в нашего бота! Этот бот предназначен для регистрации на спортивные игры в Стамбуле, но для начала вам нужно зарегистрироваться у нас!"
	dict["reg"] = "Зарегистрироваться"
	dict["BotRules"] = `Добро Пожаловать в нашу дружную компанию! Далее я вам напишу некоторые правила, и рекомендации как пользоваться ботом:\n1. Просмотр расписания.\n     
	- Предельно просто. Вы сможете посмтореть полное расписание всех возможных игр.\n2. Регистрация на игру.\n      
	- Бот поможет вам записаться на интересующую вас игру.\n3. Посмотреть фотографии.\n      
	- Тут будут храниться все фотографии, видео, гифки и все остальное, что будут выкладывать другие пользователи с уже прошедших игр.\n4. Посмототреть мои записи.\n      
	- Нажав сюда, вы сможете посмотреть игры, которые вы ожидаете. Так же тут вы сможете отменить запись.\n\n\n
	P.S. Если что, то у бота есть волшебная команда /menu. Эта команда всегда сможет вас вернуть в главное меню. 
	Естесвенно, если вы пришлете ее, когда у вас есть какой то незаконченный процес в боте, то прогресс не сохраниться`
	dict["allright"] = "Все понятно!"
	dict["WelcomeToMainMenu"] = "Добро пожаловать в гланое меню"
	dict["first"] = "Просмотр расписания"
	dict["second"] = "Регистрация на игру"
	dict["third"] = "Наши фото и видео"
	dict["fourth"] = "Настройки | Мои игры"
	dict["ChooseAnyGame"] = "Выберите интересующую вас игру"
	dict["freeSpace"] = "Свободный мест осталось:"
	dict["next"] = ">>"
	dict["previous"] = "<<"
	dict["ChooseSeats"] = "Выберите или введите желаемое количество мест. Прошу обратить ваше внимание на то, что цена за одно место на эту игру составляет %d %s"
	dict["NoMoreSeats"] = "К сожалению, вы ввели либо слишком большое количество свободных мест для игры, которых не существует"
	dict["1"] = "Я буду один"
	dict["2"] = "Нас будет двое"
	dict["3"] = "Нас будет трое"
	dict["ChoosePaymethod"] = "Выберите способ оплаты"
	dict["payByCard"] = "Онлайн оплата"
	dict["payByCash"] = "Наличкой администратору"
	dict["MainMenu"] = "Главное Меню"
	dict["WaitForYourMoney"] = `Вам нужно перевести %d %s. Выше я выслал вам QR-code, по которому вы можете перейти и оплатить или же нажмите на кнопку ниже "Оплатить"`
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
	https://www.google.com/maps?q=%f,%f`
	dict["SeatsAreFull"] = "Упс! Похоже, места, на которые вы рассчитывали, уже заняты! Мне очень жаль. У нас осталось мест: %s.\n"
	dict["Review"] = "Хотите пересмотреть свои планы?"
	dict["go-ahead"] = "Да, хочу"
	dict["no-deleteAll"] = "Нет, не хочу"
	dict["Can'tUnderstend"] = "К сожалению, я вас не понимаю. Выберите опцию"
	dict["Schedule"] = "<b>%d.</b> <em>Спорт:</em> <b>%s</b>, <em>Дата:</em> <b>%s</b>, <em>Время:</em> <b>%s</b>, <em>Осталось свободных мест:</em> <b>%d</b>, <em>Цена за одно место:</em> <b>%d %s</b>\n\n\n"
	dict["NoGames"] = "Игра пока что нет\n\n"
	dict["WeDon'tHaveAnyGames"] = "Расписание полностью пусто! Проверьте позже\n\n"
	dict["UpAndUn"] = "Вы хотите загрузить или посмотреть?"
	dict["UnloadGames"] = "Игр на которые можно загрузить фотографии пока что нет. Сейчас можно только посмотреть\n"
	dict["UploadGames"] = "Игр с уже загруженными медиафайлами еще нет, но вы можете стать первым!\n"
	dict["WaitForYourFiles"] = "Прекрасно! Вы выбрали игру. У нее осталось %d свободных мест для медиафайлов (максимум 20). Присылайте ваших фотографий или видео, которые были сделаны на этой игре. За один раз можно прислать не более 10 фалов. Кончено же, желательно, чтобы это число не привышало свободных мест"
	dict["HereYouAre"] = "Вот все медиа по это игре"
	dict["NotEnoughSpace"] = "К сожалению, похоже что у меня не осталось свободного места для загрузки ваших файлов на эту игру"
	dict["unload"] = "Посмотреть"
	dict["upload"] = "Загрузить"
	dict["Succesful"] = "Все файлы успешно загружены"

}

func en(dict map[string]string) {
	dict["volleyball"] = "Volleyball"
	dict["football"] = "Football"
	dict["WelcomeToBot"] = "Welcome to our bot! This bot is designed for registration for sports games in Istanbul, but first, you need to register with us!"
	dict["reg"] = "Register"
	dict["BotRules"] = `Welcome to our friendly company! Next, I will write you some rules and recommendations on how to use the bot:
	1. View Schedule.
			- Extremely simple. You can see the full schedule of all possible games.
	2. Game Registration.
			-The bot will help you sign up for the game you are interested in.
	3. View Photos.
			-Here will be stored all photos, videos, gifs, and everything else that other users will post from past games.
	4. Settings.
			-By clicking here, you can see the games you are expecting. Also, here you can cancel the registration.
	
				
	P.S. If anything, the bot has a magical command /menu. This command will always take you back to the main menu. 
	Naturally, if you send it when you have some unfinished process in the bot, the progress will not be saved`
	dict["allright"] = "All clear!"
	dict["WelcomeToMainMenu"] = "Welcome to the main menu"
	dict["first"] = "View schedule"
	dict["second"] = "Game registration"
	dict["third"] = "Our photos and videos"
	dict["fourth"] = "Settings | My games"
	dict["fourth"] = "Settings | My games"
	dict["ChooseAnyGame"] = "Choose the game you are interested in"
	dict["freeSpace"] = "Available seats left:"
	dict["next"] = ">>"
	dict["previous"] = "<<"
	dict["ChooseSeats"] = "Choose or enter the desired number of seats. Please note that the price for one seat for this game is %d %s"
	dict["NoMoreSeats"] = "Unfortunately, you entered either too many available seats for the game that do not exist"
	dict["1"] = "I will be alone"
	dict["2"] = "There will be two of us"
	dict["3"] = "There will be three of us"
	dict["ChoosePaymethod"] = "Choose a payment method"
	dict["payByCard"] = "Pay by card"
	dict["payByCash"] = "Pay by cash"
	dict["MainMenu"] = "Main Menu"
	dict["WaitForYourMoney"] = `You need to transfer %d %s. Above, I sent you a QR code, which you can use to proceed and make the payment, or click the "Pay" button below`
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
	https://www.google.com/maps?q=%s,%s`
	dict["SeatsAreFull"] = "Oops! It seems the seats you were counting on are already taken! I'm sorry. We have seats left: %s.\n"
	dict["Review"] = "Would you like to reconsider your plans?"
	dict["go-ahead"] = "Yes, I want"
	dict["no-deleteAll"] = "No, I don't want"
	dict["Can'tUnderstend"] = "Unfortunately, I don't understand you. Please choose an option"
	dict["Schedule"] = "<b>%d.</b> <em>Sport:</em> <b>%s</b>, <em>Date:</em> <b>%s</b>, <em>Time:</em> <b>%s</b>, <em>Remaining seats:</em> <b>%d</b>, <em>Price per seat:</em> <b>%d %s</b>\n\n\n"
	dict["NoGames"] = "The game is not available yet\n\n"
	dict["WeDon'tHaveAnyGames"] = "The schedule is completely empty! Check back later\n\n"
	dict["UpAndUn"] = "Do you want to upload or view?"
	dict["UnloadGames"] = "There are currently no games where you can upload photos. Right now, you can only view them\n"
	dict["UploadGames"] = "There are no games with already uploaded media files yet, but you can be the first!\n"
	dict["WaitForYourFiles"] = "Great! You've selected the game. It has %d free slots for media files (maximum 20). Send your photos or videos taken in this game. You can send up to 10 files at a time. Of course, it's preferable that this number doesn't exceed the available slots"
	dict["HereYouAre"] = "Here is all the media for this game"
	dict["NotEnoughSpace"] = "Unfortunately, it seems that I don't have enough space left to upload your files for this game"
	dict["unload"] = "View"
	dict["upload"] = "Upload"
	dict["Succesful"] = "All files have been successfully uploaded"
}

func tur(dict map[string]string) {
	dict["volleyball"] = "Voleybol"
	dict["football"] = "Futbol"
	dict["WelcomeToBot"] = "Botumuza hoş geldiniz! Bu bot, İstanbul'daki spor oyunlarına kaydolmak için tasarlanmıştır, ancak önce bizimle kayıt olmanız gerekiyor!"
	dict["reg"] = "Kayıt Ol"
	dict["BotRules"] = `Hoş geldiniz! Arkadaşça bir şirkete hoş geldiniz! Aşağıda size bazı kurallar ve botu nasıl kullanacağınıza dair öneriler yazacağım:

	1. Programı Görüntüle.
			-Son derece basit. Tüm olası oyunların tam programını görebilirsiniz.
	2. Oyun Kaydı.
			-Bot, ilgilendiğiniz oyun için kaydolmanıza yardımcı olacaktır.
	3.Fotoğrafları Görüntüle.
			-Burada, diğer kullanıcıların geçmiş oyunlardan paylaşacakları tüm fotoğraflar, videolar, gif'ler ve diğer şeyler saklanacaktır.
	4.Ayarlar.
			-Buraya tıklayarak beklediğiniz oyunları görebilirsiniz. Ayrıca burada kaydınızı iptal edebilirsiniz.
			
	
	Not: Botun sihirli bir komutu var: /menu. Bu komut sizi her zaman ana menüye götürecektir. 
	Elbette, eğer bot içinde tamamlanmamış bir işlem varsa, ilerleme kaydedilmeyecektir`
	dict["allright"] = "Tamam"
	dict["WelcomeToMainMenu"] = "Ana menüye hoş geldiniz"
	dict["first"] = "Programı görüntüle"
	dict["second"] = "Oyun kaydı"
	dict["third"] = "Bizim fotoğraflar ve videolar"
	dict["fourth"] = "Ayarlar | Oyunlarım"
	dict["ChooseAnyGame"] = "İlgilendiğiniz oyunu seçin"
	dict["freeSpace"] = "Kalan boş koltuklar:"
	dict["next"] = ">>"
	dict["previous"] = "<<"
	dict["ChooseSeats"] = "İstenilen koltuk sayısını seçin veya girin. Lütfen unutmayın ki bu oyun için bir koltuk fiyatı %d %s'dir"
	dict["NoMoreSeats"] = "Maalesef, var olmayan oyun için fazla sayıda boş koltuk girdiniz"
	dict["1"] = "Ben yalnız olacağım"
	dict["2"] = "İkimiz olacağız"
	dict["3"] = "Üç kişi olacağız"
	dict["ChoosePaymethod"] = "Ödeme yöntemi seçin"
	dict["payByCard"] = "Online ödeme"
	dict["payByCash"] = "Nakit, yöneticiye"
	dict["MainMenu"] = "Ana Menü"
	dict["WaitForYourMoney"] = `%d %s göndermeniz gerekiyor. Yukarıda size bir QR kodu gönderdim, bunu kullanarak ödeme yapabilir veya aşağıdaki "Ödeme Yap" düğmesine tıklayabilirsiniz`
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
	
	❤️❤️❤️Sizi bekliyoruz: %s
	\nhttps://www.google.com/maps?q=%s,%s`
	dict["SeatsAreFull"] = "Hoop! Hesapladığınız koltuklar maalesef zaten dolu! Üzgünüm. Kalan yerlerimiz: %s.\n"
	dict["Review"] = "Planlarınızı tekrar düşünmek ister misiniz?"
	dict["go-ahead"] = "Evet, istiyorum"
	dict["no-deleteAll"] = "Hayır, istemiyorum"
	dict["Can'tUnderstend"] = "Üzgünüm, sizi anlamıyorum. Lütfen bir seçenek seçin"
	dict["Schedule"] = "<b>%d.</b> <em>Spor:</em> <b>%s</b>, <em>Tarih:</em> <b>%s</b>, <em>Saat:</em> <b>%s</b>, <em>Kalan boş yerler:</em> <b>%d</b>, <em>Bir koltuk için fiyat:</em> <b>%d %s</b>\n\n\n"
	dict["NoGames"] = "Oyun henüz mevcut değil\n\n"
	dict["WeDon'tHaveAnyGames"] = "Program tamamen boş! Daha sonra kontrol edin\n\n"
	dict["UpAndUn"] = "Yüklemek mi istiyorsunuz yoksa görmek mi?"
	dict["UnloadGames"] = "Şu anda fotoğraf yükleyebileceğiniz oyun yok. Şu anda sadece onları görüntüleyebilirsiniz\n"
	dict["UploadGames"] = "Henüz yüklenmiş medya dosyaları olan oyunlar yok, ancak siz ilk olan olabilirsiniz!\n"
	dict["WaitForYourFiles"] = "Harika! Oyunu seçtiniz. Onun için %d medya dosyası için boş yer var (maksimum 20). Bu oyunda çekilen fotoğraflarınızı veya videolarınızı gönderin. Bir seferde en fazla 10 dosya gönderebilirsiniz. Tabii ki, bu sayının boş yerleri aşmaması tercih edilir"
	dict["HereYouAre"] = "Bu oyun için tüm medya burada"
	dict["NotEnoughSpace"] = "Maalesef, görünüşe göre bu oyun için dosyalarınızı yüklemek için yeterli alanım kalmadı"
	dict["unload"] = "Görüntüle"
	dict["upload"] = "Yükle"
	dict["Succesful"] = "Tüm dosyalar başarıyla yüklendi."
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
