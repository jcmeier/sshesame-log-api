package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	mongodbURI, exists := os.LookupEnv("MONGODB_CONNECTION_STRING")
	if !exists {
		panic("Env variable MONGODB_CONNECTION_STRING needs to be set")
	}

	repo := CreateRepository(mongodbURI)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.GET("/attacks", func(c echo.Context) error {
		return c.JSON(http.StatusOK, repo.FindAll())
	})
	e.Logger.Fatal(e.Start(":1323"))
}
