package controllers

import (
	"bookstore/config"
	"bookstore/models"
	"bookstore/utils"
	"bookstore/validators"
	"strconv"
	"net/http"
	"github.com/labstack/echo/v4"
)

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
	err:=config.Db.Preload("Authors").Limit(limit).Offset(offset).Find(&book).Error

	if err!=nil{
		return c.JSON(http.StatusInternalServerError,echo.Map{"Error":err.Error()})
	}
	return c.JSON(http.StatusOK,book)
}

func GetBookById(c echo.Context) error{
	id :=c.Param("id")
	var book models.Book

	err:=config.Db.Preload("Authors").First(&book,id) 
	if err.Error!=nil{
		return c.JSON(http.StatusNotFound,echo.Map{"error":err.Error.Error()}) 
	}

	return c.JSON(http.StatusFound,book)
}
func AddBook(c echo.Context) error{
	book:=new(models.Book)

	err:=c.Bind(book) 
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

	upbook:=new(models.Book)
	upbook.ID=book.ID
	upbook.Title=book.Title
	upbook.Price=book.Price
	upbook.Content=book.Content

	t:=config.Db.Begin()
	err3:=t.Where("book_id=?",id).Delete(&models.Author{}).Error
	if err3!=nil{
		t.Rollback()
		return c.JSON(http.StatusInternalServerError,echo.Map{"Error":err3.Error()})
	}

	for i := range book.Authors{
		book.Authors[i].BookID=id
	}

	upbook.Authors=book.Authors

	err4:=t.Save(upbook).Error
	if err4!=nil{
		t.Rollback()
		return c.JSON(http.StatusInternalServerError,echo.Map{"Error":err4.Error()})
	}

	t.Commit()

	return c.JSON(http.StatusOK,*book)

}
func DeleteBook(c echo.Context) error{
	id:=c.Param("id")

	err1:=config.Db.Where("book_id=?",id).Delete(&models.Author{}).Error
	if err1!=nil{
		return c.JSON(http.StatusInternalServerError,echo.Map{"error":err1.Error()})
	}

	err2:=config.Db.Delete(&models.Book{},id)
	if err2.Error!=nil{
		return c.JSON(http.StatusInternalServerError,echo.Map{"error":err2.Error.Error()})
	}

	if err2.RowsAffected==0{
		return c.JSON(http.StatusNotFound,echo.Map{"Error":err2.Error.Error()})
	}

	return c.NoContent(http.StatusOK)
}

func Downloadcsv (c echo.Context) error{
	var book []models.Book
	config.Db.Preload("Authors").Find(&book)
	return utils.Csvexport(book,c)
}

