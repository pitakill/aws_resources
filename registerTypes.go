package main

import (
	"errors"
	"reflect"
	"runtime"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type typeRegister map[string]reflect.Type

func (t typeRegister) Set(i interface{}) {
	typ := reflect.TypeOf(i).Elem()
	t[typ.Name()] = typ
}

func (t typeRegister) Get(name string) (interface{}, error) {
	if typ, ok := t[name]; ok {
		return reflect.New(typ).Elem().Interface(), nil
	}

	return nil, errors.New("not valid type registered: " + name)
}

var typeRegistry = make(typeRegister)

func init() {
	typeRegistry.Set(new(ec2.DescribeInternetGatewaysInput))
	typeRegistry.Set(new(ec2.DescribeSecurityGroupsInput))
	typeRegistry.Set(new(ec2.DescribeInstancesInput))
	typeRegistry.Set(new(ec2.DescribeVpcsInput))
	typeRegistry.Set(new(ec2.DescribeRouteTablesInput))
	typeRegistry.Set(new(ec2.DescribeSubnetsInput))
	runtime.GC()
}
