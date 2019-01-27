package handlers

import (
	"net/http"

	"github.com/ibadsiddiqui/RESTful-APIs-Go/user"
)

func usersGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})

}
