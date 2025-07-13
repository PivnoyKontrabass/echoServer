package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Todo:сделать ручку на метод POST которая бы принимала тело запроса и приводила через это добро в структуру
func main() {
	e := echo.New()
	e.GET("/", GetHandler, Middleware)
	e.POST("/", GetHandler)
	e.Logger.Fatal((e.Start(":8080")))
}
func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		value := c.QueryParam("value")
		if value == "" {
			return c.String(http.StatusBadRequest, "no value given")
		}
		c.Set("value", value)
		return next(c)
	}
}

func GetHandler(c echo.Context) error {
	value, ok := c.Get("value").(string)
	if !ok {
		return c.String(http.StatusInternalServerError, "error: no value found in context")

	}
	return c.String(http.StatusOK, "given value:"+value)
}

func PostHandler(c echo.Context) error {
	return nil
}
