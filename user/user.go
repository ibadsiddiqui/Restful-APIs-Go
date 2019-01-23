package user

import (
	// "errors"

	"gihub.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}

const dbPath = "users.db"

var ErrRecordInvalid = error.New("Record is invalid")

// All retrieves all users from the database
func All() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// One returns a single user record from the databae
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}

	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}