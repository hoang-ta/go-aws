package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) ResisterUserHandler(event types.RegisterUser) error {
	if event.UserName == "" || event.Password == "" {
		return fmt.Errorf("Empty params")
	}
	userExists, error := api.dbStore.DoesUserExist(event.UserName);
	if error != nil {
		return fmt.Errorf("There is some error from Dynamo %w", error)
	}
	if userExists {
		return fmt.Errorf("That username already exists")
	}
	error = api.dbStore.InsertUser(event)
	if error != nil {
		return fmt.Errorf("Error register user %w", error)
	}
	return nil
}