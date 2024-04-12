package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// basic echo server
	s := echo.New()
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())

	// init jaeger middleware
	closer := jaegertracing.New(s, nil)
	defer closer.Close()

	s.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// start server
	s.Logger.Fatal(s.Start(":8080"))
}
