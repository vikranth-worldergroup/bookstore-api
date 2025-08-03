package main

import (
	"bookstore/validators"
	"bookstore/routes"
	"github.com/labstack/echo/v4"
	"bookstore/config"
)

func main(){
	config.Connect()
	e:=echo.New()
	validators.Validates()
	routes.Inroutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
