package dynamo

import (
	"accounts/internal/domain/repository"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamoDBRepositoryRegistry struct {
	client *dynamodb.Client
}

//NewDynamoDBRepositoryRegistry builds dynamoDBRepositoryRegistry with its dependencies
func NewDynamoDBRepositoryRegistry(client *dynamodb.Client) repository.Registry {
	return dynamoDBRepositoryRegistry{
		client: client,
	}
}

//AccountReadOnlyRepository returns read only dynamo repo for Accounts table
func (d dynamoDBRepositoryRegistry) AccountReadOnlyRepository() repository.AccountReadOnly {
	return newAccountReadOnlyRepo(d.client)
}

//AccountWriteOnlyRepository returns write only dynamo repo for Accounts table
func (d dynamoDBRepositoryRegistry) AccountWriteOnlyRepository() repository.AccountWriteOnly {
	return newAccountWriteOnlyRepo(d.client)
}

//DoInTransaction still don't make things transactional, but it's necessary to implement repository.Registry interface
func (d dynamoDBRepositoryRegistry) DoInTransaction(txFunc repository.InTransaction) (out interface{}, err error) {
	return txFunc(d)
}
