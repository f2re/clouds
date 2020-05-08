// server
package main

import (
	"database/sql"
	"./handlers"

	"github.com/labstack/echo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)




func main() {

	db := initDB("tcmstorage.db")
	migrate(db)

	// Create a new instance of Echo
	e := echo.New()

	e.File("/", "public/index.html")

	e.GET("/api/items", handlers.GetItems(db))
    e.PUT("/api/items", handlers.PutItem(db))
    e.DELETE("/api/items/:id", handlers.DeleteItem(db))

	// Start as a web server
	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS items(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
        name VARCHAR NOT NULL
        name VARCHAR NOT NULL
        name VARCHAR NOT NULL
        name VARCHAR NOT NULL
        name VARCHAR NOT NULL
        name VARCHAR NOT NULL
        name VARCHAR NOT NULL
        name VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
