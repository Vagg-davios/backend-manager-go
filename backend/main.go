package main

import (
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

func main() {

	router := gin.Default()

	// GET request for localhost:8080/items
	router.GET("/items", getItems)

	// POST request for localhost:8080/items
	router.POST("/items", addItem)

	router.Run()
}
