package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"../models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetItems endpoint
func GetItems(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetItems(db))
	}
}

// PutItem endpoint
func PutItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new Item
		var Item models.Item
		// Map imcoming JSON body to the new Item
		c.Bind(&Item)
		// Add a Item using our new model
		id, err := models.PutItem(db, Item.Name)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

// DeleteItem endpoint
func DeleteItem(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a Item
		_, err := models.DeleteItem(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// Handle errors
		} else {
			return err
		}
	}
}