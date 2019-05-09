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
	"github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
)

type MarketplaceEntitlementServiceType struct {
	service      *marketplaceentitlementservice.MarketplaceEntitlementService
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type MarketplaceEntitlementServiceTypeConfig struct {
	resourceType string
}

func MarketplaceEntitlementServiceFactory(cfg aws.Config) Factory {
	i := new(MarketplaceEntitlementServiceType)

	i.SetService(cfg)

	return i
}

func (i *MarketplaceEntitlementServiceType) SetPartialName() {
	// "AWS::MarketplaceEntitlementService::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::MarketplaceEntitlementService::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *MarketplaceEntitlementServiceType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *MarketplaceEntitlementServiceType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *MarketplaceEntitlementServiceType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *MarketplaceEntitlementServiceType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *MarketplaceEntitlementServiceType) Configure(param interface{}) error {
	config, ok := param.(MarketplaceEntitlementServiceTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (MarketplaceEntitlementServiceTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *MarketplaceEntitlementServiceType) SetService(cfg aws.Config) {
	srv := marketplaceentitlementservice.New(cfg)

	i.service = srv
}

func (i *MarketplaceEntitlementServiceType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("marketplaceentitlementservice", i.inputName)
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

func (i *MarketplaceEntitlementServiceType) GetResources() {}

func (i *MarketplaceEntitlementServiceType) GetResourcesDetail() {}