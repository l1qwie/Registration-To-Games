package dict

var Dictionary map[string]map[string]string

func ru(dict map[string]string) {
	dict["first"] = "Просмотр расписания"
	dict["second"] = "Регистрация на игру"
	dict["third"] = "Наши фото и видео"
	dict["fourth"] = "Настройки | Мои игры"
	dict["unload"] = "Посмотреть"
	dict["upload"] = "Загрузить"
	dict["UpAndUn"] = "Вы хотите загрузить или посмотреть?"
	dict["ChooseAnyGame"] = "Выберите интересующую вас игру"
	dict["HereYouAre"] = "Вот все медиа по это игре"
	dict["MainMenu"] = "Главное Меню"
	dict["UnloadGames"] = "Игр на которые можно загрузить фотографии пока что нет. Сейчас можно только посмотреть\n"
	dict["UploadGames"] = "Игр с уже загруженными медиафайлами еще нет, но вы можете стать первым!\n"
	dict["NoMGames"] = "Расписание полностью пусто! Проверьте позже\n\n"
	dict["volleyball"] = "Волейбол"
	dict["football"] = "Футбол"
	dict["WaitForYourFiles"] = "Прекрасно! Вы выбрали игру. У нее осталось %d свободных мест для медиафайлов (максимум 20). Присылайте ваших фотографий или видео, которые были сделаны на этой игре. За один раз можно прислать не более 10 фалов. Кончено же, желательно, чтобы это число не привышало свободных мест"
	dict["NotEnoughSpace"] = "К сожалению, похоже что у меня не осталось свободного места для загрузки ваших файлов на эту игру"
	dict["Succesful"] = "Все файлы успешно загружены"
	dict["Oops!"] = "Ой! Кажется, что-то внутри сломалось! Повторите попытку позже. Если проблема не исчезнет, обратитесь к администратору!"
}

func en(dict map[string]string) {
	dict["first"] = "View schedule"
	dict["second"] = "Game registration"
	dict["third"] = "Our photos and videos"
	dict["fourth"] = "Settings | My games"
	dict["unload"] = "View"
	dict["upload"] = "Upload"
	dict["UpAndUn"] = "Do you want to upload or view?"
	dict["ChooseAnyGame"] = "Choose the game you are interested in"
	dict["HereYouAre"] = "Here is all the media for this game"
	dict["MainMenu"] = "Main Menu"
	dict["UnloadGames"] = "There are currently no games where you can upload photos. Right now, you can only view them\n"
	dict["UploadGames"] = "There are no games with already uploaded media files yet, but you can be the first!\n"
	dict["NoMGames"] = "The schedule is completely empty! Check back later\n\n"
	dict["volleyball"] = "Volleyball"
	dict["football"] = "Football"
	dict["WaitForYourFiles"] = "Great! You've selected the game. It has %d free slots for media files (maximum 20). Send your photos or videos taken in this game. You can send up to 10 files at a time. Of course, it's preferable that this number doesn't exceed the available slots"
	dict["NotEnoughSpace"] = "Unfortunately, it seems that I don't have enough space left to upload your files for this game"
	dict["Succesful"] = "All files have been successfully uploaded"
	dict["Oops!"] = "Oops! It looks like something broke inside! Please try again later. If the problem persists, contact the administrator!"
}

func tur(dict map[string]string) {
	dict["first"] = "Programı görüntüle"
	dict["second"] = "Oyun kaydı"
	dict["third"] = "Bizim fotoğraflar ve videolar"
	dict["fourth"] = "Ayarlar | Oyunlarım"
	dict["unload"] = "Görüntüle"
	dict["upload"] = "Yükle"
	dict["UpAndUn"] = "Yüklemek mi istiyorsunuz yoksa görmek mi?"
	dict["ChooseAnyGame"] = "İlgilendiğiniz oyunu seçin"
	dict["HereYouAre"] = "Bu oyun için tüm medya burada"
	dict["MainMenu"] = "Ana Menü"
	dict["UnloadGames"] = "Şu anda fotoğraf yükleyebileceğiniz oyun yok. Şu anda sadece onları görüntüleyebilirsiniz\n"
	dict["UploadGames"] = "Henüz yüklenmiş medya dosyaları olan oyunlar yok, ancak siz ilk olan olabilirsiniz!\n"
	dict["NoMGames"] = "Program tamamen boş! Daha sonra kontrol edin\n\n"
	dict["volleyball"] = "Voleybol"
	dict["football"] = "Futbol"
	dict["WaitForYourFiles"] = "Harika! Oyunu seçtiniz. Onun için %d medya dosyası için boş yer var (maksimum 20). Bu oyunda çekilen fotoğraflarınızı veya videolarınızı gönderin. Bir seferde en fazla 10 dosya gönderebilirsiniz. Tabii ki, bu sayının boş yerleri aşmaması tercih edilir"
	dict["NotEnoughSpace"] = "Maalesef, görünüşe göre bu oyun için dosyalarınızı yüklemek için yeterli alanım kalmadı"
	dict["Succesful"] = "Tüm dosyalar başarıyla yüklendi"
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
