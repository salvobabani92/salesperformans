package apicontrollers

import (
	"github.com/gin-gonic/gin"
	"github.com/salvobabani92/salesperformans.com/config"
	"github.com/salvobabani92/salesperformans.com/models"
	"github.com/salvobabani92/salesperformans.com/libs"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"strconv"
	"os"
	"io"
	"github.com/tealeg/xlsx"

)


// Ürün oluştur
func POST_Customer(c *gin.Context) {

	user, _ := libs.GetUser_Company(c)

	form := models.Customer{}
	form.Customer = user.Customer

	if val, hasValue := c.GetPostForm("no"); hasValue {
		form.No = val
	}

	if val, hasValue := c.GetPostForm("name"); hasValue {
		form.Name = val
	}


	if config.DB.NewRecord(&form) {
		config.DB.Create(&form)
	}

	config.DB.NewRecord(form)
	log.Println("Yeni bir müşteri eklendi.")
	c.JSON(http.StatusCreated, models.GetGenericStatusResponse("201", "Kaydınız başarı ile alınmıştır."))

}


// Müşteri Listesini getir
func GET_Customer(c *gin.Context) {
	user, _ := libs.GetUser_Company(c)
	// Get all matched records

	var Customer  []models.Customer
	config.DB.Where("customer_id = ?", user.CustomerID).Find(&Customer)
	c.JSON(http.StatusOK, Customer)
}

// Müşteri kaydını getir
func GET_CustomerByID(c *gin.Context) {
	log.Println("id'si bilinen bir customer kaydını getir")
	user, _ := libs.GetUser_Customer(c)

	var Customer models.Customer
	var id = c.Params.ByName("id")
	config.DB.Where("company_id = ? AND ID = ?", user.CompanyID, id).First(&Customer)
	if Customer.ID != 0 {
		c.JSON(http.StatusOK, Customer)
	} else {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	}
}

// Ürün güncelle
func PUT_Customer(c *gin.Context) {
	//TODO: Organizasyon güncelle fonksiyonu yazılacak. CompanyID filtresini dikkate al
	user, _ := libs.GetUser_Customer(c)
	form := models.Customer{}
	var id = c.Params.ByName("id")
	config.DB.Where("customer_id = ? AND ID = ?", user.CustomerID, id).First(&form)
	if form.ID == 0 {
		c.JSON(http.StatusNotFound, models.GetGenericStatusResponse("404", "Kayıt bulunamadı."))
	} else {

		if val, hasValue := c.GetPostForm("no"); hasValue {
			form.No = val
		}

		if val, hasValue := c.GetPostForm("name"); hasValue {
			form.Name = val
		}


		if config.DB.NewRecord(&form) {
			config.DB.Create(&form)
		}
		config.DB.Save(&form)
		c.JSON(http.StatusOK, models.GetGenericStatusResponse("200", "Kaydınız başarı ile alınmıştır."))
	}

}


// Ürün Sil
func DELETE_Customer(c *gin.Context) {
	user, _ := libs.GetUser_Customer(c)
	form := models.Customer{}
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


	user, _ := libs.GetUser_Customer(c)

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

		var directoryName = "./upload/Customer/" + strconv.FormatUint(uint64(user.customerID), 10) + "/"
		exist, _ := libs.FileOrDirectoryExists(directoryName)
		if exist == false {
			os.Mkdir(directoryName, 0700)
		}

		form := models.Upload{}
		form.Customer = user.CustomerID
		form.UserID = user.ID

		form.FileExtension = extension
		form.RelatedTableName = "Customer"

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
			if sheet.Name == "Customer" {
				for _, curRow := range sheet.Rows {
					rowNumber ++
					if rowNumber > 1 {
						No, _ := curRow.Cells[0].String()
						Name, _:= curRow.Cells[1].String()
						form := models.Customer{}
						form.CustomerID = user.CustomerID
						form.No = no
						form.Name = name


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

// Konum bilgilerini batchjob olarak aktarılması
func Upload_Location_From_Json_Array(c *gin.Context) {
	//TODO: Json içersinde array olarak içeri aktarılması ile ilgili fonksiyon yazılacak. Bu fonksiyon upsert gibi çalışacak. CompanyID filtresini dikkate al

}

