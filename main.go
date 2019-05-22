package aws_resources

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
		StackStatus: []cloudformation.StackStatus{
			cloudformation.StackStatusCreateComplete,
		},
		//StackName: "CloudFormationExample",
		//StackName: "awseb-e-82ibs3r2xk-stack",
		StackName: "aws-serverless-repository-alexa-skills-kit-nodejs-factskill-marin",
	}

	iCF := Relations["cloudformation"](cfg)
	if err := iCF.Configure(*cfConfig); err != nil {
		panic(err)
	}
	iCF.GetServices()
	iCF.GetResources()
	iCF.GetResourcesDetail()
}
