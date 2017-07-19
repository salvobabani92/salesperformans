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


// Satış Başlığı oluştur
func POST_SalesHeader(c *gin.Context) {

	user, _ := libs.GetUser_Company(c)

	form := models.SalesHeader{}
	form.Customer = user.CustomerID

	if val, hasValue := c.GetPostForm("no"); hasValue {
		form.No = val
	}

	if val, hasValue := c.GetPostForm("amount"); hasValue {
		int64Val, _ := strconv.ParseUint(val, 10, 64)
		form.Amount = int64(int64Val)
	}

	if val, hasValue := c.GetPostForm("unit_of_measure_ID"); hasValue {
		uintVal, _ := strconv.ParseUint(val, 10, 64)
		form.UnitofMeasureID = uint(uintVal)
	}

	if val, hasValue := c.GetPostForm("unit_of_measure_code"); hasValue {
		form.UnitofMeasureCode = val
	}

	if config.DB.NewRecord(&form) {
		config.DB.Create(&form)
	}

	config.DB.NewRecord(form)
	log.Println("Yeni bir satış başlığı eklendi.")
	c.JSON(http.StatusCreated, models.GetGenericStatusResponse("201", "Kaydınız başarı ile alınmıştır."))

}


// Satış Başlığı Listesini getir
func GET_SalesHeader(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	// Get all matched records

	var SalesHeader  []models.SalesHeader
	config.DB.Where("customer_id = ?", user.CustomerID).Find(&SalesHeader)
	c.JSON(http.StatusOK, Item)
}

// Satış Başlığı kaydını getir
func GET_SalesHeaderByID(c *gin.Context) {
	log.Println("id'si bilinen bir satış başlığı kaydını getir")
	user, _ := libs.GetUser_Company(c)

	var SalesHeader models.SalesHeader
	var id = c.Params.ByName("id")
	config.DB.Where("company_id = ? AND ID = ?", user.CustomerID, id).First(&SalesHeader)
	if SalesHeader.ID != 0 {
		c.JSON(http.StatusOK, SalesHeader)
	} else {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	}
}

// Satış Başlığı güncelle
func PUT_SalesHeader(c *gin.Context) {
	//TODO: Organizasyon güncelle fonksiyonu yazılacak. CompanyID filtresini dikkate al
	user, _ := libs.GetUser_Company(c)
	form := models.SalesHeader{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {

		if val, hasValue := c.GetPostForm("no"); hasValue {
			form.No = val
		}

		if val, hasValue := c.GetPostForm("amount"); hasValue {
			int64Val, _ := strconv.ParseUint(val, 10, 64)
			form.Amount = int64(int64Val)
		}

		if val, hasValue := c.GetPostForm("unit_of_measure_ID"); hasValue {
			uintVal, _ := strconv.ParseUint(val, 10, 64)
			form.UnitofMeasureID = uint(uintVal)
		}

		if val, hasValue := c.GetPostForm("unit_of_measure_code"); hasValue {
			form.UnitofMeasureCode = val
		}

		if config.DB.NewRecord(&form) {
			config.DB.Create(&form)
		}
		config.DB.Save(&form)
		c.JSON(http.StatusOK, models.GetGenericStatusResponse("200", "Kaydınız başarı ile alınmıştır."))
	}

}


// Satış Başlığı Sil
func DELETE_SalesHeader(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	form := models.SalesHeader{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {
		config.DB.Delete(&form)
	}
}

// Satış Başlığı Bilgilerini Excel ile içeri aktar
func Upload_SalesHeader_From_Excel(c *gin.Context) {
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

		var directoryName = "./upload/SalesHeader/" + strconv.FormatUint(uint64(user.CustomerID), 10) + "/"
		exist, _ := libs.FileOrDirectoryExists(directoryName)
		if exist == false {
			os.Mkdir(directoryName, 0700)
		}

		form := models.Upload{}
		form.Customer = user.CustomerID
		form.UserID = user.ID

		form.FileExtension = extension
		form.RelatedTableName = "SalesHeader"

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
			if sheet.Name == "SalesHeader" {
				for _, curRow := range sheet.Rows {
					rowNumber ++
					if rowNumber > 1 {
						No, _ := curRow.Cells[0].String()

						sVal, _ := curRow.Cells[1].String()
						idVal, _ := strconv.ParseUint(sVal, 10, 64)
						Amount := int64(idVal)

						UnitofMeasureCode, _ := curRow.Cells[2].String()

						form := models.SalesHeader{}
						form.CustomerID = user.CustomerID
						form.No = No
						form.Amount = Amount
						form.UnitofMeasureCode = UnitofMeasureCode


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

// Satış Başlığı bilgilerini batchjob olarak aktarılması
func Upload_SalesHeader_From_Json_Array(c *gin.Context) {
	//TODO: Json içersinde array olarak içeri aktarılması ile ilgili fonksiyon yazılacak. Bu fonksiyon upsert gibi çalışacak. CompanyID filtresini dikkate al

}

