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

type reservations struct {
	ID               bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty"`
	ReservedDates    []int64         `json:"ReservedDates" bson:"ReservedDates"`
	ConcernedObjects []bson.ObjectId `json:"ConcernedObjects" bson:"ConcernedObjects"`
	UserID           bson.ObjectId   `json:"userid" bson:"userid"`
}

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

func (h *Handler) AddObject(c echo.Context) (err error) {
	//Bind
	item := &hotelObject{ID: bson.NewObjectId()}
	if err = c.Bind(item); err != nil {
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

func (h *Handler) CheckReservation(c echo.Context) (err error) {
	protoReservation := new(reservations)
	item := new(hotelObject)
	if err = c.Bind(protoReservation); err != nil {
		return
	}

	db := h.DB.Clone()
	defer db.Close()
	for _, element := range protoReservation.ConcernedObjects {
		if err == db.DB("hotel").C("hotelobject").FindId(element).One(item) {
			for _, date := range protoReservation.ReservedDates {
				for _, alreadyReserved := range item.Reservations {
					if date == alreadyReserved {
						return c.JSON(http.StatusBadRequest, "Object :"+bson.ObjectId.Hex(element)+" Already reserved for one ore less date")
					}
				}
			}
		}
	}

	return c.JSON(http.StatusOK, "Reservation Validate")
}

func (h *Handler) ValidateReservation(c echo.Context) (err error) {
	protoReservation := new(reservations)
	item := new(hotelObject)
	if err = c.Bind(protoReservation); err != nil {
		return
	}

	db := h.DB.Clone()
	defer db.Close()
	for _, element := range protoReservation.ConcernedObjects {
		if err = db.DB("hotel").C("hotelobject").FindId(element).One(item); err != nil {
			if err == mgo.ErrNotFound {
				return &echo.HTTPError{Code: http.StatusNotFound, Message: "invalid object Id"}
			}
		}
		for _, date := range protoReservation.ReservedDates {
			for _, alreadyReserved := range item.Reservations {
				if date == alreadyReserved {
					return c.JSON(http.StatusBadRequest, "Object :"+bson.ObjectId.Hex(element)+" Has benn reserved for one ore less date")
				}
			}
			item.Reservations = append(item.Reservations, date)
		}
		if err = db.DB("hotel").C("hotelobject").Update(bson.M{"_id": element}, item); err != nil {
			return
		}
	}
	return c.JSON(http.StatusOK, "Object Validate")
}
