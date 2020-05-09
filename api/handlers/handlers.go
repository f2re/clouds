package handlers

import (
	"os"
	"io"
	// "io/ioutil"
	"github.com/jinzhu/gorm"
	"path/filepath"
	"strconv"
	"time"
	"github.com/gosimple/slug"
	"net/http"
	"encoding/json"

	"../models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetItems endpoint
func GetItems(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
        var Items []models.Item
		return c.JSON(http.StatusOK, db.Preload("Tabel").Preload("Image").Find(&Items) )
	}
}

// GetItems endpoint
func GetItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var Item models.Item
		slug := c.Param("slug")
		return c.JSON(http.StatusOK, db.Preload("Tabel").Preload("Image").Where(&models.Item{ Slug:slug }).First(&Item) )
	}
}

// PutItem endpoint
func SaveItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new Item
		// var Item models.Item
		//-----------
		// Read file
		//-----------
		img,  err := c.FormFile("Image")

		if err != nil {
			return err
		}
		src, err := img.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		crutime := time.Now().Unix()
		path := "/uploads/"+strconv.FormatInt(crutime, 10)+filepath.Ext(img.Filename)
		dst, err := os.Create( "../web/dist"+path )

		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		name := c.FormValue("name")
		// category := c.FormValue("category")
		index := c.FormValue("index")
		snabjenie := c.FormValue("snabjenie")
		kvt := c.FormValue("kvt")
		nomenklature := c.FormValue("nomenklature")
		dovorgan := c.FormValue("dovorgan")
		reqorgan := c.FormValue("reqorgan")
		explorgan := c.FormValue("explorgan")
		creator := c.FormValue("creator")
		description := c.FormValue("description")
		destination := c.FormValue("destination")
		composition := c.FormValue("composition")
		tth := c.FormValue("tth")
		// decode tabel variable
		var tabel models.Tabel
		json.Unmarshal([]byte(c.FormValue("tabel")), &tabel)

		res := db.Create(&models.Item{ 
			Name: name, 
			Slug: slug.Make(name),
			// Category: category,
			Index: index,
			Snabjenie: snabjenie,
			KVT: kvt,
			Nomenklature: nomenklature,
			Dovorgan: dovorgan,
			Reqorgan: reqorgan,
			Explorgan: explorgan,
			Creator: creator,
			Description: description,
			Destination: destination,
			Composition: composition,
			TTH: tth,
			Image: models.Image{Path:path},
			Tabel: tabel,
		})
		if res.Error != nil {
			errors := res.GetErrors()
			errstr := ""
			for _, err := range errors {
				errstr = errstr+err.Error()
			}
			return c.JSON(http.StatusOK, errstr  )
		}
		return c.String(http.StatusOK, "OK")
	}
}

// PutItem endpoint
func UpdateItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new Item
		var Item models.Item
		// Map imcoming JSON body to the new Item
        c.Bind(&Item)
        return nil
		// Add a Item using our new model
		// id, err := models.PutItem(db, Item.Name)
		// // Return a JSON response if successful
		// if err == nil {
		// 	return c.JSON(http.StatusCreated, H{
		// 		"created": id,
		// 	})
		// 	// Handle any errors
		// } else {
		// 	return err
		// }
	}
}

// DeleteItem endpoint
func DeleteItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
        // id, _ := strconv.Atoi(c.Param("id"))
        return nil
		// Use our new model to delete a Item
		// _, err := models.DeleteItem(db, id)
		// // Return a JSON response on success
		// if err == nil {
		// 	return c.JSON(http.StatusOK, H{
		// 		"deleted": id,
		// 	})
		// 	// Handle errors
		// } else {
		// 	return err
		// }
	}
}

/**
 *
 * КАТЕГОРИИ
 * 
 */


// GetItems endpoint
func GetCategories(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
        var Category models.Category
		return c.JSON(http.StatusOK, db.Find(&Category) )
	}
}

// PutCategory endpoint
func SaveCategories(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new Category
		var Category models.Category
		// Map imcoming JSON body to the new Category
        c.Bind(&Category)
        return nil
		// Add a Category using our new model
		// id, err := models.PutCategory(db, Category.Name)
		// // Return a JSON response if successful
		// if err == nil {
		// 	return c.JSON(http.StatusCreated, H{
		// 		"created": id,
		// 	})
		// 	// Handle any errors
		// } else {
		// 	return err
		// }
	}
}

// PutCategory endpoint
func UpdateCategories(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new Category
		var Category models.Category
		// Map imcoming JSON body to the new Category
        c.Bind(&Category)
        return nil
		// Add a Category using our new model
		// id, err := models.PutCategory(db, Category.Name)
		// // Return a JSON response if successful
		// if err == nil {
		// 	return c.JSON(http.StatusCreated, H{
		// 		"created": id,
		// 	})
		// 	// Handle any errors
		// } else {
		// 	return err
		// }
	}
}

// DeleteCategory endpoint
func DeleteCategories(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
        // id, _ := strconv.Atoi(c.Param("id"))
        return nil
		// Use our new model to delete a Category
		// _, err := models.DeleteCategory(db, id)
		// // Return a JSON response on success
		// if err == nil {
		// 	return c.JSON(http.StatusOK, H{
		// 		"deleted": id,
		// 	})
		// 	// Handle errors
		// } else {
		// 	return err
		// }
	}
}