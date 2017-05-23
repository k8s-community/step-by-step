package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/takama/router"
)

var log = logrus.New()

// Run server: go build; env SERVICE_PORT=8000 step-by-step
// Try requests: curl http://127.0.0.1:8000/test
func main() {
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Required parameter service port is not set")
	}

	r := router.New()
	r.Logger = logger
	r.GET("/", home)
	r.Listen("0.0.0.0:" + port)
}
