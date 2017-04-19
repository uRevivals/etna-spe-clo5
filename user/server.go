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

	//Create indices
	if err = db.Copy().DB("hotel").C("users").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	//Initialize handler
	h := &Handler{DB: db}

	//Routes
	e.POST("/users", h.AddUser)
	e.PUT("/users/:id", h.UpdateUser)
	e.GET("/users/:id", h.GetUser)
	e.DELETE("/users/:id", h.DeleteUser)
	e.POST("/users/find", h.FindUser)

	e.Logger.Fatal(e.Start(":" + port))
}
