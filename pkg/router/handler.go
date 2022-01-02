package router

import "net/http"

type Context struct {
	params []string
}

type Handler func(http.ResponseWriter, *http.Request, *Context)
