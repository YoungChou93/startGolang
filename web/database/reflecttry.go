package main

import (
	"reflect"
	"fmt"
	"time"
)

type Person struct {
	Name string
	sex bool
	birthday time.Time
}

func NewPerson() Person{
	return Person{}
}

func orm(arg interface{}){
	t:=reflect.TypeOf(arg)
	//v:=reflect.ValueOf(*request)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		fmt.Println(t.Elem().Name())
	}

}

func main()  {

        p:=NewPerson()
	//orm(p)
	t:=reflect.TypeOf(p)

	fmt.Println(t.Name())
	/*fmt.Println(v.NumField())
	vf:=v.Field(1)
	fv:=v.Field(1).Interface()
	fmt.Println(vf)
	fmt.Println(fv)*/
}