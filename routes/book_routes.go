package routes

import (
	"bookstore/controllers"
	"github.com/labstack/echo/v4"
)

func Inroutes(e *echo.Echo){
	e.PUT("/books/:id",controllers.UpdateBook)
	e.GET("/books",controllers.GetBooks)
	e.GET("/books/:id",controllers.GetBookById)
	e.POST("/books",controllers.AddBook)
	e.GET("/books/download",controllers.Downloadcsv)
	e.DELETE("/books/:id",controllers.DeleteBook)
}