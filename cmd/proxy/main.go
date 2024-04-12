package main

import (
	"log"
	"net/url"

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

	// init proxy middleware
	urlString := []string{
		"http://server1:8080",
		"http://server2:8080",
	}
	urls, err := parseUrls(urlString)
	if err != nil {
		log.Fatal(err)
	}
	s.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(urls)))

	// start server
	s.Logger.Fatal(s.Start(":8080"))
}

func parseUrls(urls []string) ([]*middleware.ProxyTarget, error) {
	var targets []*middleware.ProxyTarget
	for _, u := range urls {
		u, err := url.Parse(u)
		if err != nil {
			return nil, err
		}
		targets = append(targets, &middleware.ProxyTarget{URL: u})
	}
	return targets, nil
}
