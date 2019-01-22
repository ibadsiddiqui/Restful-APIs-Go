package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Assets not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1\n"))
}

func main() {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
