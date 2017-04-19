package main

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	reservations struct {
		ID               bson.ObjectId   `json:"id" bson:"_id,omitempty"`
		ReservedDates    []int64         `json:"ReservedDates" bson:"ReservedDates"`
		ConcernedObjects []bson.ObjectId `json:"ConcernedObjects" bson:"ConcernedObjects"`
		UserID           bson.ObjectId   `json:"userid" bson:"userid"`
	}
)
