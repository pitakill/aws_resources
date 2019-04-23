package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

var cfg aws.Config

func main() {
	config, err := external.LoadDefaultAWSConfig(
		external.WithSharedConfigProfile("dou"),
	)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	cfg = config

	cfConfig := &CloudFormationTypeConfig{
		stackStatus: []cloudformation.StackStatus{
			cloudformation.StackStatusCreateComplete,
		},
		stackName: "CloudFormationExample",
	}

	iCF := relations["cloudformation"](cfg)
	if err := iCF.Configure(*cfConfig); err != nil {
		panic(err)
	}
	iCF.GetServices()
	iCF.GetResources()
	iCF.GetResourcesDetail()
}
