package controller

type Controller interface {
	Get()
	Init(ct *Context, cn string)
}
