// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-24 10:15:27.630300518 -0500 CDT m=+0.000216236
package main

import (
	"errors"
	"fmt"
	"log"
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

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsInput or ListVpcsInput" accordingly
	prefix := search[r{resource: "route53resolver", kind: partialName}]
	name := fmt.Sprintf("%s%sInput", prefix, partialName)

	i.inputName = name
}

func (i *Route53ResolverType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsOutput or ListVpcsOutput" accordingly
	prefix := search[r{resource: "route53resolver", kind: partialName}]
	name := fmt.Sprintf("%s%sOutput", prefix, partialName)

	i.outputName = name
}

func (i *Route53ResolverType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsRequest or ListVpcsRequest" accordingly
	prefix := search[r{resource: "route53resolver", kind: partialName}]
	name := fmt.Sprintf("%s%sRequest", prefix, partialName)

	i.methodName = name
}

func (i *Route53ResolverType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *Route53ResolverType) Configure(param interface{}) error {
	config, ok := param.(TypeConfig)
	if !ok {
		return errors.New("config is not a valid param (TypeConfig)")
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
		log.Println(err)
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
