package database

import (
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDBClient() DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamoDBClient{
		databaseStore: db,
	}
}

const (
	TABLE_NAME = "user_table"
)

func (db DynamoDBClient) DoesUserExist(username string) (bool, error) {
	result, err := db.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})
	if (err != nil) {
		return true, err
		}	
	if result.Item == nil {
		return false, nil
	}
	return true, nil
}

func (db DynamoDBClient) InsertUser(user types.RegisterUser) error {
	item := & dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.UserName),
			},
			"password": {
				S: aws.String((user.Password)),
			},
		},
	}
	_, err := db.databaseStore.PutItem(item)
	
	if err != nil {
		return err
	}
	return nil
}