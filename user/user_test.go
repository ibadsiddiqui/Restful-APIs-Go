package user

import (
	"reflect"
	"testing"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

func TestCRUD(t *testing.T) {
	t.Log("CREATE")
	u := &User{
		ID:   bson.NewObjectId(),
		Name: "John",
		Role: "Tester",
	}
	err := u.Save()

	if err != nil {
		t.Fatalf("Error saving a record %s", err)
	}

	t.Log("READ")
	u2, err := One(u.ID)
	if err != nil {
		t.Fatalf("Error retrieving a record %s", err)
	}
	if !reflect.DeepEqual(u2, u) {
		t.Error("Records do not match")
	}

	t.Log("UPDATE")
	u.Role = "developer"
	err = u.Save()
	if err != nil {
		t.Fatalf("Error saving a record %s", err)
	}
	u3, err := One(u.ID)
	if err != nil {
		t.Fatalf("Error retrieving a record %s", err)
	}
	if !reflect.DeepEqual(u3, u) {
		t.Error("Records do not match")
	}

	t.Log("DELETE")
	err = Delete(u.ID)
	if err != nil {
		t.Fatalf("Error deleting a record %s", err)
	}

	_, err = One(u.ID)
	if err == nil {
		t.Fatal("Records should not exist anymore ")
	}
	if err != storm.ErrNotFound {
		t.Fatalf("Error in retrieving non-existing record: %s ", err)

	}

	t.Log("READ ALL")
	u2.ID = bson.NewObjectId()
	u3.ID = bson.NewObjectId()
	err = u2.Save()

	if err != nil {
		t.Fatalf("Erro in retrieving non-existing record: %s ", err)
	}
	err = u3.Save()

	if err != nil {
		t.Fatalf("Erro in retrieving non-existing record: %s ", err)
	}

	users, err := All()
	if len(users) != 3 {
		t.Errorf("Different number of records retrieved. Expected: 3 / Actual : %d ", len(users))

	}
}
