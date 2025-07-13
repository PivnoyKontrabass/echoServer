package main

import (
	"echoServerJson/cfg"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", GetHandler, Middleware)
	e.POST("/tasks", GetHandler)
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
	var task cfg.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	if err := c.Validate(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	task.ID = "task_" + uuid.New().String()
	task.Metadata.CreatedAt = time.Now()
	return c.JSON(http.StatusCreated, task)
}
