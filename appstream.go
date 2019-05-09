// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-09 13:01:20.611969172 -0500 CDT m=+0.000147082
package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
)

type AppStreamType struct {
	service      *appstream.AppStream
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type AppStreamTypeConfig struct {
	resourceType string
}

func AppStreamFactory(cfg aws.Config) Factory {
	i := new(AppStreamType)

	i.SetService(cfg)

	return i
}

func (i *AppStreamType) SetPartialName() {
	// "AWS::AppStream::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::AppStream::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *AppStreamType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *AppStreamType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *AppStreamType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *AppStreamType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *AppStreamType) Configure(param interface{}) error {
	config, ok := param.(AppStreamTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (AppStreamTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *AppStreamType) SetService(cfg aws.Config) {
	srv := appstream.New(cfg)

	i.service = srv
}

func (i *AppStreamType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("appstream", i.inputName)
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

func (i *AppStreamType) GetResources() {}

func (i *AppStreamType) GetResourcesDetail() {}