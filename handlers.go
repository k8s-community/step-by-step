package main

import (
	"fmt"

	"github.com/k8s-community/step-by-step/version"
	"github.com/takama/router"
)

// home returns the path of current request
func home(c *router.Control) {
	fmt.Fprintf(c.Writer, "Repo: %s, Commit: %s, Version: %s", version.REPO, version.COMMIT, version.RELEASE)
}

// logger provides a log of requests
func logger(c *router.Control) {
	remoteAddr := c.Request.Header.Get("X-Forwarded-For")
	if remoteAddr == "" {
		remoteAddr = c.Request.RemoteAddr
	}
	log.Infof("%s %s %s", remoteAddr, c.Request.Method, c.Request.URL.Path)
}
