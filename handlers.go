package main

import (
	"fmt"

	"github.com/takama/router"
)

// home returns the path of current request
func home(c *router.Control) {
	fmt.Fprintf(c.Writer, "Processing URL %s...", c.Request.URL.Path)
}

// logger provides a log of requests
func logger(c *router.Control) {
	remoteAddr := c.Request.Header.Get("X-Forwarded-For")
	if remoteAddr == "" {
		remoteAddr = c.Request.RemoteAddr
	}
	log.Infof("%s %s %s", remoteAddr, c.Request.Method, c.Request.URL.Path)
}
