package main

import (
	"fmt"

	"github.com/takama/router"
)

// home returns the path of current request
func home(c *router.Control) {
	fmt.Fprintf(c.Writer, "Processing URL %s...", c.Request.URL.Path)
}
