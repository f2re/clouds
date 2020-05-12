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

// response for request with redirect urls
type Response struct{
	Status string `json:"status"`
	URL string `json:"URL"`
}

// GetItems endpoint
func GetItems(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
        var Items []models.Item
		return c.JSON(http.StatusOK, db.Preload("Tabel").Preload("Image").Find(&Items) )
	}
}

// GetItems endpoint
func GetCategory(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var Items []models.Item
		category, err := strconv.ParseUint(c.Param("category"),10,32)
		if ( err == nil ){
			return c.JSON(http.StatusOK, db.Preload("Tabel").Preload("Image").Where(&models.Item{ Category:category }).Find(&Items) )
		}
		return c.JSON(http.StatusOK, err)
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

// GetItems endpoint
func SearchItems(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var Items []models.Item
		q := c.Param("q")
		return c.JSON(http.StatusOK, db.Preload("Tabel").Preload("Image").Where("name LIKE ?", "%"+q+"%").Find(&Items) )
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

		path := ""
		if err == nil {
			
			src, err := img.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			// Destination
			crutime := time.Now().Unix()
			path = "/uploads/"+strconv.FormatInt(crutime, 10)+filepath.Ext(img.Filename)
			dst, err := os.Create( "web/dist"+path )

			if err != nil {
				return err
			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}
		}

		name         := c.FormValue("name")
		category, err:= strconv.ParseUint(c.FormValue("category"),10,32)
		index        := c.FormValue("index")
		snabjenie    := c.FormValue("snabjenie")
		kvt          := c.FormValue("kvt")
		nomenklature := c.FormValue("nomenklature")
		dovorgan     := c.FormValue("dovorgan")
		reqorgan     := c.FormValue("reqorgan")
		explorgan    := c.FormValue("explorgan")
		creator      := c.FormValue("creator")
		description  := c.FormValue("description")
		destination  := c.FormValue("destination")
		composition  := c.FormValue("composition")
		tth          := c.FormValue("tth")
		// decode tabel variable
		var tabel models.Tabel
		json.Unmarshal([]byte(c.FormValue("tabel")), &tabel)

		res := db.Save(&models.Item{ 
			Name: name, 
			Slug: slug.Make(name),
			Category: category,
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
		// get record from DB
		var Item models.Item
		s := c.Param("slug")
		db.Preload("Tabel").Preload("Image").Where("slug = ?", s).First(&Item)

		//-----------
		// Read file
		//-----------
		img,  err := c.FormFile("Image")

		path := ""
		if err == nil {
			
			src, err := img.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			// Destination
			crutime := time.Now().Unix()
			path = "/uploads/"+strconv.FormatInt(crutime, 10)+filepath.Ext(img.Filename)
			dst, err := os.Create( "web/dist"+path )

			if err != nil {
				return err
			}
			defer dst.Close()

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}
		}else{
			path = ""
		}
		category, err := strconv.ParseUint(c.FormValue("category"),10,32)
		Item.Name         = c.FormValue("name")
		Item.Category     = category
		Item.Slug         = slug.Make(c.FormValue("name"))
		Item.Index        = c.FormValue("index")
		Item.Snabjenie    = c.FormValue("snabjenie")
		Item.KVT          = c.FormValue("kvt")
		Item.Nomenklature = c.FormValue("nomenklature")
		Item.Dovorgan     = c.FormValue("dovorgan")
		Item.Reqorgan     = c.FormValue("reqorgan")
		Item.Explorgan    = c.FormValue("explorgan")
		Item.Creator      = c.FormValue("creator")
		Item.Description  = c.FormValue("description")
		Item.Destination  = c.FormValue("destination")
		Item.Composition  = c.FormValue("composition")
		Item.TTH          = c.FormValue("tth")
		// decode tabel variable
		// var tabel models.Tabel
		json.Unmarshal([]byte(c.FormValue("tabel")), &Item.Tabel)
		if ( path != "" ){
			if (db.Model(&Item).Association("Image").Count()>0){
				Item.Image.Path = path
			}else{
				Item.Image = models.Image{Path:path}
			}
		}

		res := db.Save(&Item)

		if res.Error != nil {
			errors := res.GetErrors()
			errstr := ""
			for _, err := range errors {
				errstr = errstr+err.Error()
			}
			return c.JSON(http.StatusOK, errstr  )
		}

		resp := &Response{
			Status:"OK" ,
			URL:Item.Slug,
		}

		return c.JSON(http.StatusOK, resp )
		// return c.String(http.StatusOK, "OK")
	}
}

// DeleteItem endpoint
func DeleteItem(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var Item models.Item
		s := c.Param("slug")
		db.Preload("Tabel").Preload("Image").Where("slug = ?", s).Delete(&Item)
		
        return c.String(http.StatusOK, "OK")
	}
}
