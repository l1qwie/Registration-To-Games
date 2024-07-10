package servers

import (
	"Game/app"
	"Game/app/handler"
	"Game/apptype"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Make []string of errors (missing data)
func whatMiss(req *apptype.Request) []string {
	m := make([]string, 6)
	if req.Id == 0 {
		m[0] = `"id" = 0`
	}
	if req.Act == "" {
		m[1] = `"action" = ""`
	}
	if req.Language == "" {
		m[2] = `"language" = ""`
	}
	if req.Limit == 0 {
		m[3] = `"limit" = 0`
	}
	if req.Req == "" {
		m[4] = `"request" = ""`
	}
	return m
}

// Make []string of errors (diffrent data is awaited)
func whatDif(req *apptype.Request) []string {
	m := make([]string, 2)
	if req.Act != "game" {
		m[0] = `"action" isn't equal "game"`
	}
	if (req.Language != "ru") && (req.Language != "en") && (req.Language != "tur") {
		m[1] = `"language" isn't equal "ru" or "en" or "tur"`
	}
	return m
}

// Make string
func fromArrToStr(mes []string) string {
	var message string
	for _, m := range mes {
		if m != "" {
			message += fmt.Sprintf("%s\n", m)
		}
	}
	return message
}

// Error message wording
// Starts from []string and then make it to string
func mesofErr(req *apptype.Request, kind bool) string {
	var m []string
	if kind {
		m = whatMiss(req)
	} else {
		m = whatDif(req)
	}
	return fromArrToStr(m)
}

// Check for an error
// Return answer (string) and found (bool)
func checkError(req *apptype.Request) (mes string, f bool) {
	if (req.Id == 0) || (req.Act == "") || (req.Language == "") || (req.Limit == 0) || (req.Req == "") {
		if req.Req == "" {
			mes = "Not enough data: "
			mes += mesofErr(req, true)
			f = true
		}
	}
	if (req.Act != "game") || (req.Language != "ru" && req.Language != "en" && req.Language != "tur") {
		mes += "Diffrent data is awaited: "
		mes += mesofErr(req, false)
		f = true
	}
	return mes, f
}

func loadPrivateKey() *rsa.PrivateKey {
	// Загрузите свой закрытый ключ
	privateKeyPEM, err := os.ReadFile("server.key")
	if err != nil {
		log.Fatal("Failed to read private key: ", err)
	}
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil {
		log.Fatal("Failed to decode PEM block containing private key")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			log.Fatal("Failed to parse private key: ", err)
		}
	}
	return privateKey.(*rsa.PrivateKey)
}

func encryptData(mes *apptype.KeyResponse, data, key []byte) ([]byte, error) {
	var (
		block cipher.Block
		gcm   cipher.AEAD
		err   error
		nonce []byte
	)
	block, err = aes.NewCipher(key)
	if err == nil {
		gcm, err = cipher.NewGCM(block)
	}
	if err == nil {
		nonce = make([]byte, gcm.NonceSize())
		_, err = io.ReadFull(rand.Reader, nonce)
	}
	if err != nil {
		mes.Status = false
		mes.Message = fmt.Sprintf("1 Error (if exists): %s\n2 Error: %s", mes.Message, "Failed to encrypt response message")
		mes.Err = err
	}
	return gcm.Seal(nonce, nonce, data, nil), err
}

func isItSymKey(mes *apptype.KeyResponse, statreq *int, c *gin.Context) ([]byte, bool) {
	encryptedSymmetricKey := c.Request.Header.Get("X-Symmetric-Key")
	encryptedSymmetricKeyBytes, err := base64.StdEncoding.DecodeString(encryptedSymmetricKey)
	if err != nil {
		mes.Status = false
		mes.Message = "Invalid symmetric key encoding"
		mes.Err = err
		*statreq = http.StatusBadRequest
	}
	return encryptedSymmetricKeyBytes, err == nil
}

func decryptionSymKey(mes *apptype.KeyResponse, statreq *int, privatekey *rsa.PrivateKey, encrypsymkey []byte, ip string) []byte {
	symkey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privatekey, encrypsymkey, nil)
	if err != nil {
		mes.Status = false
		mes.Message = "Failed to decrypt symmetric key"
		*statreq = http.StatusInternalServerError
	} else {
		mes.Status = true
		mes.Message = "The key has been accepted"
		*statreq = http.StatusOK
		con := new(handler.Conn)
		con.Db = apptype.ConnectToDatabase()
		con.SaveTheKey(symkey, ip)
	}
	return symkey
}

func prepareResponse(mes *apptype.KeyResponse) []byte {
	ok := true
	resp, err := json.Marshal(mes)
	if err != nil {
		for i := 0; i < 3 && err != nil; i++ {
			mes.Message = fmt.Sprintf("There's the first error while serializing in Json: %v", err)
			resp, err = json.Marshal(mes)
			if i+1 == 3 && err != nil {
				ok = false
			}
		}
		if !ok {
			resp = []byte("SOMETHING REALLY BAD IN THE SERVER! CANNOT WORK NOW!")
		}
	}
	return resp
}

func KeyManager() {
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	privatekey := loadPrivateKey()
	router.POST("/postSymKey", func(c *gin.Context) {
		var (
			statreq                  int
			symmetricKey, encryptMes []byte
			response                 any
		)
		message := new(apptype.KeyResponse)
		encrypsymkey, ok := isItSymKey(message, &statreq, c)
		if ok {
			symmetricKey = decryptionSymKey(message, &statreq, privatekey, encrypsymkey, c.ClientIP())
		}
		resp := prepareResponse(message)
		encryptMes, err = encryptData(message, resp, symmetricKey)
		if err != nil {
			response = prepareResponse(message)
		} else {
			response = base64.StdEncoding.EncodeToString(encryptMes)
		}
		c.JSON(statreq, response)
	})
	err = router.RunTLS(":8444", "server.crt", "server.key")
	if err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

// Game starts the server.
// Main logic of the server
func Game() {
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	router.POST("/Game", func(c *gin.Context) {
		req := new(apptype.Request)
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		} else {
			_ = json.Unmarshal(body, &req)
			mes, found := checkError(req)
			if found {
				c.JSON(http.StatusOK, gin.H{"error": mes})
			} else {
				resp := app.Receiving(req)
				c.JSON(http.StatusOK, resp)
			}
		}
	})
	certFile := "server.crt"
	keyFile := "server.key"

	log.Println("Starting HTTPS server on :8099")
	err = router.RunTLS(":8099", certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}
}

func getSymKey(ip string, statreq *int, mes *string) ([]byte, error) {
	con := new(handler.Conn)
	con.Db = apptype.ConnectToDatabase()
	key, err := con.GetTheKey(ip)
	if err != nil {
		*statreq = http.StatusInternalServerError
		*mes = "Failed to retrieve symmetric key"
	}
	return key, err
}

func decryptData(data, key []byte, statreq *int, mes *string) ([]byte, error) {
	var (
		gcm                        cipher.AEAD
		noncesize                  int
		nonce, ciphertext, newdata []byte
	)
	block, err := aes.NewCipher(key)
	if err == nil {
		gcm, err = cipher.NewGCM(block)
	}
	if err == nil {
		noncesize := gcm.NonceSize()
		if len(data) < noncesize {
			err = fmt.Errorf("ciphertext too short")
		}
	}
	if err == nil {
		nonce, ciphertext = data[:noncesize], data[noncesize:]
	}
	if err == nil {
		newdata, err = gcm.Open(nil, nonce, ciphertext, nil)
	}
	if err != nil {
		*statreq = http.StatusInternalServerError
		*mes = "Failed to decrypt data"
	}
	return newdata, err
}

func Game2() {
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	router.POST("/Game", func(c *gin.Context) {
		var (
			data, symkey []byte
			response     any
			statreq      int
			mes          string
		)
		req := new(apptype.Request)
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			statreq = http.StatusBadRequest
			response = gin.H{"error": err}
		} else {
			symkey, err = getSymKey(c.ClientIP(), &statreq, &mes)
		}
		if err == nil {
			data, err = decryptData(body, symkey, &statreq, &mes)
		}
		if err == nil {
			err = json.Unmarshal(data, req)
			if err != nil {
				statreq = http.StatusBadRequest
				mes = fmt.Sprintf("%s", err)
			}
		}
		if err == nil {
			meserr, found := checkError(req)
			if found {
				statreq = http.StatusOK
				mes = meserr
			} else {
				response = app.Receiving(req)
			}
		}
		if err != nil {
			response = gin.H{"error": mes}
		}
		c.JSON(statreq, response)
	})

	certFile := "server.crt"
	keyFile := "server.key"

	log.Println("Starting HTTPS server on :8099")
	err = router.RunTLS(":8099", certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}
}

/*
router.POST("/Game", func(c *gin.Context) {
    // Получаем IP-адрес клиента
    clientIP := c.ClientIP()
    log.Printf("Request received from IP: %s", clientIP)

    req := new(apptype.Request)
    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }
    err = json.Unmarshal(body, req)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
    }

    symKey, err := getSymmetricKey()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve symmetric key"})
        return
    }

    // Расшифровка данных (если требуется)
    // decryptedData, err := decryptData(req.Data, symKey)
    // if err != nil {
    //     c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decrypt data"})
    //     return
    // }

    // Обработка запроса
    mes, found := checkError(req)
    if found {
        c.JSON(http.StatusOK, gin.H{"error": mes})
    } else {
        resp := app.Receiving(req)
        c.JSON(http.StatusOK, resp)
    }
})*/
