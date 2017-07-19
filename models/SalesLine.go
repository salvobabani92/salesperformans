

package models

import (
	"github.com/salvobabani92/salesperformans.com/config"
	"time"
)


// Satış Sırası Tablosu.
// swagger: model
type SalesLine struct {
	// ID
	ID        uint `gorm:"primary_key" json:"id"`
	// Oluşturma Tarihi
	CreatedAt time.Time `json:"-"`
	// Son Değişiklik tarihi
	UpdatedAt time.Time `json:"-"`
	// Dosya Numarası
	DocumentNo string `json:"document_no"`
	// Sıra Numarası
	LineNo string `json:"line_no"`
	// Ürün kimiği
	ItemName  string `json:"item_name"`
	//Miktarı
	Quantity string `json:"quantity"`
	//Ürün Fiyatı
	UnitPrice string `json:"unit_price"`
	// Yüzdelik
	Amount int64 `json:"amount"`
	//Vergisi
	VAT int64 `json:"vat"`
	// Vergi Yüzdesi
	VATAmount int64 `json:"vat_amount"`
	ItemID uint `json:"item_id"`
	// Ürün İlişkisi
	Item Item `json:"_"`

}

func (this SalesLine) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}


// Satış Başlığı bilgileri geridönüş değeri
// swagger:response
type SalesLineResponse struct {
	// in:body
	Body SalesHeader
}

// Satış Başlığı bilgileri
// swagger:parameters Satış Başlığı-Oluştur
type SalesLineInputPatch struct {
	// Satış Başlığı Bilgileri
	// in:body
	SalesLine SalesLine
}

// Satış Başlığı Güncelle
// swagger:parameters Satış Başlığı-Güncelle
type SalesLineInputPUT struct {
	// Satış Başlığı Bilgileri
	// in:body
	SalesLine SalesLine
}


// Satış Başlığı Bilgisini Getir
// swagger:parameters Satış Başlığı-Getir
type SalesLineIDQueryGET struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Satış Başlığı Bilgisini Güncelle
// swagger:parameters Satış Başlığı-Güncelle
type SalesLineIDQueryPUT struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Satış Başlığı Bilgisini Sil
// swagger:parameters Satış Başlığı-Sil
type SalesLineIDQueryDELETE struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}

// Satış Başlığı bilgilerini excel'den yükle
// swagger:parameters Satış Başlığı-Excel-Yükle
type UploadSalesLineExcelInput struct {
	// Satış Başlığı bilgilerini excel'den yükle
	// in:body
	Body struct{ ExcelFile string `json:"ExcelFile"` }
}


// Satış Başlığı Yükle
// swagger:parameters Satış Başlığı-Json-Array
type SalesLineArrayInputPOST struct {
	// Satış Başlığı Bilgileri
	// in:body
	SalesLine [] SalesLine
}
