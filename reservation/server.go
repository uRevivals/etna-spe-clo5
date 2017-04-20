package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	// In microservice port depends of Env variable
	port := os.Getenv("APP_PORT")
	// If none is defined we start on default one
	if port == "" {
		port = "8080"
	}

	//set Up framework
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	//set up Db
	db, err := mgo.Dial("localhost:27017")
	if err != nil {
		e.Logger.Fatal(err)
	}

	//Initialize handler
	h := &Handler{
		DB:       db,
		hotelURL: "http://localhost:1234",
		userURL:  "http://localhost:1235",
	}

	//Routes
	e.POST("/reservations", h.CreateReservation)
	e.DELETE("/reservations/:id", h.DeleteObject)
	e.POST("/reservations/find", h.FindReservation)

	e.Logger.Fatal(e.Start(":" + port))
}
