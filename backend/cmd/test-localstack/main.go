package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

func main() {
	// LocalStackに接続するためのAWS設定
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:           "http://localhost:4566",
						SigningRegion: "us-east-1",
					}, nil
				},
			),
		),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(
			aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     "test",
					SecretAccessKey: "test",
				}, nil
			}),
		),
	)
	if err != nil {
		log.Fatalf("Unable to load AWS SDK config: %v", err)
	}

	// S3サービスのテスト
	s3Client := s3.NewFromConfig(cfg)
	buckets, err := s3Client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("Unable to list S3 buckets: %v", err)
	}
	fmt.Println("S3 Buckets:")
	for _, bucket := range buckets.Buckets {
		fmt.Printf("- %s\n", *bucket.Name)
	}

	// CloudWatch Logsのテスト
	cwlClient := cloudwatchlogs.NewFromConfig(cfg)
	logGroups, err := cwlClient.DescribeLogGroups(context.Background(), &cloudwatchlogs.DescribeLogGroupsInput{})
	if err != nil {
		log.Fatalf("Unable to list CloudWatch Log groups: %v", err)
	}
	fmt.Println("\nCloudWatch Log Groups:")
	for _, group := range logGroups.LogGroups {
		fmt.Printf("- %s\n", *group.LogGroupName)
	}

	// X-Rayのテスト
	xrayClient := xray.NewFromConfig(cfg)
	samplingRules, err := xrayClient.GetSamplingRules(context.Background(), &xray.GetSamplingRulesInput{})
	if err != nil {
		log.Fatalf("Unable to get X-Ray sampling rules: %v", err)
	}
	fmt.Println("\nX-Ray Sampling Rules:")
	for _, rule := range samplingRules.SamplingRuleRecords {
		fmt.Printf("- %s\n", *rule.SamplingRule.RuleName)
	}

	fmt.Println("\nAll tests completed successfully!")
}
