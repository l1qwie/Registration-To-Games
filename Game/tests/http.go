package tests

import (
	"Game/apptype"
	"Game/tests/functional"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io"
	"log"
	"net/http"
	"os"
)

// Gegerates a symmetrick key and loads the public server's key
func generateAndLoad() ([]byte, *rsa.PublicKey, error) {
	//AES-256
	symkey := make([]byte, 32)
	_, err := rand.Read(symkey)
	if err != nil {
		panic(err)
	}
	pubKeyPEM, err := os.ReadFile("server.crt")
	if err != nil {
		log.Fatal("Failed to read public key: ", err)
	}
	block, _ := pem.Decode(pubKeyPEM)
	if block == nil {
		log.Fatal("Failed to decode PEM block containing public key")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal("Failed to parse public key: ", err)
	}
	pubkey := cert.PublicKey.(*rsa.PublicKey)
	return symkey, pubkey, err
}

// Creates HTTPS client with settings
func createClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // only for testing
			},
		},
	}
	return client
}

func encryptSymKey(symKey []byte, pubKey *rsa.PublicKey) ([]byte, error) {
	encryptedKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, symKey, nil)
	return encryptedKey, err
}

func decryptData(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func callgame(body []byte) *apptype.Response {
	symkey, pubkey, _ := generateAndLoad()
	encryptedsymkey, err := encryptSymKey(symkey, pubkey)
	if err != nil {
		log.Fatal("Failed to encrypt symmetric key: ", err)
	}
	client := createClient()
	req, err := http.NewRequest("POST", "https://localhost:8443/hello", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Failed to create request: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Symmetric-Key", base64.StdEncoding.EncodeToString(encryptedsymkey))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed to do request: ", err)
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response: ", err)
	}

	// Расшифровка ответа сервера
	res := new(apptype.Response)
	err = json.Unmarshal(respbody, &res)
	if err != nil {
		log.Fatal("Failed to unmarshal response: ", err)
	}

	// Декодирование зашифрованного сообщения из base64
	encryptedMessage, err := base64.StdEncoding.DecodeString(res.Message)
	if err != nil {
		log.Fatal("Failed to decode response message: ", err)
	}

	// Расшифровка данных
	decryptedMessage, err := decryptData(encryptedMessage, symkey)
	if err != nil {
		log.Fatal("Failed to decrypt response message: ", err)
	}

	// Замена зашифрованного сообщения на расшифрованное в структуре
	res.Message = string(decryptedMessage)

	// Вывод расшифрованного сообщения
	log.Print("Response:", res)
	return res
}

func createGame() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	defer deleteTGame()
	ts := new(TestStuct)
	ts.Round = 10
	ts.Name = "Create-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendSport, sendDate, sendTime, sendSeats, sendPrice, sendCurrency, sendLink, sendAddress, sendSave}
	ts.FuncRes = []func(*apptype.Response){functional.SelectSport, functional.SelectDate, functional.SelectTime,
		functional.SelectSeats, functional.SelectPrice, functional.SelectCurrency, functional.SelectLink,
		functional.SelectAddress, functional.SemiFinal, functional.Final}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, trash2, trash3, trash4, trash5, trash6, trash7, trash8, trash9, trash10, trash11, trash18, trash19}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ts.DoTest()
}

func changeSport() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeSport-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseSport, chsendSport, sendChSaveSport}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChSport, functional.ChSemiFinalSport, functional.ChFinalSport}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrash8, chtrash9, chtrash10, chtrash11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeDate() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeDate-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseDate, chsendDate, sendChSaveDate}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChDate, functional.ChSemiFinalDate, functional.ChFinalDate}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashdate8, chtrashdate9, chtrashdate10, chtrashdate11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeTime() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeTime-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseTime, chsendTime, sendChSaveTime}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChTime, functional.ChSemiFinalTime, functional.ChFinalTime}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashtime8, chtrashtime9, chtrashtime10, chtrashtime11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeSeats() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeSeats-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseSeats, chsendSeats, sendChSaveSeats}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChSeats, functional.ChSemiFinalSeats, functional.ChFinalSeats}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashseats8, chtrashseats9, chtrashtseats10, chtrashseats11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changePrice() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangePrice-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, chosePrice, chsendPrice, sendChSavePrice}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChPrice, functional.ChSemiFinalPrice, functional.ChFinalPrice}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashprice8, chtrashprice9, chtrashtprice10, chtrashprice11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeCurency() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	unverifiable = true
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeCurrency-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseCurrency, chsendCurrency, sendChSaveCurrency}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChCurrency, functional.ChSemiFinalCurrency, functional.ChFinalCurrency}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashtcurrency10, chtrashcurrency11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeLink() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	unverifiable = true
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeLink-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseLink, chsendLink, sendChSaveLink}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChLink, functional.ChSemiFinalLink, functional.ChFinalLink}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashtlink10, chtrashlink11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func changeAddress() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	unverifiable = true
	ts := new(TestStuct)
	ts.Round = 6
	ts.Name = "ChangeAddress-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionChange, sendGame, choseAddress, chsendAddress, sendChSaveAddress}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseGame, functional.ChooseChangeable,
		functional.ChAddress, functional.ChSemiFinalAddress, functional.ChFinalAddress}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrash4, chtrash5, chtrash6, chtrash7, chtrashaddress10, chtrashaddress11}
	ts.UpdtLevel = []int{0, 1, 2, 3, 4, 5}
	ts.DoTest()
}

func deleteGame() {
	apptype.Db = apptype.ConnectToDatabase()
	defer apptype.Db.Close()
	createTestGame()
	defer deleteChGame()
	unverifiable = false
	ts := new(TestStuct)
	ts.Round = 3
	ts.Name = "Delete-Game-Test"
	ts.FuncReq = []func() *apptype.Request{sendHello, sendDiretionDel, sendDelGame}
	ts.FuncRes = []func(*apptype.Response){functional.ChooseOneOfThree, functional.ChooseDelGame, functional.FinalDel}
	ts.FuncTrsh = []func() *apptype.Request{trash, trash1, chtrash2, chtrash3, chtrashdel4, chtrashdel5}
	ts.UpdtLevel = []int{0, 1, 2}
	ts.DoTest()
}

func changeGame() {
	changeSport()
	changeDate()
	changeTime()
	changeSeats()
	changePrice()
	changeCurency()
	changeLink()
	changeAddress()
}

func Start() {
	createGame()
	changeGame()
	deleteGame()
}
