// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-09 11:54:14.958521953 -0500 CDT m=+0.000134534
package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
)

type Route53ResolverType struct {
	service      *route53resolver.Route53Resolver
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type Route53ResolverTypeConfig struct {
	resourceType string
}

func Route53ResolverFactory(cfg aws.Config) Factory {
	i := new(Route53ResolverType)

	i.SetService(cfg)

	return i
}

func (i *Route53ResolverType) SetPartialName() {
	// "AWS::Route53Resolver::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::Route53Resolver::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *Route53ResolverType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *Route53ResolverType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *Route53ResolverType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *Route53ResolverType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *Route53ResolverType) Configure(param interface{}) error {
	config, ok := param.(Route53ResolverTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (Route53ResolverTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *Route53ResolverType) SetService(cfg aws.Config) {
	srv := route53resolver.New(cfg)

	i.service = srv
}

func (i *Route53ResolverType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("route53resolver", i.inputName)
	if err != nil {
		// We can ignore this kind of errors because there is not resources by the
		// i.inputName
		//log.Println(err)
		return
	}

	method := reflect.ValueOf(i.service).MethodByName(i.methodName)
	called := method.Call([]reflect.Value{reflect.ValueOf(instance)})

	send := reflect.Indirect(called[0]).MethodByName("Send")
	calledSend := send.Call([]reflect.Value{})

	res := calledSend[0]

	fmt.Printf("%v\n", res)
}

func (i *Route53ResolverType) GetResources() {}

func (i *Route53ResolverType) GetResourcesDetail() {}
