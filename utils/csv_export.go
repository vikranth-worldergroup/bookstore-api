package utils

import (
	"bookstore/models"
	"encoding/csv"
	"fmt"
	"strings"
	"github.com/labstack/echo/v4"
)

func Csvexport(book []models.Book,c echo.Context) error{
	c.Response().Header().Set(echo.HeaderContentType,"text/csv")
	c.Response().Header().Set(echo.HeaderContentDisposition,"attachment,filename:Books.csv")

	cswriter:=csv.NewWriter(c.Response())
	cswriter.Write([]string{"ID","Title","Author","Price","Content"})

	for _,b := range book{
		
		authjoined:=[]string{}
		for _,i := range b.Authors{
			authjoined=append(authjoined,i.Name)
		}
		authnames:=strings.Join(authjoined,", ")
		
		cswriter.Write([]string{b.ID,b.Title,authnames,fmt.Sprintf("%.2f",b.Price),b.Content})
	}

	cswriter.Flush()
	return nil
}