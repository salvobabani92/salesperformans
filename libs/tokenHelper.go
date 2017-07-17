package libs

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/salvobabani92/salesperformans.com/models"
	"time"
	"github.com/salvobabani92/salesperformans.com/config"
)


func CreateToken(user models.User) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	Claims := make(jwt.MapClaims)
	Claims["id"] = user.ID
	Claims["exp"] = time.Now().Add(config.TokenExpireDuration).Unix()
	Claims["expd"] = time.Now().Add(config.TokenExpireDuration).Format("2006-01-02T15:04:05.000Z")
	Claims["iat"] = time.Now().Unix()
	token.Claims = Claims
	tokenString, err := token.SignedString([]byte(config.MySecretKey))
	return tokenString, err
}

// Parse with Token
func ParseToken(myToken string, ClaimsName string) (interface{}) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.MySecretKey), nil
	})

	if err == nil && token.Valid {
		return token.Claims.(jwt.MapClaims)[ClaimsName]
	} else {
		return ""
	}
}

