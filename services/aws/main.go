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
	Ec2Client := ec2.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))

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

	result, err := Ec2Client.DescribeInstances(Input)

	if err != nil {
		fmt.Println(err.Error());
	}

	fmt.Println(result)
	fmt.Println("resutl");
}