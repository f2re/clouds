package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Item is a struct containing Item data
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`  // название изделия
	Slug string `json:"slug"`  // строка адреса
	Index string `json:"index"` // индекс иделия
	Snabjenie string `json:"snabjenie"` // принят на снабжение
	KVT string `json:"KVT"` // код КВТ
	nomenklature string `json:"nomenklature"` // номенклатурный номер
	dovorgan string `json:"dovorgan"`  // довольствующий орган
	reqorgan string `json:"reqorgan"` // заказывающий орган
	explrgan string `json:"explrgan"` // эксплуатирующий орган
	creator string `json:"creator"` // изготовитель
	description string `json:"description"` // описание
	destination string `json:"destination"` // назначение
	composition string `json:"composition"` // состав
	TTH string `json:"TTH"` // ТТХ
	normal string `json:"normal"` // нормы табелизации
}

// ItemCollection is collection of Items
type ItemCollection struct {
	Items []Item `json:"items"`
}

// GetItems from the DB
func GetItems(db *sql.DB) ItemCollection {
	sql := "SELECT * FROM Items"
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := ItemCollection{}
	for rows.Next() {
		Item := Item{}
		err2 := rows.Scan(&Item.ID, &Item.Name)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Items = append(result.Items, Item)
	}
	return result
}

// PutItem into DB
func PutItem(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO Items(name) VALUES(?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteItem from DB
func DeleteItem(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM Items WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}