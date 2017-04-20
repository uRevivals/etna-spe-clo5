package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	dbtesturl string

	h1 = []string{
		`{"ObjectNumber":1,"ObjectUsers":3,"NameType":"S","Pricing":720,"Reservations":[],"hotelid":1}`,
		`{"ObjectNumber":2,"ObjectUsers":2,"NameType":"JS","Pricing":500,"Reservations":[],"hotelid":1}`,
		`{"ObjectNumber":3,"ObjectUsers":2,"NameType":"CD","Pricing":300,"Reservations":[],"hotelid":1}`,
		`{"ObjectNumber":4,"ObjectUsers":2,"NameType":"CS","Pricing":150,"Reservations":[],"hotelid":1}`,
		`{"ObjectNumber":5,"ObjectUsers":2,"NameType":"CS","Pricing":150,"Reservations":[],"hotelid":1}`,
		`{"ObjectNumber":6,"ObjectUsers":1,"NameType":"P","Pricing":25,"Reservations":[],"hotelid":1}`,
		`{"ObjectNumber":7,"ObjectUsers":1,"NameType":"P","Pricing":25,"Reservations":[],"hotelid":1}`,
		`{"ObjectNumber":8,"ObjectUsers":1,"NameType":"P","Pricing":25,"Reservations":[],"hotelid":1}`,
		`{"ObjectNumber":9,"ObjectUsers":1,"NameType":"B","Pricing":0,"Reservations":[],"hotelid":1}`,
		`{"ObjectNumber":10,"ObjectUsers":1,"NameType":"B","Pricing":0,"Reservations":[],"hotelid":1}`,
	}

	h2 = []string{
		`{"ObjectNumber":11,"ObjectUsers":5,"NameType":"SR","Pricing":1000,"Reservations":[],"hotelid":2}`,
		`{"ObjectNumber":12,"ObjectUsers":5,"NameType":"SR","Pricing":1000,"Reservations":[],"hotelid":2}`,
		`{"ObjectNumber":13,"ObjectUsers":1,"NameType":"P","Pricing":25,"Reservations":[],"hotelid":2}`,
		`{"ObjectNumber":14,"ObjectUsers":1,"NameType":"P","Pricing":25,"Reservations":[],"hotelid":2}`,
		`{"ObjectNumber":15,"ObjectUsers":1,"NameType":"B","Pricing":0,"Reservations":[],"hotelid":2}`,
		`{"ObjectNumber":16,"ObjectUsers":1,"NameType":"B","Pricing":0,"Reservations":[],"hotelid":2}`,
	}

	idMap = []bson.ObjectId{}
)

func TestPopulateDb(t *testing.T) {
	//Setup
	e := echo.New()
	for _, element := range h1 {
		req, err := http.NewRequest(echo.POST, "/objects", strings.NewReader(element))
		if assert.NoError(t, err) {
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			//set up Db
			db, err := mgo.Dial("localhost:27017")
			if err != nil {
				panic(err)
			}

			//Initialize handler
			h := &Handler{DB: db}
			if assert.NoError(t, h.AddObject(c)) {
				assert.Equal(t, http.StatusCreated, rec.Code)
				var hObject0 hotelObject
				var hObject1 hotelObject
				json.Unmarshal([]byte(element), &hObject0)
				json.Unmarshal([]byte(rec.Body.String()), &hObject1)
				hObject0.ID = hObject1.ID
				assert.Equal(t, &hObject0, &hObject1)
				idMap = append(idMap, hObject1.ID)
			}
		}
	}

	for _, element := range h2 {
		req, err := http.NewRequest(echo.POST, "/objects", strings.NewReader(element))
		if assert.NoError(t, err) {
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			//set up Db
			db, err := mgo.Dial("localhost:27017")
			if err != nil {
				panic(err)
			}

			//Initialize handler
			h := &Handler{DB: db}
			if assert.NoError(t, h.AddObject(c)) {
				assert.Equal(t, http.StatusCreated, rec.Code)
				var hObject0 hotelObject
				var hObject1 hotelObject
				json.Unmarshal([]byte(element), &hObject0)
				json.Unmarshal([]byte(rec.Body.String()), &hObject1)
				hObject0.ID = hObject1.ID
				assert.Equal(t, &hObject0, &hObject1)
				idMap = append(idMap, hObject1.ID)
			}
		}
	}
}

/*
func TestDeleteDb(t *testing.T) {
	//Setup
	e := echo.New()
	for _, element := range idMap {
		req, err := http.NewRequest(echo.DELETE, "/objects/"+bson.ObjectId.Hex(element), strings.NewReader("{}"))
		if assert.NoError(t, err) {
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/objects/:id")
			c.SetParamNames("id")
			c.SetParamValues(bson.ObjectId.Hex(element))
			//set up Db
			db, err := mgo.Dial("localhost:27017")
			if err != nil {
				panic(err)
			}

			//Initialize handler
			h := &Handler{DB: db}
			if assert.NoError(t, h.DeleteObject(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
			}
		}
	}
}
*/
