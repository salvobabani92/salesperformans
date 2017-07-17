

package main

import (
	"log"
	"os"
	"html/template"
	"github.com/gin-gonic/gin"

	"github.com/salvobabani92/salesperformans.com/config"
	"github.com/salvobabani92/salesperformans.com/controller"
	"github.com/salvobabani92/salesperformans.com/models"

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
	models.UnitofMeasure{}.CreateTable()
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

}