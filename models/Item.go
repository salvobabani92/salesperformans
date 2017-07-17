package models

import (
	"github.com/salvobabani92/salesperformans.com/config"
	"time"
)


// Lokasyon Tablosu.
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

}

func (this Location) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}


// Konum bilgileri geridönüş değeri
// swagger:response
type LocationResponse struct {
	// in:body
	Body Location
}

// Konum bilgileri
// swagger:parameters Konum-Oluştur
type LocationInputPatch struct {
	// Konum Bilgileri
	// in:body
	Location Location
}

// Konum Güncelle
// swagger:parameters Konum-Güncelle
type LocationInputPUT struct {
	// Konum Bilgileri
	// in:body
	Location Location
}


// Konum Bilgisini Getir
// swagger:parameters Konum-Getir
type LocationIDQueryGET struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Konum Bilgisini Güncelle
// swagger:parameters Konum-Güncelle
type KonumIDQueryPUT struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Konum Bilgisini Sil
// swagger:parameters Konum-Sil
type LocationIDQueryDELETE struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}

// Konum bilgilerini excel'den yükle
// swagger:parameters Konum-Excel-Yükle
type UploadLocationExcelInput struct {
	// Konum bilgilerini excel'den yükle
	// in:body
	Body struct{ ExcelFile string `json:"ExcelFile"` }
}


// Konum Yükle
// swagger:parameters Konum-Json-Array
type LocationArrayInputPOST struct {
	// Konum Bilgileri
	// in:body
	Location [] Location
