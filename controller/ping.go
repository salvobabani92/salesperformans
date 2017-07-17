

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/salvobabani92/salesperformans.com/models"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, models.Pong{Message: "pong..."})
}

