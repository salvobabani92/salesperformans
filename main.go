

package main

import (
	"log"
	"os"
	"html/template"
	"github.com/gin-gonic/gin"

	"github.com/salvobabani92/salesperformans.com/config"
	"github.com/salvobabani92/salesperformans.com/controller"
	"github.com/salvobabani92/salesperformans.com/models"
	"github.com/salvobabani92/salesperformans.com/libs"

	"github.com/salvobabani92/salesperformans.com/apicontrollers"
)

//go:generate swagger generate spec
func main() {
	port := os.Getenv("PORT")
	log.Println("Connecting DB")

	config.InitDB()
	models.User{}.CreateTable()
	models.Customer{}.CreateTable()
	models.Item{}.CreateTable()
	models.SalesHeader{}.CreateTable()
	models.SalesLine{}.CreateTable()
	models.SalesPrice{}.CreateTable()
	models.Upload{}.CreateTable()

	if port == "" {
		port = "8000"
	}

	app := gin.Default()

	app.SetHTMLTemplate(template.Must(template.ParseFiles(
		"templates/header.html",
		"templates/index.html",
		"templates/footer.html",
		"templates/ReDoc.html",
	)))

	// Static Files
	app.Static("/images", "./images")
	//app.StaticFile("./swagger.json","./swagger.json")


	// Home Page
	app.GET("/", controller.Index)


	// Redoc Documentation
	app.GET("/doc", controller.Redoc)


	/*
		-------------------------------------------
		-------- Register & Login Pages -----------
		-------------------------------------------
		*/




	// swagger:route POST /register Kullanıcılar register
	// Yeni bir kullanıcı ekle
	//
	// Sisteme yeni bir kullanıcı eklemek için kullanır ve geriye kullanıcı bilgilerini gönderir.
	//
	//
	// responses:
	//   201: StatusResponse201
	app.POST("/register", controller.Register)





	// swagger:route POST /login Kullanıcılar Kullanıcı-Girişi
	// Kullanıcı Girişi
	//
	// Kullanıcı girişi sonrası bilet kullanılacaktır..
	//
	// .
	// responses:
	//   200: LoginResponce
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	app.POST("/login", controller.Login)






	// api roote group
	api := app.Group("/api", libs.BiletKontrol)




	// swagger:route GET /api/ping Genel send-ping
	//
	// Ping isteği gönder
	//
	// Uygulamanın ayakta olup olmadığını basitçe anlamaya
	// yarayan servis. Geriye pong cevabını gönderir.
	//
	// responses:
	//   200: PongResponse
	//     description: OK
	api.GET("/ping", controller.Ping)





	// swagger:route GET /api/me Kullanıcılar Benim-Bilgilerim
	// Aktif Kullanıcı Bilgileri
	//
	// Sisteme bağlı olan kullanıcıların bilgilerini gösterir
	//
	// .
	// responses:
	//   200: UserResponce
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.GET("/me", apicontrollers.GetMe)




	/*
	-------------------------------------------
	-------- Customer								  -----------
	-------------------------------------------
	*/




	// swagger:route POST /api/Customer Müşteri-Oluştur
	// Ekle
	//
	// Yeni bir Müşteri kaydı oluşturur
	//
	// responses:
	//   201: StatusResponse201
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/Customer", apicontrollers.POST_Customer)





	// swagger:route GET /api/customer Müşteri-Listesini-Getir
	// Liste
	//
	// Müşteri Listesini getirir
	//
	// responses:
	//   200: CustomerResponse
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.GET("/Customer",apicontrollers.GET_Customer)



	// swagger:route GET /api/customer/:id Müşteri-Getir
	// Oku.
	//
	// ID paramretresi ile belirtilen Müşteri kaydını getirir.
	//
	// responses:
	//   200: CustomerResponse
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	//   404: ErrorResponse404
	api.GET("/customer/:id",apicontrollers.GET_CustomerByID)




	// swagger:route PUT /api/customer/:id  Müşteri-Güncelle
	// Güncelle
	//
	// Müşteri bilgisini günceller
	//
	// responses:
	//   200: StatusResponse200
	//   404: ErrorResponse404
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.PUT("/customer/:id",apicontrollers.PUT_Customer)



	// swagger:route DELETE /api/customer/:id Müşteri-Sil
	// Sil
	//
	// Müşteri bilgisini sil
	//
	// responses:
	//   200: StatusResponse200
	//   404: ErrorResponse404
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.DELETE("/customer/:id",apicontrollers.DELETE_Customer)



	// swagger:route POST /api/customer/upload/excel Müşteri-Excel-Yükle
	// Excel Yükle
	//
	// Aktarım şablonu ile excel üzerinden Müşteri bilgilerinin yüklenmesi için kullanılır.
	//
	// .
	// responses:
	//   200: StatusResponse200
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/customer/upload/excel",apicontrollers.Upload_Customer_From_Excel)



	// swagger:route POST /api/title/upload/json_array Ünvan Ünvan-Json-Array
	// Aktar
	//
	// Bu fonksiyon entegrasyon işlemlerinde Json array olarak Ünvan bilgilerinin aktarılması için kullanılır..
	//
	// .
	// responses:
	//   200: StatusResponse200
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/customer/upload/json_array",apicontrollers.Upload_Customer_From_Json_Array)









	/*
	-------------------------------------------
	-------- Item								  -----------
	-------------------------------------------
	*/




	// swagger:route POST /api/position Ürün-Oluştur
	// Ekle
	//
	// Yeni bir Ürün kaydı oluşturur
	//
	// responses:
	//   201: StatusResponse201
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/Item", apicontrollers.POST_Item)





	// swagger:route GET /api/position  Ürün-Listesini-Getir
	// Liste
	//
	// Ürün Listesini getirir
	//
	// responses:
	//   200: PositionResponse
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.GET("/Item",apicontrollers.GET_Item)



	// swagger:route GET /api/Item/:id  Ürün-Getir
	// Oku.
	//
	// ID paramretresi ile belirtilen Ürün kaydını getirir.
	//
	// responses:
	//   200: PositionResponse
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	//   404: ErrorResponse404
	api.GET("/Item/:id",apicontrollers.GET_ItemByID)




	// swagger:route PUT /api/Item/:id  Ürün-Güncelle
	// Güncelle
	//
	// Ürün bilgisini günceller
	//
	// responses:
	//   200: StatusResponse200
	//   404: ErrorResponse404
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.PUT("/Item/:id",apicontrollers.PUT_Item)



	// swagger:route DELETE /api/Item/:id  Ürün-sil
	// Sil
	//
	// Ürün bilgisini sil
	//
	// responses:
	//   200: StatusResponse200
	//   404: ErrorResponse404
	//   400: ErrorResponse400
	//   401: ErrorResponse400
	api.DELETE("/Item/:id",apicontrollers.DELETE_Item)



	// swagger:route POST /api/Item/upload/excel  Ürün-Excel-Yükle
	// Excel Yükle
	//
	// Aktarım şablonu ile excel üzerinden Ürün bilgilerinin yüklenmesi için kullanılır.
	//
	// .
	// responses:
	//   200: StatusResponse200
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/Item/upload/excel",apicontrollers.Upload_Item_From_Excel)



	// swagger:route POST /api/Item/upload/json_array  Ürün-Json-Array
	// Aktar
	//
	// Bu fonksiyon entegrasyon işlemlerinde Json array olarak Pozisyon bilgilerinin aktarılması için kullanılır..
	//
	// .
	// responses:
	//   200: StatusResponse200
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/position/upload/json_array",apicontrollers.Upload_Item_From_Json_Array)








	/*
	-------------------------------------------
	-------- SalesHeader								  -----------
	-------------------------------------------
	*/




	// swagger:route POST /api/SalesHeader Satış Başlığı-Oluştur
	// Ekle
	//
	// Yeni bir Satış Başlığı kaydı oluşturur
	//
	// responses:
	//   201: StatusResponse201
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/SalesHeader", apicontrollers.POST_SalesHeader)





	// swagger:route GET /api/SalesHeader Satış Başlığı-Listesini-Getir
	// Liste
	//
	// Satış Başlığı Listesini getirir
	//
	// responses:
	//   200: LocationResponse
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.GET("/SalesHeader",apicontrollers.GET_SalesHeader)



	// swagger:route GET /api/location/:id Konum Konum-Getir
	// Oku.
	//
	// ID paramretresi ile belirtilen Konum kaydını getirir.
	//
	// responses:
	//   200: LocationResponse
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	//   404: ErrorResponse404
	api.GET("/SalesHeader/:id",apicontrollers.GET_SalesHeaderByID)




	// swagger:route PUT /api/SalesHeader/:id Satış Başlığı-Güncelle
	// Güncelle
	//
	// Satış Başlığı bilgisini günceller
	//
	// responses:
	//   200: StatusResponse200
	//   404: ErrorResponse404
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.PUT("/SalesHeader/:id",apicontrollers.PUT_SalesHeader)



	// swagger:route DELETE /api/SalesHeader/:id Satış Başlığı-Sil
	// Sil
	//
	// Satış Başlığı bilgisini sil
	//
	// responses:
	//   200: StatusResponse200
	//   404: ErrorResponse404
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.DELETE("/SalesHeader/:id",apicontrollers.DELETE_SalesHeader)



	// swagger:route POST /api/SalesHeader/upload/excel SalesHeader-Excel-Yükle
	// Excel Yükle
	//
	// Aktarım şablonu ile excel üzerinden SAtış Başlığı bilgilerinin yüklenmesi için kullanılır.
	//
	// .
	// responses:
	//   200: StatusResponse200
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/SalesHeader/upload/excel",apicontrollers.Upload_SalesHeader_From_Excel)



	// swagger:route POST /api/SalesHeader/upload/json_array Satış Başlığı-Json-Array
	// Aktar
	//
	// Bu fonksiyon entegrasyon işlemlerinde Json array olarak Satış Başlığı bilgilerinin aktarılması için kullanılır..
	//
	// .
	// responses:
	//   200: StatusResponse200
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/SalesHeader/upload/json_array",apicontrollers.Upload_SalesHeader_From_Json_Array)





	/*
	-------------------------------------------
	-------- SalesLine								  -----------
	-------------------------------------------
	*/




	// swagger:route POST /api/SalesLine  Satış Sırası-Oluştur
	// Ekle
	//
	// Yeni bir Satış Sırası kaydı oluşturur
	//
	// responses:
	//   201: StatusResponse201
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/SalesLine", apicontrollers.POST_SalesLine)





	// swagger:route GET /api/SalesLine  Satış Sırası-Listesini-Getir
	// Liste
	//
	// Satış Sırası Listesini getirir
	//
	// responses:
	//   200: OrganizationResponse
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.GET("/SalesLine",apicontrollers.GET_SalesLine)



	// swagger:route GET /api/SalesLine/:id  Satış Sırası-Getir
	// Oku.
	//
	// ID paramretresi ile belirtilen Satış Sırası kaydını getirir.
	//
	// responses:
	//   200: SalesLineResponse
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	//   404: ErrorResponse404
	api.GET("/SalesLine/:id",apicontrollers.GET_SalesLineByID)




	// swagger:route PUT /api/SalesLine/:id  Satış Sırası-Güncelle
	// Güncelle
	//
	// Satış Sırası bilgisini günceller
	//
	// responses:
	//   200: StatusResponse200
	//   404: ErrorResponse404
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.PUT("/SalesLine/:id",apicontrollers.PUT_SalesLine)



	// swagger:route DELETE /api/SalesLine/:id  Satış Sırası-Sil
	// Sil
	//
	// Satış Sırası bilgisini sil
	//
	// responses:
	//   200: StatusResponse200
	//   404: ErrorResponse404
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.DELETE("/SalesLine/:id",apicontrollers.DELETE_SalesLine)



	// swagger:route POST /api/SalesLine/upload/excel  Satış Sırası-Excel-Yükle
	// Excel Yükle
	//
	// Aktarım şablonu ile excel üzerinden Satış Sırası bilgilerinin yüklenmesi için kullanılır.
	//
	// .
	// responses:
	//   200: StatusResponse200
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/SalesLine/upload/excel",apicontrollers.Upload_SalesLine_From_Excel)



	// swagger:route POST /api/SalesLine/upload/json_array  Satış Sırası-Json-Array
	// Aktar
	//
	// Bu fonksiyon entegrasyon işlemlerinde Json array olarak Satış Sırası bilgilerinin aktarılması için kullanılır..
	//
	// .
	// responses:
	//   200: StatusResponse200
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/SalesLine/upload/json_array",apicontrollers.Upload_SalesLine_From_Json_Array)






	/*
	-------------------------------------------
	-------- SalesPrice								  -----------
	-------------------------------------------
	*/


	// swagger:route POST /api/SalesPrice  Satış Fiyatı-Oluştur
	// Ekle
	//
	// Yeni bir Satış Fiyatı kaydı oluşturur
	//
	// responses:
	//   201: StatusResponse201
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/SalesPrice", apicontrollers.POST_SalesPrice)





	// swagger:route GET /api/SalesPrice  SalesPrice-Listesini-Getir
	// Liste
	//
	// Personel Listesini getirir
	//
	// responses:
	//   200: StaffResponse
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.GET("/SalesPrice", apicontrollers.GET_SalesPrice)



	// swagger:route GET /api/SalesPrice/:id  Satış Fiyatı-Getir
	// Oku.
	//
	// ID paramretresi ile belirtilen Satış Fiyatı kaydını getirir.
	//
	// responses:
	//   200: StaffResponse
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	//   404: ErrorResponse404
	api.GET("/SalesPrice/:id", apicontrollers.GET_SalesPriceByID)




	// swagger:route PUT /api/SalesPrice/:id  Satış Fiyatı-Güncelle
	// Güncelle
	//
	// Satış Fiyatı bilgisini günceller
	//
	// responses:
	//   200: StatusResponse200
	//   404: ErrorResponse404
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.PUT("/SalesPrice/:id", apicontrollers.PUT_SalesPrice)



	// swagger:route DELETE /api/SalesPrice/:id  Satış Fiyatı-Sil
	// Sil
	//
	// Satış Fiyatı bilgisini sil
	//
	// responses:
	//   200: StatusResponse200
	//   404: ErrorResponse404
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.DELETE("/SalesPrice/:id", apicontrollers.DELETE_SalesPrice)



	// swagger:route POST /api/SalesPrice/upload/excel  Satış Fiyatı-Excel-Yükle
	// Excel Yükle
	//
	// Aktarım şablonu ile excel üzerinden Satış Fiyatı bilgilerinin yüklenmesi için kullanılır.
	//
	// .
	// responses:
	//   200: StatusResponse200
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/SalesPrice/upload/excel", apicontrollers.Upload_SalesPrice_From_Excel)



	// swagger:route POST /api/SalesPrice/upload/json_array  Satış Fiyatı-Json-Array
	// Aktar
	//
	// Bu fonksiyon entegrasyon işlemlerinde Json array olarak Satış Fiyatı bilgilerinin aktarılması için kullanılır..
	//
	// .
	// responses:
	//   200: StatusResponse200
	//   400: ErrorResponse400
	//   401: ErrorResponse401
	api.POST("/SalesPrice/upload/json_array", apicontrollers.Upload_SalesPrice_From_Json_Array)



	app.Run(":" + port)
}

