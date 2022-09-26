package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Quantity uint8   `json:"quantity"`
	Price    float32 `json:"price"`
}

var items = []Item{
	{ID: "1", Title: "Brush", Quantity: 2, Price: 12.00},
	{ID: "2", Title: "Bed", Quantity: 20, Price: 100.00},
	{ID: "3", Title: "Pan", Quantity: 2, Price: 10.50},
}

// Display items as a JSON format
func getItems(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, items)
}

// If all the fields are correct, append an item into the items array
func addItem(context *gin.Context) {

	var newItem Item

	// Error checking for JSON (correct object names etc.)
	if err := context.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)

	// Display the JSON that was just appended
	context.IndentedJSON(http.StatusCreated, newItem)
}

// Get the id from the url and display that certain item
func getItem(context *gin.Context) {
	id := context.Param("id")

	item, err := getItemById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, item)
}

// Loop through items and find the item with the given id
func getItemById(id string) (*Item, error) {
	for i, t := range items {
		if t.ID == id {
			return &items[i], nil
		}
	}

	return nil, errors.New("Item not found")
}

func main() {

	router := gin.Default()

	// GET request for all items
	router.GET("/items", getItems)

	// POST request for adding items
	router.POST("/items", addItem)

	// GET request for single item
	router.GET("/items/:id", getItem)

	router.Run()
}
