
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


// Satış Fiyatı oluştur
func POST_SalesPrice(c *gin.Context) {

	form := models.SalesPrice{}

	if val, hasValue := c.GetPostForm("type"); hasValue {
		form.Type = val
	}
	if val, hasValue := c.GetPostForm("Item_ID"); hasValue {
		uintVal, _ := strconv.ParseUint(val, 10, 64)
		form.ItemID = uint(uintVal)
	}

	if val, hasValue := c.GetPostForm("price"); hasValue {
		form.Price = val
	}

	if val, hasValue := c.GetPostForm("Customer_ID"); hasValue {
		uintVal, _ := strconv.ParseUint(val, 10, 64)
		form.CustomerID = uint(uintVal)
	}

	if config.DB.NewRecord(&form) {
		config.DB.Create(&form)
	}

	config.DB.NewRecord(form)
	log.Println("Yeni bir satış fiyatı eklendi.")
	c.JSON(http.StatusCreated, models.GetGenericStatusResponse("201", "Kaydınız başarı ile alınmıştır."))

}


// Satış Fiyatı Listesini getir
func GET_SalesPrice(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	// Get all matched records

	var SalesPrice  []models.SalesPrice
	config.DB.Where("customer_id = ?", user.CustomerID).Find(&SalesPrice)
	c.JSON(http.StatusOK, SalesPrice)
}

// Satış Fiyatı kaydını getir
func GET_SalesPriceByID(c *gin.Context) {
	log.Println("id'si bilinen bir Satış Fiyatı kaydını getir")
	user, _ := libs.GetUser_Company(c)

	var SalesPrice models.SalesPrice
	var id = c.Params.ByName("id")
	config.DB.Where("company_id = ? AND ID = ?", user.CustomerID, id).First(&SalesPrice)
	if SalesPrice.ID != 0 {
		c.JSON(http.StatusOK, SalesPrice)
	} else {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	}
}

// Satış Fiyatı güncelle
func PUT_SalesPrice(c *gin.Context) {
	//TODO: Organizasyon güncelle fonksiyonu yazılacak. CompanyID filtresini dikkate al
	user, _ := libs.GetUser_Company(c)
	form := models.SalesPrice{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {

		if val, hasValue := c.GetPostForm("type"); hasValue {
			form.Type = val
		}
		if val, hasValue := c.GetPostForm("Item_ID"); hasValue {
			uintVal, _ := strconv.ParseUint(val, 10, 64)
			form.ItemID = uint(uintVal)
		}

		if val, hasValue := c.GetPostForm("price"); hasValue {
			form.Price = val
		}

		if val, hasValue := c.GetPostForm("Customer_ID"); hasValue {
			uintVal, _ := strconv.ParseUint(val, 10, 64)
			form.CustomerID = uint(uintVal)
		}

		if config.DB.NewRecord(&form) {
			config.DB.Create(&form)
		}
		config.DB.Save(&form)
		c.JSON(http.StatusOK, models.GetGenericStatusResponse("200", "Kaydınız başarı ile alınmıştır."))
	}

}


// Satış Fiyatı Sil
func DELETE_SalesPrice(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	form := models.SalesPrice{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {
		config.DB.Delete(&form)
	}
}

// Satış Fiyatı Bilgilerini Excel ile içeri aktar
func Upload_SalesPrice_From_Excel(c *gin.Context) {
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

		var directoryName = "./upload/SalesPrice/" + strconv.FormatUint(uint64(user.CustomerID), 10) + "/"
		exist, _ := libs.FileOrDirectoryExists(directoryName)
		if exist == false {
			os.Mkdir(directoryName, 0700)
		}

		form := models.Upload{}
		form.UserID = user.ID

		form.FileExtension = extension
		form.RelatedTableName = "SalesPrice"

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
			if sheet.Name == "SalesPrice" {
				for _, curRow := range sheet.Rows {
					rowNumber ++
					if rowNumber > 1 {
						Type, _ := curRow.Cells[0].String()
						Price, _ := curRow.Cells[1].String()


						form := models.SalesPrice{}
						form.CustomerID = user.CustomerID
						form.Type = Type
						form.Price = Price

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

// Satış Fiyatı bilgilerini batchjob olarak aktarılması
func Upload_SalesPrice_From_Json_Array(c *gin.Context) {
	//TODO: Json içersinde array olarak içeri aktarılması ile ilgili fonksiyon yazılacak. Bu fonksiyon upsert gibi çalışacak. CompanyID filtresini dikkate al

}


