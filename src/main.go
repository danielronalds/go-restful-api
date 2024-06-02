package main

import (
	"net/http"

	"github.com/danielronalds/go-restful-api/src/resources"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/hello-world", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	e.GET("/events", resources.GetEvents)
	e.GET("/events/:id", resources.GetEventById)
	e.POST("/events", resources.CreateEvent)
	e.PUT("/events", resources.UpdateEvent)
	e.DELETE("/events/:id", resources.DeleteEvent)

	e.Logger.Fatal(e.Start(":3000"))
}
