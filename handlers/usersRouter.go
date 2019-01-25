package handlers

import (
	"net/http"
	"strings"
)

func UsersRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			return
		case http.MethodPost:
			return
		// case http.MethodPut:
		// 	return
		// case http.MethodPatch:
		// 	return
		default:
			postError(w, http.StatusMethodNotAllowed)
		}
	}
}
