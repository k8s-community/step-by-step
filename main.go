package main

import (
	"net/http"
)

// Run server: go build && step-by-step
// Try requests: curl http://127.0.0.1:8000/test
func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}
