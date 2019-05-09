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
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

type ApiGatewayV2Type struct {
	service      *apigatewayv2.ApiGatewayV2
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type ApiGatewayV2TypeConfig struct {
	resourceType string
}

func ApiGatewayV2Factory(cfg aws.Config) Factory {
	i := new(ApiGatewayV2Type)

	i.SetService(cfg)

	return i
}

func (i *ApiGatewayV2Type) SetPartialName() {
	// "AWS::ApiGatewayV2::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::ApiGatewayV2::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *ApiGatewayV2Type) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *ApiGatewayV2Type) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *ApiGatewayV2Type) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *ApiGatewayV2Type) SetResourceType(t string) {
	i.resourceType = t
}

func (i *ApiGatewayV2Type) Configure(param interface{}) error {
	config, ok := param.(ApiGatewayV2TypeConfig)
	if !ok {
		return errors.New("config is not a valid param (ApiGatewayV2TypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *ApiGatewayV2Type) SetService(cfg aws.Config) {
	srv := apigatewayv2.New(cfg)

	i.service = srv
}

func (i *ApiGatewayV2Type) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("apigatewayv2", i.inputName)
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

func (i *ApiGatewayV2Type) GetResources() {}

func (i *ApiGatewayV2Type) GetResourcesDetail() {}
