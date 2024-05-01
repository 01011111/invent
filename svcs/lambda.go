package svcs

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func ListLambdas(cfg aws.Config) {
	// Create an Amazon Lambda service client
	client := lambda.NewFromConfig(cfg)

	// Get the first page of results for ListFunctions for a region
	output, err := client.ListFunctions(context.TODO(), &lambda.ListFunctionsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, fn := range output.Functions {
		fmt.Println(aws.ToString(fn.FunctionArn))
	}
}

func ListLayers(cfg aws.Config) {
	// Create an Amazon Lambda service client
	client := lambda.NewFromConfig(cfg)

	// Get the first page of results for ListLayers for a region
	output, err := client.ListLayers(context.TODO(), &lambda.ListLayersInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, layer := range output.Layers {
		fmt.Println(aws.ToString(layer.LayerArn))
	}
}
