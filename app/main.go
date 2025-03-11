package main

import (
	"net/http"

	"github.com/Melidee/goth-chat/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	app.Logger.Fatal(app.Start(":8080"))
}