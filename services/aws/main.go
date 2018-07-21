package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"	
 	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetAWSEC2InstanceInformation(RandomString string) string {
	AWSTry()

	return RandomString
}

func AWSTry () {
// aws ec2 describe-instances --filters 'Name=tag:Name,Values=dev-server-*'
	Ec2Service := ec2.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))

	Input := &ec2.DescribeInstancesInput{
	    Filters: []*ec2.Filter{
	        {
	            Name: aws.String("instance-type"),
	            Values: []*string{
	                aws.String("t2.micro"),
	            },
	        },
	    },
	}

	result, err := Ec2Service.DescribeInstances(Input)

	if err != nil {
		fmt.Println(err.Error());
	}

	fmt.Println(result)
	fmt.Println("resutl");
}