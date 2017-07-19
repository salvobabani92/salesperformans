

package models

import (
	"github.com/salvobabani92/salesperformans.com/config"
	"time"
)


// Satış başlığı Tablosu.
// swagger: model
type SalesHeader struct {
	// ID
	ID        uint `gorm:"primary_key" json:"id"`
	// Oluşturma Tarihi
	CreatedAt time.Time `json:"-"`
	// Son Değişiklik tarihi
	UpdatedAt time.Time `json:"-"`
	// Numarası
	No  string `json:"no"`
	//Müşteri Numarası
	CustomerNo string  `json:"customer_no"`
	//Müşteri İsmi
	CustomerName string `json:"customer_name"`
	//Müşteri Kimliği
	CustomerID  uint `json:"customer_id"`
	//Yüzdelik
	Amount   int64 `json:"amount"`
}

func (this SalesHeader) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}


// Satış Başlığı bilgileri geridönüş değeri
// swagger:response
type LocationResponse struct {
	// in:body
	Body SalesHeader
}

// Satış Başlığı bilgileri
// swagger:parameters Satış Başlığı-Oluştur
type SalesHeaderInputPatch struct {
	// Satış Başlığı Bilgileri
	// in:body
	SalesHeader SalesHeader
}

// Satış Başlığı Güncelle
// swagger:parameters Satış Başlığı-Güncelle
type SalesHeaderInputPUT struct {
	// Satış Başlığı Bilgileri
	// in:body
	SalesHeader SalesHeader
}


// Satış Başlığı Bilgisini Getir
// swagger:parameters Satış Başlığı-Getir
type SalesHeaderIDQueryGET struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Satış Başlığı Bilgisini Güncelle
// swagger:parameters Satış Başlığı-Güncelle
type SalesHeaderIDQueryPUT struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Satış Başlığı Bilgisini Sil
// swagger:parameters Satış Başlığı-Sil
type SalesHeaderIDQueryDELETE struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}

// Satış Başlığı bilgilerini excel'den yükle
// swagger:parameters Satış Başlığı-Excel-Yükle
type UploadSalesHeaderExcelInput struct {
	// Satış Başlığı bilgilerini excel'den yükle
	// in:body
	Body struct{ ExcelFile string `json:"ExcelFile"` }
}


// Satış Başlığı Yükle
// swagger:parameters Satış Başlığı-Json-Array
type SalesHeaderArrayInputPOST struct {
	// Satış Başlığı Bilgileri
	// in:body
	SalesHeader [] SalesHeader
}
