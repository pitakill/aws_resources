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
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

type ElastiCacheType struct {
	service      *elasticache.ElastiCache
	resourceType string
	methodName   string
	inputName    string
	outputName   string
	partialName  string
}

type ElastiCacheTypeConfig struct {
	resourceType string
}

func ElastiCacheFactory(cfg aws.Config) Factory {
	i := new(ElastiCacheType)

	i.SetService(cfg)

	return i
}

func (i *ElastiCacheType) SetPartialName() {
	// "AWS::ElastiCache::VPC" to "VPC"
	name := strings.ReplaceAll(i.resourceType, "AWS::ElastiCache::", "")

	if name == "VPC" {
		// Ex: "VPC" to "vpc"
		name = strings.ToLower(name)

		// Ex: "vpc" to "Vpc"
		name = strings.Title(name)
	}

	i.partialName = name
}

func (i *ElastiCacheType) SetInputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsInput"
	name := fmt.Sprintf("Describe%ssInput", i.partialName)

	i.inputName = name
}

func (i *ElastiCacheType) SetOutputName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsOutput"
	name := fmt.Sprintf("Describe%ssOutput", i.partialName)

	i.outputName = name
}

func (i *ElastiCacheType) SetMethodName() {
	if i.partialName == "" {
		return
	}

	// Ex: "Vpc" to "DescribeVpcsRequest"
	name := fmt.Sprintf("Describe%ssRequest", i.partialName)

	i.methodName = name
}

func (i *ElastiCacheType) SetResourceType(t string) {
	i.resourceType = t
}

func (i *ElastiCacheType) Configure(param interface{}) error {
	config, ok := param.(ElastiCacheTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (ElastiCacheTypeConfig)")
	}

	i.SetResourceType(config.resourceType)
	i.SetPartialName()
	i.SetMethodName()
	i.SetInputName()
	i.SetOutputName()

	return nil
}

func (i *ElastiCacheType) SetService(cfg aws.Config) {
	srv := elasticache.New(cfg)

	i.service = srv
}

func (i *ElastiCacheType) GetServices() {
	if i.methodName == "" {
		return
	}

	instance, err := typeRegistry.Get("elasticache", i.inputName)
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

func (i *ElastiCacheType) GetResources() {}

func (i *ElastiCacheType) GetResourcesDetail() {}
