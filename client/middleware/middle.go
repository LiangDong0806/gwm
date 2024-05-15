package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zg5/Homework01/common"
)

func Middle(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSONP(http.StatusAccepted, gin.H{
			"code":    http.StatusAccepted,
			"message": "token ItCanTBeEmpty",
		})
		return
	}
	ok, _ := common.GetJwtToken(token)
	if !ok {
		c.JSONP(http.StatusAccepted, gin.H{
			"code":    http.StatusAccepted,
			"message": "token error",
		})
		return
	}
}
