package dynamo

import (
	"accounts/internal/domain/entity"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

const tableName = "Accounts"

type (
	accountReadOnlyRepository struct {
		client *dynamodb.Client
	}

	accountWriteOnlyRepository struct {
		client *dynamodb.Client
	}
)

func newAccountReadOnlyRepo(client *dynamodb.Client) accountReadOnlyRepository {
	return accountReadOnlyRepository{
		client: client,
	}
}

func newAccountWriteOnlyRepo(client *dynamodb.Client) accountWriteOnlyRepository {
	return accountWriteOnlyRepository{
		client: client,
	}
}

//FindById implementation for DynamoDB query by ID
func (a accountReadOnlyRepository) FindById(ctx context.Context, id string) (entity.Account, error) {
	attrValueMap, err := attributevalue.MarshalMap(struct{ Id string }{Id: id})
	if err != nil {
		return entity.Account{}, errors.New("couldn't marshal input")
	}

	input := dynamodb.GetItemInput{
		Key:       attrValueMap,
		TableName: aws.String(tableName),
	}

	output, err := a.client.GetItem(ctx, &input)
	if err != nil {
		return entity.Account{}, err
	}

	if output.Item == nil {
		return entity.Account{}, errors.New("account not found")
	}

	account := entity.Account{}

	err = attributevalue.UnmarshalMap(output.Item, &account)
	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

//Create implementation for DynamoDB put request
func (a accountWriteOnlyRepository) Create(ctx context.Context, account entity.Account) (entity.Account, error) {
	attrValueMap, err := attributevalue.MarshalMap(struct {
		Id             string
		DocumentNumber string
	}{
		Id:             uuid.NewString(),
		DocumentNumber: account.DocumentNumber,
	})

	input := dynamodb.PutItemInput{
		Item:      attrValueMap,
		TableName: aws.String(tableName),
	}

	_, err = a.client.PutItem(ctx, &input)
	if err != nil {
		return entity.Account{}, err
	}

	acc := entity.Account{}

	err = attributevalue.UnmarshalMap(input.Item, &acc)
	if err != nil {
		return entity.Account{}, err
	}

	return acc, nil
}
