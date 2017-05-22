# step-by-step

An example of a service, which "evolves" step by step.

## Step 01. Preparation

It might be helpful to set up git pre-commit hooks before writing the code.
For example, you can use [pre-commit by Yelp](http://pre-commit.com) to manage this kind of hooks.

See `.pre-commit-config.yaml`.

## Step 02. Add the simplest web-server

The simplest way to make a web-server: register a handler function and listen on some TCP network address:
 
    http.HandleFunc("/", home)
    http.ListenAndServe(":8000", nil)

## Step 03. Add tests

`TestHandler` function (see `handlers_test.go`) is the simplest way to test our web-server: check if the handler
returns expected result. Use `go test -v` command to run test.
