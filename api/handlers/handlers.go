package handlers

import (
	"github.com/jinzhu/gorm"
	"net/http"
	// "strconv"

	"../models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetItems endpoint
func GetItems(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
        var Item models.Item
		return c.JSON(http.StatusOK, db.Find(&Item) )
	}
}

// PutItem endpoint
func SaveItem(db *gorm.DB) echo.HandlerFunc {
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