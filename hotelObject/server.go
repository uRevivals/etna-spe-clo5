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
		port = "1234"
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
	h := &Handler{DB: db}

	//Routes
	e.POST("/objects", h.AddObject)
	e.PUT("/objects/:id", h.UpdateObject)
	e.GET("/objects/:id", h.GetObject)
	e.DELETE("/objects/:id", h.DeleteObject)
	e.POST("/objects/find", h.FindType)
	e.POST("/objects/checkreservation/", h.CheckReservation)
	e.POST("/objects/reservation/", h.ValidateReservation)

	e.Logger.Fatal(e.Start(":" + port))
}
