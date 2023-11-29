package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	repo := CreateRepository()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.GET("/attacks", func(c echo.Context) error {
		return c.JSON(http.StatusOK, repo.FindAll())
	})
	e.Logger.Fatal(e.Start(":1323"))
}
