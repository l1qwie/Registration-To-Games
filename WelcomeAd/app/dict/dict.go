package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["SayHello"] = "Добро пожаловать в нашего бота! Этот бот может пригодиться вам, если вы приобрели его первую часть и вам требуется административная часть. Если это ваш случай, то нажмите на кнопку снизу"
	dict["reg"] = "Зарегистрироваться"
	dict["LetsStart"] = "Начнем же!"
	dict["AdminRules"] = `Вот пара рекомендаций для того, как пользоваться этим ботом:
	
	1. <b>Игры</b>         
			-Тут вы сможете сделать все что угодно с вашими играми (создание, настройка, удаление и тд.)
	2. <b>Клиенты</b>         
			-Тут вы сможете сделать все со своими клиентами (создать, поменять, удалить)
	3. <b>Активность</b>
	        -Тут вы сможете сделать своеобразную игру для чата ваших клиентов 
			 (суть игры в том, что ваши клиенты должны будут наперегонки регистрироваться на игру). 
			 А так же тут вы сможете посмотреть куда и кто записался, а так же найти контакты нужного вам человека
	4. <b>Деньги</b>
	        -Здесь вы сможете посмотреть кто вам должен за игры, от кого ожидается оплата, а кто уже оплатил. 
			 Так же именно тут вы сможете увидеть небольшую статистику по всей финансовым оборотам в боте
	5. <b>Настройки</b>          
			-Тут вы сможете изменить настрйоки бота. Например язык`
	dict["first"] = "Игры"
	dict["second"] = "Клиенты"
	dict["third"] = "Активность"
	dict["fourth"] = "Деньги"
	dict["fifth"] = "Настройки"
	dict["MainMenu"] = "Главное Меню"
}

func en(dict map[string]string) {
	dict["SayHello"] = "Welcome to our bot! This bot can be useful if you have purchased its first part and need the administrative section. If this is your case, then click the button below"
	dict["reg"] = "Register"
	dict["LetsStart"] = "Let's get started!"
	dict["AdminRules"] = `Here are a couple of recommendations on how to use this bot:
	
	1. <b>Games</b>          
			-Here, you can do anything with your games (create, configure, delete, etc.)
	2. <b>Clients</b>          
			-Here, you can manage everything related to your clients (create, modify, delete)
	3. <b>Activity</b>          
			-Here you can create a unique game for your clients in the chat 
			 (the essence of the game is that your clients will have to register for the game as quickly as possible). 
			 Also, here you can see where and who has registered, as well as find the contacts of the person you need
	4. <b>Finances</b>          
			-Here you can see who owes you for the games, who payment is expected from, and who has already paid. 
			 Also, right here you can view a small statistics on all financial transactions in the bot
	5. <b>Settings</b>          
			-Here, you can change the bot settings. For example, the language`
	dict["first"] = "Games"
	dict["second"] = "Clients"
	dict["third"] = "Activity"
	dict["fourth"] = "Finances"
	dict["fifth"] = "Settings"
	dict["MainMenu"] = "Main Menu"
}

func tur(dict map[string]string) {
	dict["SayHello"] = "Hoş geldiniz! Botumuza hoş geldiniz! Bu bot, eğer ilk bölümünü satın aldıysanız ve yönetim bölümüne ihtiyacınız varsa, size yardımcı olabilir. Eğer durumunuz bu ise, o zaman aşağıdaki düğmeye tıklayın"
	dict["reg"] = "Kayıt Ol"
	dict["LetsStart"] = "Hadi başlayalım!"
	dict["AdminRules"] = `İşte bu botu nasıl kullanacağınıza dair birkaç öneri:
	
	1. <b>Oyunlar</b>          
			-Burada oyunlarınızla ilgili her şeyi yapabilirsiniz (oluşturmak, yapılandırmak, silmek vb.)
	2. <b>Müşteriler</b>         
			-Burada müşterilerinizle ilgili her şeyi yönetebilirsiniz (oluşturmak, değiştirmek, silmek)
	3. <b>Aktivite</b>        
			-Burada müşterileriniz için sohbet içinde özgün bir oyun oluşturabilirsiniz 
			 (oyunun özü, müşterilerinizin oyuna mümkün olan en kısa sürede kaydolmaları gerektiğidir). 
			 Ayrıca, burada kimin ve nereye kaydolduğunu görebilir, aynı zamanda ihtiyacınız olan kişinin iletişim bilgilerini bulabilirsiniz
	4. <b>Finans</b>         
			-Burada size oyunlar için kim borçlu, kimden ödeme bekleniyor ve kim ödeme yaptı, görebilirsiniz. 
			 Ayrıca, tam burada, bot içindeki tüm finansal işlemlere ait küçük bir istatistiği görebilirsiniz
	5. <b>Ayarlar</b>          
			-Burada botun ayarlarını değiştirebilirsiniz. Örneğin, dil`
	dict["first"] = "Oyunlar"
	dict["second"] = "Müşteriler"
	dict["third"] = "Aktivite"
	dict["fourth"] = "Finans"
	dict["fifth"] = "Ayarlar"
	dict["MainMenu"] = "Ana menüye"
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
