package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Item struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Quantity uint8   `json:"quantity"`
	Price    float32 `json:"price"`
}

var items = []Item{
	{ID: "1", Title: "Brush", Quantity: 50, Price: 12.00},
	{ID: "2", Title: "Bed", Quantity: 200, Price: 0.80},
	{ID: "3", Title: "Pan", Quantity: 100, Price: 10.50},
	{ID: "4", Title: "Water", Quantity: 150, Price: 0.50},
	{ID: "5", Title: "Desk", Quantity: 10, Price: 100.00},
	{ID: "6", Title: "Chair", Quantity: 20, Price: 69.99},
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

	item, err, _ := getItemById(id) // Get the returned item, the error in case it's not nil, and skip the index

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"}) // Print error message
		return
	}
	context.IndentedJSON(http.StatusOK, item)
}

// Take the id from the parameter, then get the index of the item based on the id, delete it and shift everything
func removeItem(context *gin.Context) {
	id := context.Param("id")

	_, err, i := getItemById(id) // Skip the returned item, get the error and the index of the item
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"}) // Print error message
		return
	}

	copy(items[i:], items[i+1:]) // Shift a[i+1:] left one index.
	items = items[:len(items)-1] // Truncate slice.

	context.IndentedJSON(http.StatusOK, items)
}

// Loop through items and find the item with the given id
func getItemById(id string) (*Item, error, int) {
	for i, t := range items {
		if t.ID == id {
			return &items[i], nil, i
		}
	}
	return nil, errors.New("Item not found"), -1
}

func main() {

	// Router
	router := gin.Default()

	// Custom CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},                                                                  // Bad for security but this is purely a personal project and not intended for commercial purposes
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch}, // Allowed methods of sending requests
		AllowCredentials: true,
	})

	router.Use(c) // Used to bypass CORS policy, god bless whoever made this 🙏

	// GET request for all items
	router.GET("/items", getItems)

	// POST request for adding items
	router.POST("/items", addItem)

	// GET request for single item
	router.GET("/items/:id", getItem)

	// PATCH request for deleting an item (Changed from DELETE to PATCH, works better for some reason?)
	router.PATCH("/items/:id", removeItem)

	router.Run(":8080")

}
