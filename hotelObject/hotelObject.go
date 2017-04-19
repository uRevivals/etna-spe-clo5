package main

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	hotelObject struct {
		ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
		ObjectNumber int64         `json:"objectnumber,omitempty" bson:"objectnumber,omitempty"`
		Objectusers  int64         `json:"objectusers,omitempty" bons:"objectusers,omitempty"`
		NameType     string        `json:"nametype" bson:"nametype"`
		Pricing      float64       `json:"pricing" bson:"pricing"`
		Reservations []int64       `json:"reservations" bson:"reservations"`
		Hotel        int64         `json:"hotelid" bson:"hotelid"`
	}
)
