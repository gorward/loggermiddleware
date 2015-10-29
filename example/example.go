package main

import (
	"fmt"
	"github.com/gorward/loggermiddleware"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", loggermiddleware.LoggerMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "helllo")
		})))
}
