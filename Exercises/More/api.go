package main


import (
		"database/sql"
		"fmt"
		"net/http"
		"github.com/gin-gonic/gin"
		_"github.com/mattn/go-sqlite3"
)


/*

In this exercise, we are going to build a simple API using the GIN Framework and a sqlite database.
To get the GIN framework, run the following command in your terminal: go get github.com/gin-gonic/gin 
To get the sqlite driver, run the following command in your terminal: go get github.com/mattn/go-sqlite3
Analyse the code below and try to understand what is happening.
After that, complete the exercises below.

If you are struggling, you may find these guides helpful:

- https://go.dev/doc/tutorial/web-service-gin
- https://zetcode.com/golang/sqlite3/

Can always ask for help aswell =)

*/

func main (){

	router := gin.Default()

	// Basic Endpoint creation, you can access it by going to http://localhost:8080/hello

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Lets turn this API into a web server to manage books for a library 



	// 1. Start by creating a struct to represent a book. It should have the following fields: ID, Title, Artist, Price


		/* 2 . Implement a JSON serialization method in your struct, you can do this as follows 

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

		db.Exec("CREATE TABLE books (id INTEGER PRIMARY KEY, title TEXT, artist TEXT, price INTEGER)")


	*/

	/* 4 . Now that we have a database, create a endpoint to add a book to it using the POST method.

	 To test this endpoint you may use for example postman or curl:

	curl -X POST -H "Content-Type: application/json" -d '{
    "Title": "Very Cool Book",
    "Artist": "Pablo Picasso",
    "Price": 33
	}' http://localhost:8080/addbook

	( this is a example, adjust it to your api )


	*/


	// 5 . Now that we have books in the databse, create a endpoint to get all the books in the database using the GET method.


	// 6. Develop a endpoint to get a book by its ID using the GET method.


	/* (Extra)
	 Develop endpoints to:
	 - Update a book information
	 - Delete a book from the database
	*/
	
	router.Run(":8080")


}