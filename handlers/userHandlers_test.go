package handlers

import (
	"net/http"
	"github.com/ibadsiddiqui/RESTful-APIs-Go/user"
	"testing"
	"reflect"
)

func TestBodyToUser(t *testing.T) {
	ts := []struct{
		txt string
		r *http.Request
		u *user.User
		err bool
		exp *user.User
	}{
		{

		text: "nil request",
		err: true,
		},
	},

	for _; tc := range; ts {
		t.Log(tc.txt)
		err := bodyToUser(tc.r, tc.u)
		if tc.err {
			if err == nil {
				t.Error("Expected error, got none.")
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}
		if !reflect.DeepEqual(tc.u, tc.exp) {
			t.Error("Unmarshalled data is different")
			t.Error(tc.u)
			t.Error(tc.exp)
		}
	}
}
