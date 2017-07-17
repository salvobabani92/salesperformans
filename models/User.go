


package models

import (
	"time"
	"github.com/salvobabani92/salesperformans.com/config"
)


// Firmadaki IK kullanıcılarının tanımlandığı tablodur. Bu tabloda personeller tutulmaz.
// Sadece sistemi admin yetkisi ile kullanacak olan kullanıcıların bilgileri tutulmaktadır.
// swagger:model
type User struct {
	// ID
	ID           uint `gorm:"primary_key"`
	// Oluşturma Tarihi
	CreatedAt    time.Time `json:"-"`
	// Son Değişiklik tarihi
	UpdatedAt    time.Time `json:"-"`
	// Kullanıcı Mail Adresi
	// required: true
	Email        string `json:"email" sql:"type:varchar(120)" CaptionML:"enu=Email;trk=Email"`
	// Kullanıcının Adı
	Name         string `json:"name" option:"value" sql:"type:varchar(120);" CaptionML:"enu=Name;trk=Ad"`
	// Kullanıcının Soyadı
	LastName     string `json:"last_name" sql:"type:varchar(120);" CaptionML:"enu=Last Name;trk=Soyadı"`
	// Bu bilgi gizlidir. şifresi hashlenmiş olarak bu alanda tutulur ve sistemde kaydedilir.
	Hash         string `json:"-" sql:"type:varchar(100);" CaptionML:"enu=Hash;trk=Şifre"`
	// Kullanıcının şifresi bu alan tabloda kaydedilmez.
	Password     string `json:"-" sql:"-" ss:"-" CaptionML:"enu=Password;trk=Şifre"`
	// Firma Adı, kullanıcının hangi firmanın IK personeli olduğu bu tabloda tutulur.
	CustomerID    uint    `json:"-" sql:"index" CaptionML:"enu=Customer;trk=Müşteri"`
	// Firma Bilgileri
	Customer      Customer `json:"-"`
	// Kullanıcı şu anda aktif ise true, kullanıcı aktif değilse false olarak tutulur.
	Active       bool  `sql:"default:false" CaptionML:"enu=Active;trk=Aktif"`
	// Kullanıcı sistemi hangi dil kodu ile kullanıyor. Varsayılan değeri TRK (Türkçe)
	LanguageCode string `json:"lang_code" CaptionML:"enu=Language Code;trk=Dil Kodu"`
}

// Kullanıcı bilgileri
// swagger:response
type UserResponce struct {
	// in:body
	Body User
}


func (this User) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}