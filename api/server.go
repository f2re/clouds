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
	dbpath := "api/db/"

	// TCM database
	db, err := gorm.Open("sqlite3", dbpath+"tcmstorage.db")
	if err != nil {
		panic("failed to connect database tcmstorage.db")
	}
	defer db.Close()

	// User database
	// udb, err := gorm.Open("sqlite3", dbpath+"usertcm.db")
	// if err != nil {
	// 	panic("failed to connect database usertcm.db")
	// }
	// defer db.Close()

	// var Item models.Item
	// Миграция схем
	db.AutoMigrate(	&models.Item{},
					&models.Tabel{},
					&models.Image{}, 
					&models.UserItems{})
	

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
	e.Static("/uploads", "web/uploads")
	e.Static("/fonts", "web/dist/fonts")

	// работаем с техникой
	e.GET("/api/items", handlers.GetItems(db))
	e.GET("/api/useritems", handlers.GetUserItems(db))
	e.GET("/api/search", handlers.GetItems(db))
	e.GET("/api/search/:q", handlers.SearchItems(db))
	e.GET("/api/item/:slug", handlers.GetItem(db))
	e.GET("/api/useritem/:id", handlers.GetUserItem(db))
    e.POST("/api/items", handlers.SaveItem(db))
    e.POST("/api/adduseritem", handlers.AddUserItem(db))
    e.PUT("/api/item/:slug", handlers.UpdateItem(db))
    e.PUT("/api/useritem/:id", handlers.UpdateUserItem(db))
	e.DELETE("/api/item/:slug", handlers.DeleteItem(db)) 
	e.DELETE("/api/removeuseritem/:id", handlers.DeleteUserItem(db)) 

	e.GET("/api/category/:category", handlers.GetCategory(db))

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