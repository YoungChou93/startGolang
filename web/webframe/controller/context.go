package controller

import "net/http"

type Context struct {
	ResponseWriter http.ResponseWriter
	Request *http.Request
	Params map[int]string
}