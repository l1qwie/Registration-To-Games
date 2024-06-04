package servers

import (
	"Game/app"
	"Game/apptype"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

// Game starts the server.
// Main logic of the server
func Game() {
	router := gin.Default()
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
	router.Run(":8099")
}
