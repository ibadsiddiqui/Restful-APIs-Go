package handlers

import "net/http"

func UsersRouter(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API at users\n"))
}
