package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

func handleRequest(ctx context.Context, event events.DynamoDBEvent) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Println("Unable to load SDK config", err)
		return err
	}

	_ = sns.NewFromConfig(cfg)

	for _, record := range event.Records {
		log.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

		// Print new values for attributes of type String
		for name, value := range record.Change.NewImage {
			log.Println("Attribute: ", name)
			log.Println("Value: ", value.String())

			if value.DataType() == events.DataTypeString {
				log.Printf("Attribute name: %s, value: %s\n", name, value.String())
			}
		}
	}

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
