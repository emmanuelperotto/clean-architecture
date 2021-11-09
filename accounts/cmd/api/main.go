package main

import (
	"accounts/internal/adapter/database/mysql"
	"accounts/internal/adapter/web"
	"accounts/internal/domain/usecase"
	"accounts/internal/infra"
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/gofiber/fiber/v2"
	"log"
)

// SNSListTopicsAPI defines the interface for the ListTopics function.
// We use this interface to test the function using a mocked service.
type SNSListTopicsAPI interface {
	ListTopics(ctx context.Context,
		params *sns.ListTopicsInput,
		optFns ...func(*sns.Options)) (*sns.ListTopicsOutput, error)
}

// GetTopics retrieves information about the Amazon Simple Notification Service (Amazon SNS) topics
// Inputs:
//     c is the context of the method call, which includes the Region
//     api is the interface that defines the method call
//     input defines the input arguments to the service call.
// Output:
//     If success, a ListTopicsOutput object containing the result of the service call and nil
//     Otherwise, nil and an error from the call to ListTopics
func GetTopics(c context.Context, api SNSListTopicsAPI, input *sns.ListTopicsInput) (*sns.ListTopicsOutput, error) {
	return api.ListTopics(c, input)
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("pocs"),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:4566", SigningRegion: region}, nil
			})),
	)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sns.NewFromConfig(cfg)

	input := &sns.ListTopicsInput{}

	results, err := GetTopics(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving information about the SNS topics:")
		fmt.Println(err)
		return
	}

	for _, t := range results.Topics {
		fmt.Println(*t.TopicArn)
	}

	db, err := infra.ConnectMySQLDB()
	if err != nil {
		log.Fatalln(err)
	}

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Println("Can't close DB Connection")
		}
	}(db)

	app := fiber.New()
	repositoryRegistry := mysql.NewMySQLRepositoryRegistry(db) // Replace with "local.NewLocalRepositoryRegistry()" if you want to test local storage

	accountHandler := web.NewAccountHandler(
		usecase.NewCreateAccountUseCase(repositoryRegistry),
		usecase.NewGetAccountUseCase(repositoryRegistry),
	)

	app.Post("/accounts", accountHandler.CreateAccount)
	app.Get("/accounts/:id", accountHandler.GetAccount)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}
