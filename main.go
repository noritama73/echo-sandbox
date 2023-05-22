package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/noritama73/echo-sandbox/internal/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())

	handler := handler.NewApiHandler()
	e.GET("/add", handler.Add)

	e.Logger.Fatal(e.Start(":8888"))
}
