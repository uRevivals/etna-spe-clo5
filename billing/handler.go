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

func (h *Handler) EstimateBilling(c echo.Context) (err error) {
	//Bind
	item := &Billing{ID: bson.NewObjectId()}
	if err = c.Bind(Billing); err != nil {
		return
	}

	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("hotelobject").Insert(item); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, item)
}

func (h *Handler) GetObject(c echo.Context) (err error) {
	//Bind
	item := new(hotelObject)

	id := c.Param("id")
	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("hotelobject").FindId(bson.ObjectIdHex(id)).One(item); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "invalid hotelObject Id"}
		}
		return
	}

	return c.JSON(http.StatusOK, item)
}

func (h *Handler) FindType(c echo.Context) (err error) {
	//Bind
	nametype := c.Param("nametype")
	var items []hotelObject
	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("hotelobject").Find(bson.M{"nametype": nametype}).All(items); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Empty result set"}
		}
		return
	}

	return c.JSON(http.StatusOK, items)
}

func (h *Handler) UpdateObject(c echo.Context) (err error) {
	itemnew := new(hotelObject)
	if err = c.Bind(itemnew); err != nil {
		return
	}

	itemold := new(hotelObject)
	id := c.Param("id")
	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("hotelobject").FindId(bson.ObjectIdHex(id)).One(itemold); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "invalid object Id"}
		}
		return
	}

	if err = db.DB("hotel").C("hotelobject").Update(bson.M{"_id": bson.ObjectIdHex(id)}, itemnew); err != nil {
		return
	}
	itemnew.ID = itemold.ID
	return c.JSON(http.StatusOK, itemnew)
}

func (h *Handler) DeleteObject(c echo.Context) (err error) {
	id := c.Param("id")

	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("hotelobject").RemoveId(bson.ObjectIdHex(id)); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "invalid id"}
		}
		return
	}

	return c.JSON(http.StatusOK, "Object Deleted")
}
