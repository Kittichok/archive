package main

import (
	"beefbeef/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "beef beef")
	})

	client := &http.Client{}
	beefService := services.NewBeefService(client)

	e.GET("/beef/summary", func(c echo.Context) error {
		return c.JSON(200, beefService.Count())
	})

	e.Logger.Fatal(e.Start(":1323"))
}
