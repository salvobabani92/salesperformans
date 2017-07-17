

package models

import (
	"github.com/salvobabani92/salesperformans.com/config"
	"time"
)


// Satış Fiyatı Tablosu.
// swagger: model
type SalesPrice struct {
	// ID
	ID        uint `gorm:"primary_key" json:"id"`
	// Tipi
	Type  string `json:"type"`
	// Oluşturma Tarihi
	CreatedAt time.Time `json:"-"`
	// Son Değişiklik tarihi
	UpdatedAt time.Time `json:"-"`
	ItemID uint `json:"item_id"`
	Price string `json:"price"`
	CustomerID uint `json:"customer_id"`
	Customer Customer `json:"_"`
	Item Item `json:"_"`

}

func (this SalesPrice) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}


// Satış Fiyatı bilgileri geridönüş değeri
// swagger:response
type SalesPriceResponse struct {
	// in:body
	Body SalesHeader
}

// Satış Fiyatı bilgileri
// swagger:parameters Satış Fiyatı-Oluştur
type SalesPriceInputPatch struct {
	// Satış Fiyatı Bilgileri
	// in:body
	SalesHeader SalesHeader
}

// Satış Fiyatı Güncelle
// swagger:parameters Satış Fiyatı-Güncelle
type SalesPriceInputPUT struct {
	// Satış Fiyatı Bilgileri
	// in:body
	SalesPrice SalesPrice
}


// Satış Fiyatı Bilgisini Getir
// swagger:parameters Satış Fiyatı-Getir
type SalesPriceIDQueryGET struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Satış Fiyatı Bilgisini Güncelle
// swagger:parameters Satış Fiyatı-Güncelle
type SalesPriceIDQueryPUT struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Satış Fiyatı Bilgisini Sil
// swagger:parameters Satış Fiyatı-Sil
type SalesPriceIDQueryDELETE struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}

// Satış Fiyatı bilgilerini excel'den yükle
// swagger:parameters Satış Fiyatı-Excel-Yükle
type UploadSalesPriceExcelInput struct {
	// Satış Fiyatı bilgilerini excel'den yükle
	// in:body
	Body struct{ ExcelFile string `json:"ExcelFile"` }
}


// Satış Fiyatı Yükle
// swagger:parameters Satış Fiyatı-Json-Array
type SalesPriceArrayInputPOST struct {
	// Satış Fiyatı Bilgileri
	// in:body
	SalesPrice [] SalesPrice
}
