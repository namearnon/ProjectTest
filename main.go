package main

import (
	"ProjectTest/middlewares"
	"ProjectTest/route"
	"errors"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middlewares.Core)
	e.Use(middleware.Static("./"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowOrigins:     []string{"*"},
	}))
	route.TestRoute(e)
	if _, err := os.Stat("./images"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("./images", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	var port = os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	log.Fatal(e.Start(":" + port))
}
