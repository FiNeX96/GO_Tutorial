package main


import (
		"database/sql"
		"fmt"
		"net/http"
		"github.com/gin-gonic/gin"
		"os"
		_ "github.com/mattn/go-sqlite3"
)



type Book struct {

	ID int `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price int `json:"price"`

}


func main (){

	router := gin.Default()

	// Lets turn this API into a web server to manage books for a library 



	// 1. Start by creating a struct to represent a book. It should have the following fields: ID, Title, Artist, Price


	/* 2 . Implement a JSON serialization method in your struct, you can do this as follows :

		type Book struct {
			ID int `json:"id"`
			Title string `json:"title"` 
			(etc...)


		This will tell GIN how to serialize a object of your struct into a JSON object to be returned by the API.
	*/


	/* 3 . Create a sqlite3 database called books.db and create a table called books with the following fields: ID, Title, Artist, Price

		To create and open a sqlite3 database you may use the following code:

		db, err := sql.Open("sqlite3", "books.db") . Dont forget to take care of possible errors =)

		To execute comands in it you can do:

		db.Exec( sql_command_here )

		If you dont know sql or just dont remember it, here you go :

		db.Exec("CREATE TABLE books (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, artist TEXT, price INTEGER)")


	*/





	db, err := sql.Open("sqlite3", "books.db")

	if err != nil {
		fmt.Println("Error opening database!")
		os.Exit(1)
	}

	_ , err = db.Exec("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, artist TEXT, price INTEGER)")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	
	/* 4 . Now that we have a database, create a endpoint to add a book to it using the POST method.

	 To test this endpoint you may use for example postman or curl:

	curl -X POST -H "Content-Type: application/json" -d '{
    "Title": "Very Cool Book",
    "Artist": "Pablo Picasso",
    "Price": 33
	}' http://localhost:8080/addbook


	*/

	router.POST("/addbook", func(c *gin.Context) {
		addBook(c, db)
	})


	// 5 . Now that we have books in the database, create a endpoint and a function to get all the books in the database using the GET method.

	router.GET("/getbooks", func(c *gin.Context) {
		getBooks(c,db)
	})



	// 6. Develop a endpoint and a function to get a book by its ID using the GET method.

	router.GET("/getbook/:id", func(c *gin.Context) {
		getBookByID(c,db)
	})


	/* (Extra)

	 Develop endpoints to:
	 - Update a book information
	 - Delete a book from the database

	*/

	router.Run(":8080")


}


func addBook(c *gin.Context, db *sql.DB){

	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		fmt.Println(err)
		return
	}

	_, err := db.Exec("INSERT INTO books (title, artist, price) VALUES (?, ?, ?)",  newBook.Title, newBook.Artist, newBook.Price)

	if err != nil {
		fmt.Println("Error inserting book!")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book added successfully!",
	})
}

func getBooks(c *gin.Context, db *sql.DB){

	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		fmt.Println("Error getting books!")
	}

	var books []Book

	for rows.Next() {

		var book Book

		err = rows.Scan(&book.ID, &book.Title, &book.Artist, &book.Price)

		if err != nil {
			fmt.Println("Error scanning books!")
		}

		books = append(books, book)

	}

	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func getBookByID(c *gin.Context, db *sql.DB){

	id := c.Param("id")

	var book Book

	err := db.QueryRow("SELECT * FROM books WHERE id = ?", id).Scan(&book.ID, &book.Title, &book.Artist, &book.Price)

	if err != nil {
		fmt.Println("Error getting book!")
	}

	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}