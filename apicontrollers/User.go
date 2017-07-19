package apicontrollers

import (
	"github.com/gin-gonic/gin"
	"github.com/salvobabani92/salesperformans.com/libs"
	"net/http"
)

func GetMe(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	c.JSON(http.StatusOK, user)
}