package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["reg"] = "Зарегистрироваться"
	dict["WelcomeToBot"] = "Добро пожаловать в нашего бота! Этот бот предназначен для регистрации на спортивные игры в Стамбуле, но для начала вам нужно зарегистрироваться у нас!"
	dict["allright"] = "Все понятно!"
	dict["reg"] = "Зарегистрироваться"
	dict["BotRules"] = `Добро Пожаловать в нашу дружную компанию! Далее я вам напишу некоторые правила, и рекомендации как пользоваться ботом:\n1. Просмотр расписания.\n     
	- Предельно просто. Вы сможете посмтореть полное расписание всех возможных игр.\n2. Регистрация на игру.\n      
	- Бот поможет вам записаться на интересующую вас игру.\n3. Посмотреть фотографии.\n      
	- Тут будут храниться все фотографии, видео, гифки и все остальное, что будут выкладывать другие пользователи с уже прошедших игр.\n4. Посмототреть мои записи.\n      
	- Нажав сюда, вы сможете посмотреть игры, которые вы ожидаете. Так же тут вы сможете отменить запись.\n\n\n
	P.S. Если что, то у бота есть волшебная команда /menu. Эта команда всегда сможет вас вернуть в главное меню. 
	Естесвенно, если вы пришлете ее, когда у вас есть какой то незаконченный процес в боте, то прогресс не сохраниться`
	dict["first"] = "Просмотр расписания"
	dict["second"] = "Регистрация на игру"
	dict["third"] = "Наши фото и видео"
	dict["fourth"] = "Настройки | Мои игры"
	dict["WelcomeToMainMenu"] = "Добро пожаловать в гланое меню"
}

func en(dict map[string]string) {
	dict["reg"] = "Register"
	dict["WelcomeToBot"] = "Welcome to our bot! This bot is designed for registration for sports games in Istanbul, but first, you need to register with us!"
	dict["allright"] = "All clear!"
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
	dict["first"] = "View schedule"
	dict["second"] = "Game registration"
	dict["third"] = "Our photos and videos"
	dict["fourth"] = "Settings | My games"
	dict["WelcomeToMainMenu"] = "Welcome to the main menu"
}

func tur(dict map[string]string) {
	dict["reg"] = "Kayıt Ol"
	dict["WelcomeToBot"] = "Botumuza hoş geldiniz! Bu bot, İstanbul'daki spor oyunlarına kaydolmak için tasarlanmıştır, ancak önce bizimle kayıt olmanız gerekiyor!"
	dict["allright"] = "Tamam"
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
	dict["first"] = "Programı görüntüle"
	dict["second"] = "Oyun kaydı"
	dict["third"] = "Bizim fotoğraflar ve videolar"
	dict["fourth"] = "Ayarlar | Oyunlarım"
	dict["WelcomeToMainMenu"] = "Ana menüye hoş geldiniz"
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
