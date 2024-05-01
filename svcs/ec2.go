package svcs

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListEC2Instances(cfg aws.Config) {
	// Create an Amazon EC2 service client
	client := ec2.NewFromConfig(cfg)

	// Get the first page of results for DescribeInstances for a region
	output, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, reservation := range output.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Println(aws.ToString(instance.InstanceId))
		}
	}
}
