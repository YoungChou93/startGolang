package router

import (
	"github.com/YoungZhou93/startGolang/web/crawler/controller"
	"net/http"
	"reflect"
	"strings"
)

type DefaultRouter struct {
	routers []*RouterInfo
	App     *Application
}

func (p *DefaultRouter) Register(pattern string, c controller.Controller) {

	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &RouterInfo{}
	route.url = pattern
	route.routerType = t
	p.routers = append(p.routers, route)
}

func (p *DefaultRouter) ServeForward(w http.ResponseWriter, r *http.Request) {

	var started bool
	for prefix, staticDir := range p.App.StaticDir {
		if strings.HasPrefix(r.URL.Path, prefix) {
			file := staticDir + r.URL.Path[len(prefix):]
			http.ServeFile(w, r, file)
			started = true
			return
		}
	}

	requestPath := r.URL.Path
	for _, route := range p.routers {
		if requestPath == route.url {
			vc := reflect.New(route.routerType)
			method := vc.MethodByName("Get")
			method.Call(nil)
		}

	}

	if started == false {
		http.NotFound(w, r)
	}

}
