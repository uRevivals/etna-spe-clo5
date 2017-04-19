package billing

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Billing struct {
		ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
		For   bson.ObjectId
		Email string `json:"email" bson:"email"`
		Time  int64
	}
)
