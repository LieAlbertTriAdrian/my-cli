package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"	
 	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetAWSEC2InstanceInformation(RandomString string) string {
	// AWSTry()

	return RandomString
}

func AWSTry () {
	Ec2Service := ec2.New(session.New())

	Input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String("i-1234567890abcdef0"),
		},
	}

	_, err := Ec2Service.DescribeInstances(Input)

	if err != nil {
		fmt.Println(err.Error());
	}

	fmt.Println("resutl");
}
