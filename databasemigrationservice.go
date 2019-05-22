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
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
)

type DatabaseMigrationServiceType struct {
	service      *databasemigrationservice.DatabaseMigrationService
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

func DatabaseMigrationServiceFactory(cfg aws.Config) Factory {
	i := new(DatabaseMigrationServiceType)

	i.SetService(cfg)

	return i
}

func (i *DatabaseMigrationServiceType) SetPartialName() {
	// "AWS::DatabaseMigrationService::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::DatabaseMigrationService::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *DatabaseMigrationServiceType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsInput or ListVpcsInput" accordingly
	prefix := search[r{resource: "databasemigrationservice", kind: partialName}]
	name := fmt.Sprintf("%s%sInput", prefix, partialName)

	i.inputName = name
}

func (i *DatabaseMigrationServiceType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsOutput or ListVpcsOutput" accordingly
	prefix := search[r{resource: "databasemigrationservice", kind: partialName}]
	name := fmt.Sprintf("%s%sOutput", prefix, partialName)

	i.outputName = name
}

func (i *DatabaseMigrationServiceType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Add the s to the partialName
	partialName := fmt.Sprintf("%ss", i.partialName)

	// Ex: "Vpc" to "DescribeVpcsRequest or ListVpcsRequest" accordingly
	prefix := search[r{resource: "databasemigrationservice", kind: partialName}]
	name := fmt.Sprintf("%s%sRequest", prefix, partialName)

	i.methodName = name
}

func (i *DatabaseMigrationServiceType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *DatabaseMigrationServiceType) Configure(param interface{}) error {
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

func (i *DatabaseMigrationServiceType) SetService(cfg aws.Config) {
	srv := databasemigrationservice.New(cfg)

	i.service = srv
}

func (i *DatabaseMigrationServiceType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("databasemigrationservice", i.inputName)
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

func (i *DatabaseMigrationServiceType) GetResources() {}

func (i *DatabaseMigrationServiceType) GetResourcesDetail() {}
