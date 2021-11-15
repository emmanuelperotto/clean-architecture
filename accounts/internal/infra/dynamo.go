package infra

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

//ConnectDynamoDB connects to dynamodb and returns a client
func ConnectDynamoDB() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("pocs"),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalln("DynamoDB Configuration error", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return client
}
