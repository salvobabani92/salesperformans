package libs

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"github.com/salvobabani92/salesperformans.com/config"
	"net/http"
	"github.com/salvobabani92/salesperformans.com/models"

)

func BiletKontrol(c *gin.Context) {
	var tokenStr string = c.Request.Header.Get(config.TokenHeaderName)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.MySecretKey), nil
	})
	if err == nil && token.Valid {
		var user = models.User{}
		config.DB.First(&user, int64(ParseToken(tokenStr, "id").(float64)))
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Lütfen tekrar sisteme giriş yapınız."})
			c.Abort()
		}
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Lütfen tekrar sisteme giriş yapınız."})
		c.Abort()
	}
}

// context içinde tanımlı olan kullanıcı bilgilerine ve firma bilgilerine erişim sağlanır.
func GetUser_Company(c *gin.Context) (models.User, models.Customer) {
	var tokenStr string = c.Request.Header.Get(config.TokenHeaderName)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.MySecretKey), nil
	})
	if err == nil && token.Valid {
		var user = models.User{}
		config.DB.First(&user, int64(ParseToken(tokenStr, "id").(float64)))

		var customer = models.Customer{}
		config.DB.First(&customer, user.CustomerID)
		return user, customer

	} else {
		return models.User{}, models.Customer{}
	}
}
