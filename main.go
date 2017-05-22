package main

import (
	"github.com/takama/router"
)

// Run server: go build && step-by-step
// Try requests: curl http://127.0.0.1:8000/test
func main() {
	r := router.New()
	r.GET("/", home)
	r.Listen(":8000")
}
