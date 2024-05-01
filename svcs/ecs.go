package svcs

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func ListECSClusters(client *ecs.Client) []string {
	// Get the first page of results for ListClusters for a region
	output, err := client.ListClusters(context.TODO(), &ecs.ListClustersInput{})
	if err != nil {
		log.Fatal(err)
	}

	return output.ClusterArns
}

func ListECSSercices(cfg aws.Config) {
	// Create an Amazon ECS service client
	client := ecs.NewFromConfig(cfg)

	clusters := ListECSClusters(client)

	for _, cluster := range clusters {
		// Get the first page of results for ListServices for a region
		output, err := client.ListServices(context.TODO(), &ecs.ListServicesInput{
			Cluster: &cluster,
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, arn := range output.ServiceArns {
			fmt.Println(arn)
		}
	}
}
