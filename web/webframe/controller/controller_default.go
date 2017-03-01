package controller

import (
	"html/template"
	"net/http"
)

type DefaultController struct {
	Ct *Context
	Tpl *template.Template
	Data map[interface{}]interface{}
	ChildName string
	TplNames string
	Layout []string
	TplExt string
}

func (this * DefaultController) Init(ct *Context, cn string){
	this.Data = make(map[interface{}]interface{})
	this.Layout = make([]string, 0)
	this.TplNames = ""
	this.ChildName = cn
	this.Ct = ct
	this.TplExt = "tpl"
}
func (this * DefaultController) Get(){
	http.Error(this.Ct.ResponseWriter, "Method Not Allowed", 405)
}

