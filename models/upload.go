

package models

import (
"time"
"github.com/salvobabani92/salesperformans.com/config"

)


// Uploads Tablosu.
// swagger: model
type Upload struct {
	// ID
	ID               uint `gorm:"primary_key" json:"id"`
	// Oluşturma Tarihi
	CreatedAt        time.Time `json:"-"`
	// Son Değişiklik tarihi
	UpdatedAt        time.Time `json:"-"`
	// Müşteri ID
	CustomerID        uint    `CaptionML:"enu=Customer ID;trk=Müşteri ID" json:"-"`
	// Müşteri Bilgileri
	Customer          Customer `json:"-"`
	// Kullanıcı ID
	UserID           uint `json:"user_id" CaptionML:"enu=User ID;trk=Kullanıcı ID"`
	// Kullanıcı Bilgisi
	User			User `json:"_"`
	// Relate Table Name
	RelatedTableName string `json:"related_table_name" CaptionML:"enu=Related Table Name;trk=İlişkili Tablo Adı"`
	// Related Record ID
	RelatedRecordID  string `json:"related_record_id" CaptionML:"enu=Related Record ID;trk=İlişkili Kayıt ID"`
	// File Path
	FilePath         string `json:"file_path" CaptionML:"enu=File Path;trk=Dosya Adresi"`
	// File Extension
	FileExtension    string `json:"file_extension" CaptionML:"enu=File Extension;trk=Dosya Uzantısı"`
}

func (this Upload) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}


// Upload bilgileri geridönüş değeri
// swagger:response
type UploadResponse struct {
	// in:body
	Body Upload
}
