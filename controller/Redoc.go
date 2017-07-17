

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func  Redoc(c *gin.Context) {

	c.HTML(http.StatusOK, "redoc", nil)
}
