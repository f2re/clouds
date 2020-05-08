// server
package main

import (
	// "database/sql"
	"./handlers"
	"./models"

	"github.com/labstack/echo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)



func main() {

	db, err := gorm.Open("sqlite3", "db/tcmstorage.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// var Item models.Item
	// Миграция схем
	db.AutoMigrate(&models.Item{},&models.Category{},&models.Tabel{},&models.Image{})
	

	// Create a new instance of Echo
	e := echo.New()

	e.File("/", "../web/public/index.html")
	e.Static("/js", "../web/public/js")
	e.Static("/css", "../web/public/css")
	e.Static("/img", "../web/public/img")
	e.Static("/uploads", "../web/public/uploads")

	// работаем с техникой
	e.GET("/api/items", handlers.GetItems(db))
    e.POST("/api/items", handlers.SaveItem(db))
    e.PUT("/api/items/:id", handlers.UpdateItem(db))
	e.DELETE("/api/items/:id", handlers.DeleteItem(db))
	
	// работаем с категориями
	e.GET("/api/categories", handlers.GetCategories(db))
	e.POST("/api/categories", handlers.SaveCategories(db))
    e.PUT("/api/categories/:id", handlers.UpdateCategories(db))
	e.DELETE("/api/categories/:id", handlers.DeleteCategories(db))

	// Start as a web server
	e.Logger.Fatal(e.Start(":8000"))
}

