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
	"github.com/aws/aws-sdk-go-v2/service/dax"
)

type DAXType struct {
	service      *dax.DAX
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type DAXTypeConfig struct {
	resourceType string
}

func DAXFactory(cfg aws.Config) Factory {
	i := new(DAXType)

	i.SetService(cfg)

	return i
}

func (i *DAXType) SetPartialName() {
	// "AWS::DAX::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::DAX::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *DAXType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *DAXType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *DAXType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *DAXType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *DAXType) Configure(param interface{}) error {
	config, ok := param.(DAXTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (DAXTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *DAXType) SetService(cfg aws.Config) {
	srv := dax.New(cfg)

	i.service = srv
}

func (i *DAXType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("dax", i.inputName)
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

func (i *DAXType) GetResources() {}

func (i *DAXType) GetResourcesDetail() {}
