package config

import (
	"time"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/lib/pq"

)

var MySecretKey = "zQ!a11p+8xbfSU?2y30NA_zura*3AB-777"
// Kaç günlük geçerli bilet üretilecek
var TokenExpireDuration = time.Hour * 24 * 40
var TokenHeaderName = "Bilet"
var DB *gorm.DB

func InitDB() {
	cnnString := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s sslmode=disable", "localhost", "postgres", "123456", "salesperformans")


	var err error
	DB, err = gorm.Open("postgres", cnnString)
	if err != nil {
		log.Println("DB Error", err)
	}
	log.Println("DB Connected")
	//AutoMigrate()

}
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}