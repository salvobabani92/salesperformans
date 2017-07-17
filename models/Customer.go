

package models

import (
	"fmt"
	"github.com/salvobabani92/salesperformans.com/config"

)

type Customer struct {
	//Müşteri ID
	ID 				uint `gorm:"primary_key" json:"id"`
	//Müşteri Numarası
	No				uint `json:"no"`
	//Müşteri Adı
	Name			string `json:"name" sql:"type:varchar(250);" CaptionML:"trk=Müşteri Adı;enu=Customer Name"`

}

func (this Customer) CreateTable() {
	config.DB.DropTable(this)
	fmt.Println("Customer Table Dropped")
	config.DB.CreateTable(this)
	fmt.Println("Customer Table Created")

}

// Firma bilgileri geridönüş değeri
// swagger:response
type CustomerResponce struct {
	// in:body
	Body Customer
}

// Firma bilgileri
// swagger:parameters Firma-Bilgileri-Güncelle
type CustomerInputPatch struct {
	// Kullanıcıya ait bilgiler
	// in:body
	Customer Customer
}

