package server

import (
	"Welcome/app"
	"Welcome/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Make []string of errors (missing data)
func whatMiss(req *types.Request) []string {
	mes := make([]string, 3)
	if req.Id == 0 {
		mes[0] = `id = 0`
	}
	if req.Act == "" {
		mes[1] = `action = ""`
	}
	if req.Language == "" {
		mes[2] = `language = ""`
	}
	return mes
}

// Make []string of errors (diffrent data is awaited)
func whatDif(req *types.Request) []string {
	mes := make([]string, 2)
	if req.Act != "registration" {
		mes[0] = `"action" isn't equal "registration"`
	}
	if req.Language != "ru" && req.Language != "en" && req.Language != "tur" {
		mes[1] = `"language" isn't equal "ru" or "en" or "tur"`
	}
	return mes
}

// Make string
func fromArrToStr(mes []string) (message string) {
	for _, m := range mes {
		if m != "" {
			message += fmt.Sprintf("%s\n", m)
		}
	}
	return message
}

// Error message wording
// Starts from []string and then make it to string
func mesofErr(req *types.Request, kind bool) string {
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
func checkError(req *types.Request) (mes string, f bool) {
	if req.Id == 0 || req.Act == "" || req.Language == "" {
		mes = "Not enough data: "
		mes += mesofErr(req, true)
		f = true
	}
	if (req.Act != "registration") || (req.Language != "ru" && req.Language != "en" && req.Language != "tur") {
		mes += "Diffrent data is awaited: "
		mes += mesofErr(req, false)
		f = true
	}

	return mes, f
}

// Starts the server
func Welcome() {
	router := gin.Default()
	router.POST("/Welcome", func(c *gin.Context) {
		//router get a post request
		req := new(types.Request)
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		} else {
			//main logic
			_ = json.Unmarshal(body, &req)
			fmt.Println(req.Act, "!#!@#!@#!")
			mes, found := checkError(req)
			if found {
				c.JSON(http.StatusOK, gin.H{"error": mes})
			} else {
				resp := app.Receiving(req)
				c.JSON(http.StatusOK, resp)
			}
		}
	})
	router.Run(":8081")
}
