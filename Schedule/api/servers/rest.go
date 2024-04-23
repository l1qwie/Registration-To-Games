package servers

import (
	"Schedule/app"
	"Schedule/apptype"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// make []string of errors (missing data)
func whatMiss(req *apptype.Request) []string {
	m := make([]string, 6)
	if req.Id == 0 {
		m[0] = "req.Id = 0"
	}
	if req.Act == "" {
		m[1] = "req.Act = ``"
	}
	if req.Language == "" {
		m[2] = "req.Language = ``"
	}
	return m
}

// make string
func fromArrToStr(mes []string) string {
	var message string
	for i, m := range mes {
		if m != "" {
			if i == len(mes) {
				message += m
			} else {
				message += fmt.Sprintf("%s, ", m)
			}
		}
	}
	return message
}

// error message wording
// start from []string and then make it to string
func mesofErr(req *apptype.Request) string {
	m := whatMiss(req)
	return fromArrToStr(m)
}

// check for an error
// return answer (string) and found (bool)
func checkError(req *apptype.Request) (string, bool) {
	var (
		mes string
		f   bool
	)
	if req.Id == 0 || req.Act == "" || req.Language == "" {
		mes = mesofErr(req)
		f = true
	}
	return mes, f
}

// start the server
// main logic of the server
func Schedule() {
	router := gin.Default()
	router.POST("/Schedule", func(c *gin.Context) {
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
	router.Run(":8079")
}
