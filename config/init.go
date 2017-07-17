package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"fmt"
	"log"

)

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
