// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// 2019-05-22 12:14:33.477778006 -0500 CDT m=+0.000119556
package aws_resources

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transcribeservice"
)

type TranscribeServiceType struct {
	service      *transcribeservice.TranscribeService
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func TranscribeServiceFactory(cfg aws.Config) Factory {
	i := new(TranscribeServiceType)

	i.SetService(cfg)

	return i
}

func (i *TranscribeServiceType) SetPartialName() {
	// "AWS::TranscribeService::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::TranscribeService::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *TranscribeServiceType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsInput or ListVpcsInput" accordingly
	prefix := search[r{resource: "transcribeservice", kind: partialName}]
	name := fmt.Sprintf("%s%sInput", prefix, partialName)

	i.inputName = name
}

func (i *TranscribeServiceType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsOutput or ListVpcsOutput" accordingly
	prefix := search[r{resource: "transcribeservice", kind: partialName}]
	name := fmt.Sprintf("%s%sOutput", prefix, partialName)

	i.outputName = name
}

func (i *TranscribeServiceType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsRequest or ListVpcsRequest" accordingly
	prefix := search[r{resource: "transcribeservice", kind: partialName}]
	name := fmt.Sprintf("%s%sRequest", prefix, partialName)

	i.methodName = name
}

func (i *TranscribeServiceType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *TranscribeServiceType) Configure(param interface{}) error {
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

func (i *TranscribeServiceType) SetService(cfg aws.Config) {
	srv := transcribeservice.New(cfg)

	i.service = srv
}

func (i *TranscribeServiceType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("transcribeservice", i.inputName)
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

func (i *TranscribeServiceType) GetResources() {}

func (i *TranscribeServiceType) GetResourcesDetail() {}
