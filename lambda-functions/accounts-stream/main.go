package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
	"os"
	"time"
)

const (
	accountCreated = "AccountCreated"
	accountUpdated = "AccountUpdated"
	accountDeleted = "AccountDeleted"
)

var (
	snsTopic = os.Getenv("SNS_TOPIC")
)

type (
	domainEvent struct {
		EventId       string    `json:"event_id"`
		EventType     string    `json:"event_type"`
		AggregateType string    `json:"aggregate_type"`
		AggregateId   string    `json:"aggregate_id"`
		PayloadData   string    `json:"payload_data"`
		Timestamp     time.Time `json:"timestamp"`
	}

	snsMessage struct {
		Default string `json:"default"`
	}
)

func handleRequest(ctx context.Context, event events.DynamoDBEvent) error {
	var accDomainEvent domainEvent
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Println("Unable to load SDK config", err)
		return err
	}

	snsClient := sns.NewFromConfig(cfg)

	for _, record := range event.Records {
		log.Printf("Processing request data for event ID %s\n", record.EventID)

		accDomainEvent, err = mapDynamoEventToDomainEvent(ctx, record)
		if err != nil {
			return err
		}

		if err := publishToSNS(ctx, accDomainEvent, snsClient); err != nil {
			return err
		}
	}

	return nil
}

func publishToSNS(ctx context.Context, accountEvent domainEvent, snsClient *sns.Client) error {
	eventBytes, err := json.Marshal(accountEvent)
	if err != nil {
		return err
	}

	message := snsMessage{Default: string(eventBytes)}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	input := sns.PublishInput{
		Message:          aws.String(string(messageBytes)),
		MessageStructure: aws.String("json"),
		TopicArn:         aws.String(snsTopic),
	}
	output, err := snsClient.Publish(ctx, &input)
	if err != nil {
		return err
	}

	log.Printf("Message published: %s\n", *output.MessageId)
	return nil
}

func mapDynamoEventToDomainEvent(_ context.Context, record events.DynamoDBEventRecord) (domainEvent, error) {
	evenTypeMap := map[string]string{
		string(events.DynamoDBOperationTypeInsert): accountCreated,
		string(events.DynamoDBOperationTypeModify): accountUpdated,
		string(events.DynamoDBOperationTypeRemove): accountDeleted,
	}
	eventType := evenTypeMap[record.EventName]

	log.Printf("Processing %s operation as %s\n", record.EventName, eventType)

	var accId events.DynamoDBAttributeValue
	var docNumber events.DynamoDBAttributeValue

	if eventType == accountDeleted {
		accId, _ = record.Change.OldImage["Id"]
		docNumber, _ = record.Change.OldImage["DocumentNumber"]
	} else {
		accId, _ = record.Change.NewImage["Id"]
		docNumber, _ = record.Change.NewImage["DocumentNumber"]
	}

	payload := struct {
		AccountId      string `json:"account_id"`
		DocumentNumber string `json:"document_number"`
	}{
		AccountId:      accId.String(),
		DocumentNumber: docNumber.String(),
	}

	bytes, err := json.Marshal(payload)
	if err != nil {
		return domainEvent{}, err
	}

	event := domainEvent{
		EventId:       record.EventID,
		EventType:     eventType,
		AggregateType: "Account",
		AggregateId:   payload.AccountId,
		PayloadData:   string(bytes),
		Timestamp:     record.Change.ApproximateCreationDateTime.Time,
	}

	return event, nil
}

func main() {
	lambda.Start(handleRequest)
}
