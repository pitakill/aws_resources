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
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
)

type ACMPCAType struct {
	service      *acmpca.ACMPCA
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type ACMPCATypeConfig struct {
	resourceType string
}

func ACMPCAFactory(cfg aws.Config) Factory {
	i := new(ACMPCAType)

	i.SetService(cfg)

	return i
}

func (i *ACMPCAType) SetPartialName() {
	// "AWS::ACMPCA::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::ACMPCA::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *ACMPCAType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *ACMPCAType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *ACMPCAType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *ACMPCAType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *ACMPCAType) Configure(param interface{}) error {
	config, ok := param.(ACMPCATypeConfig)
	if !ok {
		return errors.New("config is not a valid param (ACMPCATypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *ACMPCAType) SetService(cfg aws.Config) {
	srv := acmpca.New(cfg)

	i.service = srv
}

func (i *ACMPCAType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("acmpca", i.inputName)
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

func (i *ACMPCAType) GetResources() {}

func (i *ACMPCAType) GetResourcesDetail() {}
