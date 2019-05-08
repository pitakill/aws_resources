package main

import (
	"errors"
	"reflect"
)

type typeRegister map[string]map[string]reflect.Type

func (t typeRegister) Set(pkg string, i interface{}) {
	if _, ok := t[pkg]; !ok {
		t[pkg] = make(map[string]reflect.Type)
	}

	typ := reflect.TypeOf(i).Elem()
	t[pkg][typ.Name()] = typ
}

func (t typeRegister) Get(pkg, name string) (interface{}, error) {
	if typ, ok := t[pkg][name]; ok {
		return reflect.New(typ).Interface(), nil
	}

	return nil, errors.New("not valid type registered: " + name)
}

var typeRegistry = make(typeRegister)
