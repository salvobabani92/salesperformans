package models

import (
	"github.com/salvobabani92/salesperformans.com/config"
	"time"
)


// Item Tablosu.
// swagger: model
type Item struct {
	// ID
	ID        uint `gorm:"primary_key" json:"id"`
	// Oluşturma Tarihi
	CreatedAt time.Time `json:"-"`
	// Son Değişiklik tarihi
	UpdatedAt time.Time `json:"-"`
	// Ürün Numarası
	No     string `json:"no"`
	//Açıklaması
	Description  string `json:"description"`
	//Birim Şekli
	UnitofMeasure  string `json:"unitof_measure"`
	//Barkod Numarası
	BarcodeNo   string `json:"barcode_no"`

}

func (this Item) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}


// Ürün bilgileri geridönüş değeri
// swagger:response
type ItemResponse struct {
	// in:body
	Body Item
}

// Ürün bilgileri
// swagger:parameters Ürün-Oluştur
type ItemInputPatch struct {
	// Ürün Bilgileri
	// in:body
	Item Item
}

// Ürün Güncelle
// swagger:parameters Ürün-Güncelle
type ItemInputPUT struct {
	// Konum Bilgileri
	// in:body
	Item Item
}


// Ürün Bilgisini Getir
// swagger:parameters Ürün-Getir
type ItemIDQueryGET struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Ürün Bilgisini Güncelle
// swagger:parameters Ürün-Güncelle
type ItemIDQueryPUT struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Ürün Bilgisini Sil
// swagger:parameters Ürün-Sil
type ItemIDQueryDELETE struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}

// Ürün bilgilerini excel'den yükle
// swagger:parameters Ürün-Excel-Yükle
type UploadItemExcelInput struct {
	// Ürün bilgilerini excel'den yükle
	// in:body
	Body struct{ ExcelFile string `json:"ExcelFile"` }
}


// Ürün Yükle
// swagger:parameters Ürün-Json-Array
type ItemArrayInputPOST struct {
	// Ürün Bilgileri
	// in:body
	Item [] Item
}