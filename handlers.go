package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type book struct{
	ID string `json:"id"`
    Title string `json:"title"`
    Author string `json:"author"`
    Price float64 `json:"price"`
    Content string `json:"content"`  
}

var books=map[string]book{}

func GetBooks(c echo.Context) error{
	bookslist:=[]book{}

	for _,b:= range books{
		bookslist = append(bookslist, b)
	}

	return c.JSON(http.StatusOK,bookslist)
}

func GetBookById(c echo.Context) error{
	id :=c.Param("id")
	book,ok := books[id]

	if !ok{
		return c.JSON(http.StatusNotFound,echo.Map{"error":"Book not found"})
	}

	return c.JSON(http.StatusFound,book)
}
func AddBook(c echo.Context) error{
	p:=new(book)
	err:=c.Bind(p)
	if p.ID == "" {
	return c.JSON(http.StatusBadRequest, echo.Map{"error": "Book ID cannot be empty"})
	}
	if err!=nil{
		return c.JSON(http.StatusBadRequest,echo.Map{"Error":"Invalid Input"})
	}
	_,ok:=books[p.ID]
	if ok{
		return c.JSON(http.StatusBadRequest,echo.Map{"Error":"Book already exists"})
	}
	books[p.ID]=*p
	return c.JSON(http.StatusCreated,p)
}

func UpdateBook(c echo.Context) error{
	id:=c.Param("id")
	_,ok:=books[id]
	if !ok{
		return c.JSON(http.StatusBadRequest,echo.Map{"Error":"Book not found"})
	}

	p:=new(book)
	err:=c.Bind(p)
	if err!=nil{
		return c.JSON(http.StatusBadRequest,echo.Map{"Error":"Invalid Input"})
	}
	
	books[id]=*p
	return c.JSON(http.StatusCreated,p)

}
func DeleteBook(c echo.Context) error{
	id:=c.Param("id")
	_,ok:=books[id]

	if !ok{
		return c.JSON(http.StatusNotFound,echo.Map{"error":"Book not found"})
	}

	delete(books,id)
	return c.NoContent(http.StatusOK)
}

