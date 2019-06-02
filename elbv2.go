// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
<<<<<<< HEAD
// 2019-06-04 17:54:17.26188712 -0500 CDT m=+0.000169213
=======
// 2019-06-02 13:56:03.549666 -0500 CDT m=+0.005104325
>>>>>>> just added info to grab the data for experiment
package aws_resources

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elbv2"
)

type ELBV2Type struct {
	service      *elbv2.ELBV2
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func ELBV2Factory(cfg aws.Config) Factory {
	i := new(ELBV2Type)

	i.SetService(cfg)

	return i
}

func (i *ELBV2Type) SetPartialName() {
	// "AWS::ELBV2::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::ELBV2::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *ELBV2Type) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsInput or ListVpcsInput" accordingly
	prefix := search[r{resource: "elbv2", kind: partialName}]
	name := fmt.Sprintf("%s%sInput", prefix, partialName)

	i.inputName = name
}

func (i *ELBV2Type) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsOutput or ListVpcsOutput" accordingly
	prefix := search[r{resource: "elbv2", kind: partialName}]
	name := fmt.Sprintf("%s%sOutput", prefix, partialName)

	i.outputName = name
}

func (i *ELBV2Type) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsRequest or ListVpcsRequest" accordingly
	prefix := search[r{resource: "elbv2", kind: partialName}]
	name := fmt.Sprintf("%s%sRequest", prefix, partialName)

	i.methodName = name
}

func (i *ELBV2Type) SetResourceType(t string) {
	i.resourceType = t
}

func (i *ELBV2Type) Configure(param interface{}) error {
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

func (i *ELBV2Type) SetService(cfg aws.Config) {
	srv := elbv2.New(cfg)

	i.service = srv
}

func (i *ELBV2Type) GetServices() (reflect.Value, error) {
	if i.methodName == "" {
		return reflect.Value{}, errors.New("method can't be an empty string")
	}

	instance, err := typeRegistry.Get("elbv2", i.inputName)
	if err != nil {
		// We can ignore this kind of errors because there is not resources by the
		// i.inputName
		log.Println(err)
		// This seems odd, we can do better
		// We need to think about this again
		return reflect.Value{}, nil
	}

	method := reflect.ValueOf(i.service).MethodByName(i.methodName)
	called := method.Call([]reflect.Value{reflect.ValueOf(instance)})

	send := reflect.Indirect(called[0]).MethodByName("Send")
	calledSend := send.Call([]reflect.Value{})

	return calledSend[0], nil
}

func (i *ELBV2Type) GetResources() error { return nil }

func (i *ELBV2Type) GetResourcesDetail() ([]reflect.Value, error) {
	return []reflect.Value{}, nil
}
