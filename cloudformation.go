package aws_resources

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

type CloudFormationType struct {
	service       *cloudformation.CloudFormation
	status        []cloudformation.StackStatus
	stackName     string
	stackResource cloudformation.StackSummary
	resources     []cloudformation.StackResource
}

type CloudFormationTypeConfig struct {
	StackStatus []cloudformation.StackStatus
	StackName   string
}

func CloudFormationFactory(cfg aws.Config) Factory {
	i := new(CloudFormationType)

	i.SetService(cfg)

	return i
}

func (i *CloudFormationType) Configure(param interface{}) error {
	config, ok := param.(CloudFormationTypeConfig)
	if !ok {
		return errors.New("config is not a valid param (CloudFormationTypeConfig)")
	}

	i.SetStackName(config.StackName)
	i.SetStatus(config.StackStatus)

	return nil
}

func (i *CloudFormationType) SetService(cfg aws.Config) {
	srv := cloudformation.New(cfg)

	i.service = srv
}

func (i *CloudFormationType) SetStatus(status []cloudformation.StackStatus) {
	i.status = status
}

func (i *CloudFormationType) SetStackName(name string) {
	i.stackName = name
}

func (i *CloudFormationType) GetServices() {
	req := i.service.ListStacksRequest(&cloudformation.ListStacksInput{
		StackStatusFilter: i.status,
	})

	res, err := req.Send()
	if err != nil {
		panic(err.Error())
	}

	for _, resource := range res.StackSummaries {
		if *resource.StackName == i.stackName {
			i.stackResource = resource
		}
	}
}

func (i *CloudFormationType) GetResources() {
	req := i.service.DescribeStackResourcesRequest(&cloudformation.DescribeStackResourcesInput{
		StackName: i.stackResource.StackName,
	})

	res, err := req.Send()
	if err != nil {
		panic("failed to get the resources, " + err.Error())
	}

	i.resources = res.StackResources
}

func (i *CloudFormationType) GetResourcesDetail() {
	for _, resource := range i.resources {
		config := &TypeConfig{
			resourceType: *resource.ResourceType,
		}

		s := getKind(*resource.ResourceType)

		fmt.Println(s)

		instance := Relations[s](cfg)
		if err := instance.Configure(*config); err != nil {
			panic(err)
		}
		instance.GetServices()
	}
}
