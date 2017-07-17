

package models

import (
	"github.com/salvobabani92/salesperformans.com/config"
	"time"
)


// Birim Ölçüsü Tablosu.
// swagger: model
type UnitofMeasure struct {
	// ID
	ID        uint `gorm:"primary_key" json:"id"`
	// Oluşturma Tarihi
	CreatedAt time.Time `json:"-"`
	// Son Değişiklik tarihi
	UpdatedAt time.Time `json:"-"`
	// Ürün adı
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

}

func (this UnitofMeasure) CreateTable() {
	config.DB.DropTable(this)
	config.DB.CreateTable(this)
}


// Birim Ölçüsü bilgileri geridönüş değeri
// swagger:response
type UnitofMeasureResponse struct {
	// in:body
	Body UnitofMeasure
}

// Birim Ölçüsü bilgileri
// swagger:parameters Birim Ölçüsü-Oluştur
type UnitofMeasureInputPatch struct {
	// Birim Ölçüsü Bilgileri
	// in:body
	UnitofMeasure UnitofMeasure
}

// Birim Ölçüsü Güncelle
// swagger:parameters Birim Ölçüsü-Güncelle
type UnitofMeasureInputPUT struct {
	// Birim Ölçüsü Bilgileri
	// in:body
	UnitofMeasure UnitofMeasure
}


// Birim Ölçüsü Bilgisini Getir
// swagger:parameters Birim Ölçüsü-Getir
type UnitofMeasureIDQueryGET struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Birim Ölçüsü Bilgisini Güncelle
// swagger:parameters Birim Ölçüsü-Güncelle
type UnitofMeasureIDQueryPUT struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}



// Birim Ölçüsü Bilgisini Sil
// swagger:parameters Birim Ölçüsü-Sil
type UnitofMeasureIDQueryDELETE struct {
	// in:query
	// type:int
	ID uint `json:"id"`
}

// Birim Ölçüsü bilgilerini excel'den yükle
// swagger:parameters Birim Ölçüsü-Excel-Yükle
type UploadUnitofMeasureExcelInput struct {
	// Birim Ölçüsü bilgilerini excel'den yükle
	// in:body
	Body struct{ ExcelFile string `json:"ExcelFile"` }
}


// Birim Ölçüsü Yükle
// swagger:parameters Birim Ölçüsü-Json-Array
type UnitofMeasureArrayInputPOST struct {
	// Birim Ölçüsü Bilgileri
	// in:body
	UnitofMeasure [] UnitofMeasure
}
