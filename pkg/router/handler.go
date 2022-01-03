package router

import "net/http"

type Context struct {
	Params map[string]string
}

type Handler func(http.ResponseWriter, *http.Request, *Context)
