package svcs

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListBuckets(cfg aws.Config) {
	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// Get the first page of results for ListObjectsV2 for a bucket
	bucketsOutput, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range bucketsOutput.Buckets {
		fmt.Printf("arn:aws:s3:::%s\n", aws.ToString(bucket.Name))
	}

	// Get the first page of results for ListObjectsV2 for a bucket
	directoryBucketsOutput, err := client.ListDirectoryBuckets(context.TODO(), &s3.ListDirectoryBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range directoryBucketsOutput.Buckets {
		fmt.Printf("arn:aws:s3:::%s\n", aws.ToString(bucket.Name))
	}
}
