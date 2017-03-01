package router

import (
	"github.com/YoungZhou93/startGolang/web/webframe/controller"
	"net/http"
)

type Router interface {

	Register(pattern string, c controller.Controller)

	ServeForward(w http.ResponseWriter, r *http.Request)
}
