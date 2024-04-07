package enter

import (
	"Welcome/app"
	"Welcome/types"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mesofErr(req types.Request) string {
	mes := `Not enough data: `
	if req.Id == 0 {
		mes += `id = 0`
	}
	if req.Act == "" {
		mes += "action = `` "
	}
	if req.Language == "" {
		mes += "language = `` "
	}
	if req.Connection == nil {
		mes += `connection = nil `
	}
	return mes
}

func Welcome() {
	router := gin.Default()
	router.POST("/Welcome", func(c *gin.Context) {
		//router get a post request
		var req types.Request
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		} else {
			//main logic
			_ = json.Unmarshal(body, &req)
			if req.Id == 0 || req.Act == "" || req.Language == "" || req.Connection == nil {
				c.JSON(http.StatusOK, gin.H{"error": mesofErr(req)})
			} else {
				resp := app.Receiving(&req)
				c.JSON(http.StatusOK, resp)
			}
		}
	})

	router.Run(":8081")
}
