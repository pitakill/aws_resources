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
	"github.com/aws/aws-sdk-go-v2/service/snowball"
)

type SnowballType struct {
	service      *snowball.Snowball
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type SnowballTypeConfig struct {
	resourceType string
}

func SnowballFactory(cfg aws.Config) Factory {
	i := new(SnowballType)

	i.SetService(cfg)

	return i
}

func (i *SnowballType) SetPartialName() {
	// "AWS::Snowball::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::Snowball::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *SnowballType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *SnowballType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *SnowballType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *SnowballType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *SnowballType) Configure(param interface{}) error {
	config, ok := param.(SnowballTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (SnowballTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *SnowballType) SetService(cfg aws.Config) {
	srv := snowball.New(cfg)

	i.service = srv
}

func (i *SnowballType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("snowball", i.inputName)
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

func (i *SnowballType) GetResources() {}

func (i *SnowballType) GetResourcesDetail() {}
