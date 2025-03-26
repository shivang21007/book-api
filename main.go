package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	OriginalQty int    `json:"original_qty"`
	CurrentQty  int    `json:"current_qty"`
}

var books = []book{
	{Id: "1", Title: "In search of Lost Time", Author: "Marcel Proust", OriginalQty: 2, CurrentQty: 2},
	{Id: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", OriginalQty: 5, CurrentQty: 5},
	{Id: "3", Title: "War and Peace", Author: "Leo Tolstoy", OriginalQty: 6, CurrentQty: 6},
	{Id: "4", Title: "NCERT Science", Author: "CBSE Board", OriginalQty: 3, CurrentQty: 3},
}

// simply fetch all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// create a new book
// create a new book
func createBook(c *gin.Context) {
	var newBookFields struct {
		Id       string `json:"id"`
		Title    string `json:"title"`
		Author   string `json:"author"`
		Quantity int    `json:"quantity"`
	}

	if err := c.BindJSON(&newBookFields); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// Check if a book with the same ID already exists
	if _, err := getBookById(newBookFields.Id); err == nil {
		// Book with the same ID exists, respond with a message
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Book with the same ID already exists. Please create a new Book with new Fields or Update it."})
		return
	}

	// Check if a book with the same Title already exists
	if bookExistsByTitle(newBookFields.Title) {
		// Book with the same ID exists, respond with a message
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Book with the same Title already exists. Please Try to Update it if required."})
		return
	}

	// Create a new book with the provided fields
	newBook := book{
		Id:          newBookFields.Id,
		Title:       newBookFields.Title,
		Author:      newBookFields.Author,
		OriginalQty: newBookFields.Quantity,
		CurrentQty:  newBookFields.Quantity,
	}

	// Append the new book
	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Book is Successfully Created.", "Created-Book": newBook})
}

// (helper func) check if a book with the same title already exists
func bookExistsByTitle(title string) bool {
	for _, b := range books {
		if b.Title == title {
			return true
		}
	}
	return false
}

// fetch any book by its ID
func bookbyId(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// Update an existing book by its ID
// Update an existing book by its ID
func updateBookById(c *gin.Context) {
	id := c.Param("id")

	// Find the book using the getBookById helper function
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	var updateFields struct {
		Id       string `json:"id"`
		Title    string `json:"title"`
		Author   string `json:"author"`
		Quantity int    `json:"quantity"`
	}

	if err := c.BindJSON(&updateFields); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// Update the book with the provided fields
	book.Id = updateFields.Id
	book.Title = updateFields.Title
	book.Author = updateFields.Author
	book.OriginalQty = updateFields.Quantity
	book.CurrentQty = updateFields.Quantity

	// Respond with a JSON message indicating that the book was successfully updated
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book is Successfully Updated", "updated_book": book})
}

// Delete any book by its ID
func deletebyId(c *gin.Context) {
	id := c.Param("id")

	// Find the book using the getBookById helper function
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	// Get the index of the book in the 'books' slice
	index, err := getIndexById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error finding book index."})
		return
	}

	// Remove the book from the 'books' slice
	books = append(books[:index], books[index+1:]...)

	// Respond with a JSON message indicating that the book was successfully deleted
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book is Successfully Deleted", "deleted_book": book})
}

// (helper func) find the index of any book by giving its ID
func getIndexById(id string) (int, error) {
	for i, b := range books {
		if b.Id == id {
			return i, nil
		}
	}
	return -1, errors.New("book not found")
}

// (helper func) find any book by giving its ID
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.Id == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// checkout book by its ID
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book doesn't exist."})
		return
	}

	if book.CurrentQty <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.CurrentQty -= 1
	c.IndentedJSON(http.StatusOK, gin.H{
		"message":          "Book is Successfully Checked-Out",
		"checked-out book": book,
	})
}

// return any book by its ID
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book doesn't exist."})
		return
	}

	// Check if returning the book would exceed the original quantity
	if book.CurrentQty >= book.OriginalQty {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Cannot return more books than the original quantity."})
		return
	}

	book.CurrentQty += 1
	c.IndentedJSON(http.StatusOK, gin.H{
		"message":            "Book is Successfully returned",
		"returned book":      book,
		"Quantity-available": book.CurrentQty,
	})
}

// Main function
func main() {
	router := gin.Default()

	router.GET("/")

	router.GET("/books", getBooks)

	router.GET("/books/:id", bookbyId)

	router.POST("/books", createBook)

	router.PUT("/books/:id", updateBookById)

	router.DELETE("/books/:id", deletebyId)

	router.PATCH("/checkout", checkoutBook)
	// http://localhost:8080/checkout?id=2
	router.PATCH("/return", returnBook)
	// http://localhost:8080/return?id=2
	router.Run(":8080")
}
