package main

import (
	"net/http"
	"strings"

	"encoding/json"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Handler struct {
	DB       *mgo.Session
	hotelURL string
	userURL  string
}

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

func (h *Handler) CreateReservation(c echo.Context) (err error) {
	//Bind
	r := &reservations{ID: bson.NewObjectId()}
	if err = c.Bind(r); err != nil {
		return
	}

	//Check validity of request
	resp, err := http.Get(h.userURL + "/users/" + bson.ObjectId.Hex(r.UserID))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return c.JSON(http.StatusBadRequest, "invalid User id")
	}

	client := &http.Client{}
	json, err := json.Marshal(r)

	req, err := http.NewRequest(echo.POST, h.hotelURL+"/objects/checkreservation/", strings.NewReader(string(json)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return c.JSON(http.StatusBadRequest, "invalid Object")
	}

	req, err = http.NewRequest(echo.POST, h.hotelURL+"/objects/reservation/", strings.NewReader(string(json)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return c.JSON(http.StatusBadRequest, "invalid Object Reservation")
	}

	db := h.DB.Clone()
	defer db.Close()
	if resp.StatusCode == 200 {
		if err = db.DB("hotel").C("reservation").Insert(r); err != nil {
			return
		}
	}

	return c.JSON(http.StatusCreated, r)
}

func (h *Handler) FindReservation(c echo.Context) (err error) {
	//Bind
	var items []reservations
	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("reservation").Find(bson.M{}).All(items); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "Empty result set"}
		}
		return
	}

	return c.JSON(http.StatusOK, items)
}

func (h *Handler) DeleteObject(c echo.Context) (err error) {
	id := c.Param("id")

	//Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("hotel").C("reservation").RemoveId(bson.ObjectIdHex(id)); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "invalid id"}
		}
		return
	}

	return c.JSON(http.StatusOK, "Object Deleted")
}
