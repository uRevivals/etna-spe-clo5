package main

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Billing struct {
		ID    bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		For   bson.ObjectId `json:"userid" bson:"userid,omitempty"`
		Email string        `json:"email" bson:"email"`
		Time  int64         `json:"time,omitempty" bson:"time,omitempty"`
	}
)
