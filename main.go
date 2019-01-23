package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ibadsiddiqui/RESTful-APIs-Go/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
