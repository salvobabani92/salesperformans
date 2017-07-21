package apicontrollers

import (
	"github.com/gin-gonic/gin"
	"github.com/salvobabani92/salesperformans.com/config"
	"github.com/salvobabani92/salesperformans.com/models"
	"github.com/salvobabani92/salesperformans.com/libs"
	"net/http"
	"log"
	"path/filepath"
	"os"
	"io"
	"strconv"
	"strings"
	"github.com/tealeg/xlsx"
)


// Satış Sırası oluştur
func POST_SalesLine(c *gin.Context) {

	form := models.SalesLine{}


	if val, hasValue := c.GetPostForm("document_no"); hasValue {
		form.DocumentNo = val
	}

	if val, hasValue := c.GetPostForm("line_no"); hasValue {
		form.LineNo = val
	}

	if val, hasValue := c.GetPostForm("Item_ID"); hasValue {
		uintVal, _ := strconv.ParseUint(val, 10, 64)
		form.ItemID = uint(uintVal)
	}

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
	log.Println("Yeni bir satış başlığı eklendi.")
	c.JSON(http.StatusCreated, models.GetGenericStatusResponse("201", "Kaydınız başarı ile alınmıştır."))

}


// Satış Sırası Listesini getir
func GET_SalesLine(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	// Get all matched records

	var SalesLine  []models.SalesLine
	config.DB.Where("customer_id = ?", user.CustomerID).Find(&SalesLine)
	c.JSON(http.StatusOK, SalesLine)
}

// Satış Sırası kaydını getir
func GET_SalesLineByID(c *gin.Context) {
	log.Println("id'si bilinen bir satış sırası kaydını getir")
	user, _ := libs.GetUser_Company(c)

	var SalesLine models.SalesLine
	var id = c.Params.ByName("id")
	config.DB.Where("company_id = ? AND ID = ?", user.CustomerID, id).First(&SalesLine)
	if SalesLine.ID != 0 {
		c.JSON(http.StatusOK, SalesLine)
	} else {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	}
}

// Satış Sırası güncelle
func PUT_SalesLine(c *gin.Context) {
	//TODO: Organizasyon güncelle fonksiyonu yazılacak. CompanyID filtresini dikkate al
	user, _ := libs.GetUser_Company(c)
	form := models.SalesLine{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {

		if val, hasValue := c.GetPostForm("document_no"); hasValue {
			form.DocumentNo = val
		}

		if val, hasValue := c.GetPostForm("line_no"); hasValue {
			form.LineNo = val
		}

		if val, hasValue := c.GetPostForm("Item_ID"); hasValue {
			uintVal, _ := strconv.ParseUint(val, 10, 64)
			form.ItemID = uint(uintVal)
		}

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


// Satış Sırası Sil
func DELETE_SalesLine(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	form := models.SalesLine{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {
		config.DB.Delete(&form)
	}
}

// Satış Sırası Bilgilerini Excel ile içeri aktar
func Upload_SalesLine_From_Excel(c *gin.Context) {
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

		var directoryName = "./upload/SalesLine/" + strconv.FormatUint(uint64(user.CustomerID), 10) + "/"
		exist, _ := libs.FileOrDirectoryExists(directoryName)
		if exist == false {
			os.Mkdir(directoryName, 0700)
		}

		form := models.Upload{}
		form.UserID = user.ID

		form.FileExtension = extension
		form.RelatedTableName = "SalesLine"

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
			if sheet.Name == "SalesLine" {
				for _, curRow := range sheet.Rows {
					rowNumber ++
					if rowNumber > 1 {
						DocumentNo, _ := curRow.Cells[0].String()
						LineNo, _ := curRow.Cells[1].String()

						ItemName, _ := curRow.Cells[2].String()
						Quantity, _ := curRow.Cells[3].String()
						UnitPrice, _ := curRow.Cells[4].String()

						sVal, _ := curRow.Cells[5].String()
						idVal, _ := strconv.ParseUint(sVal, 10, 64)
						Amount := int64(idVal)

						sVal, _ = curRow.Cells[6].String()
						idVal, _ = strconv.ParseUint(sVal, 10, 64)
						VAT := int64(idVal)

						sVal, _ = curRow.Cells[7].String()
						idVal, _ = strconv.ParseUint(sVal, 10, 64)
						VATAmount := int64(idVal)


						form := models.SalesLine{}
						form.ItemName = ItemName
						form.Quantity = Quantity
						form.UnitPrice = UnitPrice
						form.Amount = Amount
						form.VAT = VAT
						form.VATAmount = VATAmount
						form.DocumentNo = DocumentNo
						form.LineNo = LineNo

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

// Satış Sırası bilgilerini batchjob olarak aktarılması
func Upload_SalesLine_From_Json_Array(c *gin.Context) {
	//TODO: Json içersinde array olarak içeri aktarılması ile ilgili fonksiyon yazılacak. Bu fonksiyon upsert gibi çalışacak. CompanyID filtresini dikkate al

}


