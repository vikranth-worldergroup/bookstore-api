package controllers

import (
	"bookstore/config"
	"bookstore/models"
	"bookstore/utils"
	"bookstore/validators"
	"strconv"

	// "log"
	"net/http"
	// "fmt"

	"github.com/labstack/echo/v4"
)

// var books=map[string]models.Book{}

func GetBooks(c echo.Context) error{
	page,_:=strconv.Atoi(c.QueryParam("page"))
	limit,_:=strconv.Atoi(c.QueryParam("limit"))

	if page<1{
		page=1
	}
	if limit<1{
		limit=10
	}
	offset:=(page-1)*limit
	book:=[]models.Book{}
	err:=config.Db.Limit(limit).Offset(offset).Find(&book).Error

	if err!=nil{
		return c.JSON(http.StatusInternalServerError,echo.Map{"Error":err.Error()})
	}
	return c.JSON(http.StatusOK,book)
}

func GetBookById(c echo.Context) error{
	id :=c.Param("id")
	var book models.Book

	err:=config.Db.First(&book,id) //returns gorm object
	//Find a record in the books table where the id equals the value from the URL, 
	// and fill the matching row's values into the book variable.
	if err.Error!=nil{
		// fmt.Println(err.Error.Error())
		return c.JSON(http.StatusNotFound,echo.Map{"error":err.Error.Error()}) //err.Error.Error actually converts
																			   //the error to string
	}
	return c.JSON(http.StatusFound,book)
}
func AddBook(c echo.Context) error{
	book:=new(models.Book)
	err:=c.Bind(book) // returns error if the json data is malformed 
	if err!=nil{
		return c.JSON(http.StatusBadRequest,echo.Map{"Error":"Invalid Input"})
	}
	err1 := validators.Val.Struct(book)
	if  err1 != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"validation_error": err1.Error()})
    }
	err2:=config.Db.Create(book).Error
	if err2!=nil{
		return c.JSON(http.StatusBadRequest,echo.Map{"Error":err2.Error()})
	}
	return c.JSON(http.StatusCreated,*book)
}

func UpdateBook(c echo.Context) error{
	id:=c.Param("id")
	book:=new(models.Book)

	err:=config.Db.First(book,id).Error
	if err!=nil{
		return c.JSON(http.StatusBadRequest,echo.Map{"Error":err.Error()})
	}

	err1:=c.Bind(book)
	if err1!=nil{
		return c.JSON(http.StatusBadRequest,echo.Map{"Error":"Invalid Input"})
	}

	err2:=validators.Val.Struct(book)
	if err2!=nil{
		return c.JSON(http.StatusBadRequest,echo.Map{"Error":err2.Error()})
	}

	err3:=config.Db.Save(book).Error
	if err3!=nil{
		return c.JSON(http.StatusInternalServerError,echo.Map{"Error":err3.Error()})
	}

	return c.JSON(http.StatusCreated,*book)

}
func DeleteBook(c echo.Context) error{
	id:=c.Param("id")
	// var book models.Book
	// err:=config.Db.First(&book,id).Error

	// if err!=nil{
	// 	return c.JSON(http.StatusNotFound,echo.Map{"error":err.Error()})
	// }
	err1:=config.Db.Delete(&models.Book{},id).Error // if id doesn't exists no error is returned 
	if err1!=nil{
		return c.JSON(http.StatusInternalServerError,echo.Map{"error":err1.Error()})
	}

	return c.NoContent(http.StatusOK)
}

func Downloadcsv (c echo.Context) error{
	var book []models.Book
	config.Db.Find(&book)
	return utils.Csvexport(book,c)
}

