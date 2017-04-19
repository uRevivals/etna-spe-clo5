package main

import (
	"net/http"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Handler struct {
	DB *mgo.Session
}

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

func (h *Handler) AddUser(c echo.Context) (err error) {
	//Bind
	u := &User{ID: bson.NewObjectId()}
	if err = c.Bind(u); err != nil {
		return
	}

	//Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or passowrd"}
	}

	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("users").Insert(u); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) GetUser(c echo.Context) (err error) {
	//Bind
	u := new(User)

	id := c.Param("id")
	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("users").FindId(bson.ObjectIdHex(id)).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "invalid user ID"}
		}
		return
	}

	u.Password = ""
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) FindUser(c echo.Context) (err error) {
	//Bind
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}

	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("users").Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "user not found"}
		}
		return
	}

	u.Password = ""
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) UpdateUser(c echo.Context) (err error) {
	unew := new(User)
	if err = c.Bind(unew); err != nil {
		return
	}

	uold := new(User)
	id := c.Param("id")
	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("users").FindId(bson.ObjectIdHex(id)).One(uold); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "invalid user ID"}
		}
		return
	}

	if err = db.DB("hotel").C("users").Update(bson.M{"_id": bson.ObjectIdHex(id)}, unew); err != nil {
		return
	}
	return c.JSON(http.StatusOK, unew)
}

func (h *Handler) DeleteUser(c echo.Context) (err error) {
	id := c.Param("id")
	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("users").RemoveId(bson.ObjectIdHex(id)); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "invalid email or password"}
		}
		return
	}

	return c.JSON(http.StatusOK, "User Deleted")
}
