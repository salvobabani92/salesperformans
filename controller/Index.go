

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func  Index(c *gin.Context) {

	c.HTML(http.StatusOK, "index", gin.H{
		"user":   "q",
		"userID": "1",
		"Title":"Hello",
	})
}
