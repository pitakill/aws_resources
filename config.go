package aws_resources

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

var cfg aws.Config

func SetConfig (localConf aws.Config) {
    cfg = localConf
}

func Config () aws.Config {
     return cfg
}
