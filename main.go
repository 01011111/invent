package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"svcs"

	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	profile = flag.String("profile", "default", "AWS profile to use")
	region  = flag.String("region", "", "AWS region to use")
)

func main() {
	flag.Parse()

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(*profile))
	if err != nil {
		log.Fatal(err)
	}

	if *region != "" {
		cfg.Region = *region
	}

	// Profile loaded from shared config file.
	fmt.Println("Profile:", *profile)
	// Region loaded from credentials file.
	fmt.Println("Region:", cfg.Region)
	fmt.Print("\n")

	// List S3 Buckets
	fmt.Println("S3 Buckets")
	svcs.ListBuckets(cfg)
	fmt.Print("\n")

	// List Lambda Functions
	fmt.Println("Lambda Functions")
	svcs.ListLambdas(cfg)
	fmt.Print("\n")

	// List Lambda Layers
	fmt.Println("Lambda Layers")
	svcs.ListLayers(cfg)
	fmt.Print("\n")

	// List ECS Services
	fmt.Println("ECS Services")
	svcs.ListECSSercices(cfg)
	fmt.Print("\n")

	// List EC2 Instances
	fmt.Println("EC2 Instances")
	svcs.ListEC2Instances(cfg)
	fmt.Print("\n")
}
