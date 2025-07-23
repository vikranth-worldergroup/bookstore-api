package main

import (
	"github.com/labstack/echo/v4"
)

func main(){
	e:=echo.New()
	e.GET("/books",GetBooks)
	e.GET("/books/:id",GetBookById)
	e.POST("/books",AddBook)
	e.PUT("/books/:id",UpdateBook)
	e.DELETE("/books/:id",DeleteBook)

	e.Logger.Fatal(e.Start(":8080"))
}
