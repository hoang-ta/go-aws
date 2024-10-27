package database

import (
	"fmt"
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

type UserStore interface {
	DoesUserExist(username string) (bool, error)
	InsertUser(user types.RegisterUser) error
	GetUser(username string) (types.User, error)
}

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

func (db DynamoDBClient) InsertUser(registerUser types.RegisterUser) error {
	user, err := types.NewUser(registerUser)
	if err != nil {
		return err
	}
	item := & dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.UserName),
			},
			"password": {
				S: aws.String((user.PasswordHash)),
			},
		},
	}
	_, err = db.databaseStore.PutItem(item)
	
	if err != nil {
		return err
	}
	return nil
}

func (db DynamoDBClient) GetUser(username string) (types.User, error) {
	var user types.User

	result, err := db.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue {
			"username": {
				S: &username,
			},
		},
	})
	if err != nil {
		return user, err
	}
	if result.Item == nil {
		return user, fmt.Errorf("user not found")
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}