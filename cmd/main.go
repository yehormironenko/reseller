package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"reseller/controller/handlers"
)

// 2 function search bookByName // GET with params
// Buy
func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())

	e.GET("/echo", handlers.Echo)
	e.GET("/search", handlers.GetBook)

	e.Logger.Fatal(e.Start(":1323"))
}
