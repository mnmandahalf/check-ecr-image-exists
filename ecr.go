package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func main() {
	awsAccessKeyID := flag.String("access-key-id", "", "specify aws access-key-id")
	awsSecretAccessKey := flag.String("secret-access-key", "", "specify aws secret-access-key")
	region := flag.String("region", "", "specify aws region")
	RepositoryName := flag.String("repository-name", "", "specify repository name")
	imageTag := flag.String("tag", "", "specify image tag")
	flag.Parse()
	cfg, err := configure(awsAccessKeyID, awsSecretAccessKey, region)
	if err != nil {
		return
	}
	checkImageExists(cfg, RepositoryName, imageTag)
}

func configure(awsAccessKeyID *string, awsSecretAccessKey *string, region *string) (*aws.Config, error) {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// read AWS Access Keys and overwrite credential
	if *awsAccessKeyID != "" && *awsSecretAccessKey != "" {
		cfg.Credentials = credentials.NewStaticCredentialsProvider(*awsAccessKeyID, *awsSecretAccessKey, os.Getenv("AWS_SESSION_TOKEN"))
	}
	cfg.Region = *region
	return &cfg, err
}

func checkImageExists(cfg *aws.Config, repositoryName *string, imageTag *string) {
	svc := ecr.NewFromConfig(*cfg)
	paginator := ecr.NewListImagesPaginator(
		svc,
		&ecr.ListImagesInput{
			RepositoryName: repositoryName,
		},
	)

	for paginator.HasMorePages() {
		out, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, image := range out.ImageIds {
			if image.ImageTag != nil && *image.ImageTag == *imageTag {
				fmt.Println("1")
				return
			}
		}
	}

	fmt.Println("0")
	return
}
