package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2Type struct {
	service      *ec2.EC2
	resourceType string
	methodName   string
	inputName    string
	partialName  string
}

type EC2TypeConfig struct {
	resourceType string
}

func EC2Factory(cfg aws.Config) Factory {
	i := new(EC2Type)

	i.SetService(cfg)

	return i
}

var notValid = []string{
	"AWS::EC2::VPCGatewayAttachment",
	"AWS::EC2::Route",
	"AWS::EC2::SubnetRouteTableAssociation",
}

func (i *EC2Type) SetPartialName() {
	if ContainsString(notValid, i.resourceType) {
		return
	}

	// "AWS::EC2::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::EC2::", "")

	if name == "VPC" {
		// "VPC" to "vpc"
		name = strings.ToLower(name)

		// "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *EC2Type) SetInputName() {
	if i.partialName == "" {
		return
	}

	// "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *EC2Type) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *EC2Type) SetResourceType(t string) {
	i.resourceType = t
}

func (i *EC2Type) Configure(param interface{}) error {
	config, ok := param.(EC2TypeConfig)
	if !ok {
		return errors.New("config is not a valid param (EC2TypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()

	return nil
}

func (i *EC2Type) SetService(cfg aws.Config) {
	srv := ec2.New(cfg)

	i.service = srv
}

func (i *EC2Type) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get(i.inputName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", i)

	method := reflect.ValueOf(i.service).MethodByName(i.methodName)
	called := method.Call([]reflect.Value{reflect.ValueOf(instance)})

	fmt.Println(&called)

	//req := i.service.DescribeVpcsRequest(&ec2.DescribeVpcsInput{})

	//res, err := req.Send()
	//if err != nil {
	//panic(err.Error())
	//}

	//fmt.Println(res)
}

func (i *EC2Type) GetResources() {}

func (i *EC2Type) GetResourcesDetail() {}
