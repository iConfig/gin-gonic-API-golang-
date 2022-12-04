package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){


	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", addBooks)


	router.Run("127.0.0.1:8080")

}

// declaring book struct 
type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Price string `json:"price"`
}

// adding data to book struct 
var books = []book{
	{ID:"1", Title:" Go for starters", Author:"John McTomy", Price:"50"},
	{ID:"2", Title:" Go for Advanced Programmers", Author:"Julian Rachael", Price:"80"},
	{ID:"3", Title:" Go for Professioners", Author:"Dan Real", Price:"100"},
}

// function to return list of books
func getBooks(c *gin.Context){
	c.IndentedJSON( http.StatusOK, books)
}

//function to post new books  
func addBooks(c *gin.Context){

	var newbook book 

	//call bindjson to bind the received json data

	if err := c.BindJSON(&newbook); err != nil {
		return 
	}

	// add new books to slice 

	books = append(books,newbook)
	c.IndentedJSON(http.StatusCreated, newbook)
}

//function to get books by id 
func getBookById(c *gin.Context){
	id := c.Param("id")

	//loop through the list of books  
	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

