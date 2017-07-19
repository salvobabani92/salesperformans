

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/salvobabani92/salesperformans.com/models"
	"github.com/salvobabani92/salesperformans.com/config"
	"github.com/salvobabani92/salesperformans.com/libs"
	"log"

	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var user = models.User{}
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	user.Hash = config.GetMD5Hash(user.Password)
	config.DB.Where("email = ? and hash = ?", user.Email, user.Hash).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized,
			//gin.H{"error": "Kullanıcı adı veya şifre hatalı.."}
			models.GetGenericStatusResponse("401","Kullanıcı adı veya şifre hatalı"),
		)
	} else {
		strToken, err := libs.CreateToken(user)
		if err == nil {
			c.Header(config.TokenHeaderName, strToken)
			timeVal, _ := time.Parse("2006-01-02T15:04:05.000Z", libs.ParseToken(strToken, "expd").(string))

			c.JSON(http.StatusOK,
				models.LoginResp{
					TokenVal:strToken,
					Expire: timeVal,
				},
			)
			//..gin.H{config.TokenHeaderName: strToken, "expire":libs.ParseToken(strToken, "expd")})
		} else {
			log.Println(err)

			c.JSON(http.StatusBadRequest,
				models.GetGenericStatusResponse("400","Şifre oluşturma servisinde bir sorun var."),
			)
		}
	}
}

func Register(c *gin.Context) {
	var user = models.User{}
	config.DB.Where("Email = ?", c.PostForm("email")).Select("id").Find(&user)
	if user.ID == 0 {

		user.Email = c.PostForm("email")
		user.Name = c.PostForm("name")
		user.LastName = c.PostForm("lastname")
		user.Hash = config.GetMD5Hash(c.PostForm("password"))
		user.Active = true
		user.LanguageCode = "TRK"


		if config.DB.NewRecord(user) {
			config.DB.Create(&user)
		}
		config.DB.NewRecord(user)
	}
	c.JSON(http.StatusCreated, models.GetGenericStatusResponse("201","Kaydınız başarı ile alınmıştır."))


}