package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

// Type Factory
type Factory interface {
	Configure(interface{}) error
	GetServices()
	GetResources()
	GetResourcesDetail()
	SetService(aws.Config)
}

type Info func(aws.Config) Factory

var relations = map[string]Info{
	"cloudformation": CloudFormationFactory,
	"ec2":            EC2Factory,
}
