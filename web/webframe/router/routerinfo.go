package router

import (
	"regexp"
	"reflect"
)

type RouterInfo struct {
	url            string
	regex          *regexp.Regexp
	params         map[int]string
	routerType reflect.Type
}
