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


// Ürün oluştur
func POST_Item(c *gin.Context) {

	user, _ := libs.GetUser_Company(c)

	form := models.Item{}

	if val, hasValue := c.GetPostForm("no"); hasValue {
		form.No = val
	}

	if val, hasValue := c.GetPostForm("description"); hasValue {
		form.Description = val
	}

	if val, hasValue := c.GetPostForm("barcode_no"); hasValue {
		form.BarcodeNo = val
	}

	if val, hasValue := c.GetPostForm("description"); hasValue {
		form.Description = val
	}

	if val, hasValue := c.GetPostForm("unit_of_measure"); hasValue {
		form.UnitofMeasure = val
	}

	if config.DB.NewRecord(&form) {
		config.DB.Create(&form)
	}

	config.DB.NewRecord(form)
	log.Println("Yeni bir ürün eklendi.")
	c.JSON(http.StatusCreated, models.GetGenericStatusResponse("201", "Kaydınız başarı ile alınmıştır."))

}


// Ürün Listesini getir
func GET_Item(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	// Get all matched records

	var Item  []models.Item
	config.DB.Where("customer_id = ?", user.CustomerID).Find(&Item)
	c.JSON(http.StatusOK, Item)
}

// Ürün kaydını getir
func GET_ItemByID(c *gin.Context) {
	log.Println("id'si bilinen bir item kaydını getir")
	user, _ := libs.GetUser_Company(c)

	var Item models.Item
	var id = c.Params.ByName("id")
	config.DB.Where("company_id = ? AND ID = ?", user.CustomerID, id).First(&Item)
	if Item.ID != 0 {
		c.JSON(http.StatusOK, Item)
	} else {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	}
}

// Ürün güncelle
func PUT_Item(c *gin.Context) {
	//TODO: Organizasyon güncelle fonksiyonu yazılacak. CompanyID filtresini dikkate al
	user, _ := libs.GetUser_Company(c)
	form := models.Item{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {

		if val, hasValue := c.GetPostForm("no"); hasValue {
			form.No = val
		}

		if val, hasValue := c.GetPostForm("description"); hasValue {
			form.Description = val
		}

		if val, hasValue := c.GetPostForm("barcode_no"); hasValue {
			form.BarcodeNo = val
		}

		if val, hasValue := c.GetPostForm("description"); hasValue {
			form.Description = val
		}

		if val, hasValue := c.GetPostForm("unit_of_measure"); hasValue {
			form.UnitofMeasure = val
		}

		if config.DB.NewRecord(&form) {
			config.DB.Create(&form)
		}
		config.DB.Save(&form)
		c.JSON(http.StatusOK, models.GetGenericStatusResponse("200", "Kaydınız başarı ile alınmıştır."))
	}

}


// Ürün Sil
func DELETE_Item(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	form := models.Item{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {
		config.DB.Delete(&form)
	}
}

// Ürün Bilgilerini Excel ile içeri aktar
func Upload_Item_From_Excel(c *gin.Context) {
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

		var directoryName = "./upload/Item/" + strconv.FormatUint(uint64(user.CustomerID), 10) + "/"
		exist, _ := libs.FileOrDirectoryExists(directoryName)
		if exist == false {
			os.Mkdir(directoryName, 0700)
		}

		form := models.Upload{}
		form.Customer = user.CustomerID
		form.UserID = user.ID

		form.FileExtension = extension
		form.RelatedTableName = "Item"

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
			if sheet.Name == "Item" {
				for _, curRow := range sheet.Rows {
					rowNumber ++
					if rowNumber > 1 {
						No, _ := curRow.Cells[0].String()
						Description, _ := curRow.Cells[1].String()
						BarcodeNo, _ := curRow.Cells[2].String()
						UnitofMeasure, _ := curRow.Cells[3].String()
						form := models.Item{}
						form.No = No
						form.Description = Description
						form.BarcodeNo = BarcodeNo
						form.UnitofMeasure = UnitofMeasure

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

// Ürün bilgilerini batchjob olarak aktarılması
func Upload_Item_From_Json_Array(c *gin.Context) {
	//TODO: Json içersinde array olarak içeri aktarılması ile ilgili fonksiyon yazılacak. Bu fonksiyon upsert gibi çalışacak. CompanyID filtresini dikkate al

}

