
package apicontrollers

import (
	"github.com/gin-gonic/gin"
	"github.com/salvobabani92/salesperformans.com/config"
	"github.com/salvobabani92/salesperformans.com/models"
	"strconv"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"os"
	"io"
	"github.com/tealeg/xlsx"

)



// Birim Ölçüsü oluştur
func POST_UnitofMeasure(c *gin.Context) {

	user, _ := libs.GetUser_Company(c)

	form := models.UnitofMeasure{}
	form.Customer = user.CustomerID

	if val, hasValue := c.GetPostForm("Item_name"); hasValue {
		form.ItemName = val
	}

	if val, hasValue := c.GetPostForm("quantity"); hasValue {
		form.Quantity = val
	}

	if val, hasValue := c.GetPostForm("unit_price"); hasValue {
		form.UnitPrice = val
	}

	if val, hasValue := c.GetPostForm("amount"); hasValue {
		int64Val, _ := strconv.ParseUint(val, 10, 64)
		form.Amount = int64(int64Val)
	}

	if val, hasValue := c.GetPostForm("VAT"); hasValue {
		int64Val, _ := strconv.ParseUint(val, 10, 64)
		form.VAT = int64(int64Val)
	}

	if val, hasValue := c.GetPostForm("VAT_amount"); hasValue {
		int64Val, _ := strconv.ParseUint(val, 10, 64)
		form.VATAmount = int64(int64Val)
	}


	if config.DB.NewRecord(&form) {
		config.DB.Create(&form)
	}

	config.DB.NewRecord(form)
	log.Println("Yeni bir satış fiyatı eklendi.")
	c.JSON(http.StatusCreated, models.GetGenericStatusResponse("201", "Kaydınız başarı ile alınmıştır."))

}


// Birim Ölçüsü Listesini getir
func GET_UnitofMeasure(c *gin.Context) {
	user, _ := libs.GetUser_Customer(c)
	// Get all matched records

	var UnitofMeasure  []models.UnitofMeasure
	config.DB.Where("customer_id = ?", user.CustomerID).Find(&UnitofMeasure)
	c.JSON(http.StatusOK, UnitofMeasure)
}

// Birim Ölçüsü kaydını getir
func GET_UnitofMeasureByID(c *gin.Context) {
	log.Println("id'si bilinen bir Satış Fiyatı kaydını getir")
	user, _ := libs.GetUser_Company(c)

	var UnitofMeasure models.UnitofMeasure
	var id = c.Params.ByName("id")
	config.DB.Where("company_id = ? AND ID = ?", user.CustomerID, id).First(&UnitofMeasure)
	if UnitofMeasure.ID != 0 {
		c.JSON(http.StatusOK, UnitofMeasure)
	} else {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	}
}

// Birim Ölçüsü güncelle
func PUT_UnitofMeasure(c *gin.Context) {
	//TODO: Organizasyon güncelle fonksiyonu yazılacak. CompanyID filtresini dikkate al
	user, _ := libs.GetUser_Company(c)
	form := models.UnitofMeasure{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {

		if val, hasValue := c.GetPostForm("Item_name"); hasValue {
			form.ItemName = val
		}

		if val, hasValue := c.GetPostForm("quantity"); hasValue {
			form.Quantity = val
		}

		if val, hasValue := c.GetPostForm("unit_price"); hasValue {
			form.UnitPrice = val
		}

		if val, hasValue := c.GetPostForm("amount"); hasValue {
			int64Val, _ := strconv.ParseUint(val, 10, 64)
			form.Amount = int64(int64Val)
		}

		if val, hasValue := c.GetPostForm("VAT"); hasValue {
			int64Val, _ := strconv.ParseUint(val, 10, 64)
			form.VAT = int64(int64Val)
		}

		if val, hasValue := c.GetPostForm("VAT_amount"); hasValue {
			int64Val, _ := strconv.ParseUint(val, 10, 64)
			form.VATAmount = int64(int64Val)
		}

		if config.DB.NewRecord(&form) {
			config.DB.Create(&form)
		}
		config.DB.Save(&form)
		c.JSON(http.StatusOK, models.GetGenericStatusResponse("200", "Kaydınız başarı ile alınmıştır."))
	}

}


// Birim Ölçüsü Sil
func DELETE_UnitofMeasure(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	form := models.UnitofMeasure{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {
		config.DB.Delete(&form)
	}
}

// Birim Ölçüsü Bilgilerini Excel ile içeri aktar
func Upload_UnitofMeasure_From_Excel(c *gin.Context) {
	//TODO: Excel'den bu bilgilerin aktarılması ile ilgili fonksiyon yazılacak. CompanyID filtresini dikkate al


	user, _ := libs.GetUser_Company(c)

	if (user.CustomerID == 0) {
		c.JSON(http.StatusBadRequest, models.GetGenericStatusResponse("400", "Şirket bilgileriniz yanlış."))

		return
	}

	file, header, _ := c.Request.FormFile("file")
	log.Println(file)
	filename := header.Filename
	var extension = filepath.Ext(filename)

	if (!strings.HasPrefix(extension, ".xls")) {
		c.JSON(http.StatusBadRequest, models.GetGenericStatusResponse("400", "Yüklediğiniz dosya excel dosyası değil."))
	} else {

		var directoryName = "./upload/UnitofMeasure/" + strconv.FormatUint(uint64(user.customerID), 10) + "/"
		exist, _ := libs.FileOrDirectoryExists(directoryName)
		if exist == false {
			os.Mkdir(directoryName, 0700)
		}

		form := models.Upload{}
		form.Customer = user.CustomerID
		form.UserID = user.ID

		form.FileExtension = extension
		form.RelatedTableName = "UnitofMeasure"

		if config.DB.NewRecord(&form) {
			config.DB.Create(&form)
		}
		config.DB.NewRecord(&form)
		form.FilePath = directoryName + strconv.FormatUint(uint64(form.ID), 16) + extension

		out, err := os.Create(form.FilePath)

		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}

		xlFile, err := xlsx.OpenFile(form.FilePath)
		if err != nil {
			log.Println(err)
		}

		var rowNumber int = 0;


		for _, sheet := range xlFile.Sheets {
			if sheet.Name == "UnitofMeasure" {
				for _, curRow := range sheet.Rows {
					rowNumber ++
					if rowNumber > 1 {
						ItemName, _ := curRow.Cells[0].String()
						Quantity, _ := curRow.Cells[1].String()
						UnitPrice, _ := curRow.Cells[2].String()

						sVal, _ := curRow.Cells[3].String()
						idVal, _ := strconv.ParseUint(sVal, 10, 64)
						Amount := int64(idVal)

						sVal, _ = curRow.Cells[4].String()
						idVal, _ = strconv.ParseUint(sVal, 10, 64)
						VAT := int64(idVal)


						sVal, _ = curRow.Cells[5].String()
						idVal, _ = strconv.ParseUint(sVal, 10, 64)
						VATAmount := int64(idVal)


						form := models.UnitofMeasure{}
						form.CustomerID = user.CustomerID
						form.ItemName = ItemName
						form.Quantity = Quantity
						form.UnitPrice = UnitPrice
						form.Amount = Amount
						form.VAT = VAT
						form.VATAmount = VATAmount



						if config.DB.NewRecord(&form) {
							config.DB.Create(&form)
						}
						config.DB.NewRecord(form)
					}
				}
			}
		}

		c.JSON(http.StatusOK, models.GetGenericStatusResponse("200", "Dosya yüklendi."))

	}
}

// Birim Ölçüsü bilgilerini batchjob olarak aktarılması
func Upload_UnitofMeasure_From_Json_Array(c *gin.Context) {
	//TODO: Json içersinde array olarak içeri aktarılması ile ilgili fonksiyon yazılacak. Bu fonksiyon upsert gibi çalışacak. CompanyID filtresini dikkate al

}



