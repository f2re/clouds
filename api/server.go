// server
package main

import (
	// "database/sql"
	"./handlers"
	"./models"
	"fmt"
	// "os"
	"os/exec"
	"log"
	"runtime"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)



func main() {
	db, err := gorm.Open("sqlite3", "api/db/tcmstorage.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// var Item models.Item
	// Миграция схем
	db.AutoMigrate(&models.Item{},&models.Category{},&models.Tabel{},&models.Image{})
	

	// Create a new instance of Echo
	e := echo.New()
	e.Debug = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:8000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	  }))

	
	e.Static("/js", "web/dist/js")
	e.Static("/css", "web/dist/css")
	e.Static("/img", "web/dist/img")
	e.Static("/uploads", "web/dist/uploads")
	e.Static("/fonts", "web/dist/fonts")

	// работаем с техникой
	e.GET("/api/items", handlers.GetItems(db))
	e.GET("/api/search/:q", handlers.SearchItems(db))
	e.GET("/api/item/:slug", handlers.GetItem(db))
    e.POST("/api/items", handlers.SaveItem(db))
    e.PUT("/api/item/:slug", handlers.UpdateItem(db))
	e.DELETE("/api/item/:slug", handlers.DeleteItem(db)) 
	
	// работаем с категориями
	e.GET("/api/categories", handlers.GetCategories(db))
	e.POST("/api/categories", handlers.SaveCategories(db))
    e.PUT("/api/categories/:id", handlers.UpdateCategories(db))
	e.DELETE("/api/categories/:id", handlers.DeleteCategories(db))

	e.File("/*", "web/dist/index.html")
	
	openbrowser("http://localhost:8000")
	// Start as a web server
	e.Logger.Fatal(e.Start(":8000"))
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}