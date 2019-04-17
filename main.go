package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Container holds all the CloudFormation information
type Container struct {
	account   string
	config    aws.Config
	name      string
	resources []cloudformation.StackResource
	cfservice *cloudformation.CloudFormation
	stacks    []cloudformation.StackSummary
	stack     cloudformation.StackSummary
}

func main() {
	container := &Container{
		// Get these fields dynamically i.e. with a flag
		account: "dou",
		name:    "CloudFormationExample",
	}

	container.SetConfig().SetService().SetStacks(cloudformation.StackStatusCreateComplete).SetStack().GetResources()

	for _, resource := range container.resources {
		if *resource.ResourceType == "AWS::EC2::VPC" {
			svc := ec2.New(container.config)

			req := svc.DescribeVpcsRequest(&ec2.DescribeVpcsInput{})

			res, err := req.Send()
			if err != nil {
				panic(err.Error())
			}

			fmt.Println(res)
		}
	}
}

func (c *Container) SetConfig() *Container {
	cfg, err := external.LoadDefaultAWSConfig(
		external.WithSharedConfigProfile(c.account),
	)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	c.config = cfg

	return c
}

func (c *Container) SetService() *Container {
	c.cfservice = cloudformation.New(c.config)

	return c
}

func (c *Container) SetStacks(status cloudformation.StackStatus) *Container {
	req := c.cfservice.ListStacksRequest(&cloudformation.ListStacksInput{
		StackStatusFilter: []cloudformation.StackStatus{status},
	})

	res, err := req.Send()
	if err != nil {
		panic("failed to list instances, " + err.Error())
	}

	c.stacks = res.StackSummaries

	return c
}

func (c *Container) SetStack() *Container {
	for _, stack := range c.stacks {
		if c.name == *stack.StackName {
			c.stack = stack

			break
		}
	}

	return c
}

func (c *Container) GetResources() *Container {
	req := c.cfservice.DescribeStackResourcesRequest(&cloudformation.DescribeStackResourcesInput{
		StackName: c.stack.StackName,
	})

	res, err := req.Send()
	if err != nil {
		panic("failed to get the resources, " + err.Error())
	}

	c.resources = res.StackResources

	return c
}
