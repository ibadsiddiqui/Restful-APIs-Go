package user

import (
	"errors"

	"gihub.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID		bson.ObjectId 	`json:"id" storm:"id"`
	Name 	string			`json:"name"`
	Role 	string			`json:"role"`
}

const dbPath = "users.db"

var ErrRecordInvalid = error.New("Record is invalid")