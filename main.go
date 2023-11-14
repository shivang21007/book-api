package main

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	Id string  `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int  `json:"quantity"`

}

var books = []book {
	{Id: "1", Title: "In seaarch of Lost Time", Author: "Mercel Proust", Quantity: 2},
	{Id: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{Id: "3", Title: "War and Peace", Author: "Leao Tolstoy", Quantity: 6},
	{Id: "4", Title: "NCERT Science", Author: "CBSE Board", Quantity: 3},
} 

// simply fetch all books
func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

// create a new book
func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Book is Successfully Created.", "Created-Book": newBook})
}


// fetch any book by it's ID
func bookbyId(c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// (helper func) find any book by giving his ID
func getBookById(id string) (*book, error){
	for i, b := range books {
		if b.Id == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// checkout book by its ID
func checkoutBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Book doesn't exist."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, gin.H{
		"message":"Book is Successafully Checked-Out",
		"checked-out book":book,
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

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, gin.H{
		"message":"Book is Successfully returned",
		"returned book":  book,
	})
}

// Main function
func main(){
	router := gin.Default()

	router.GET("/")

	router.GET("/books", getBooks) 

	router.GET("/books/:id", bookbyId) 

	router.POST("/books", createBook)

	router.PATCH("/checkout", checkoutBook) 
	// http://localhost:8080/checkout?id=2
	router.PATCH("/return", returnBook)
	// http://localhost:8080/return?id=2
	router.Run("localhost:8080")
};