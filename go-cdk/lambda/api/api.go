package api

import (
	"encoding/json"
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) ResisterUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var registerUser types.RegisterUser

	err := json.Unmarshal([]byte(request.Body), &registerUser)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: "Invalid request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	if registerUser.UserName == "" || registerUser.Password == "" {
		return events.APIGatewayProxyResponse{
			Body: "Invalid request - empty params",
			StatusCode: http.StatusBadRequest,
		}, nil
	}
	userExists, err := api.dbStore.DoesUserExist(registerUser.UserName);
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("Internal server error - There is some error from Dynamo %w", err)
	}
	if userExists {
		return events.APIGatewayProxyResponse{
			Body: fmt.Sprintf("That username already exists"),
			StatusCode: http.StatusConflict,
		}, nil
	}
	err = api.dbStore.InsertUser(registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("Internal server error - Error register user %w", err)
	}
	return events.APIGatewayProxyResponse{
		Body: "Success register user",
		StatusCode: http.StatusCreated,
	}, nil
}

func (api ApiHandler) LoginUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	type LoginRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest

	err := json.Unmarshal([]byte(request.Body), &loginRequest)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: "Invalid request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	user, err := api.dbStore.GetUser(loginRequest.UserName)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body: "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("Internal server error - Get user %w", err)
	}

	if !types.ValidatePassword(user.PasswordHash, loginRequest.Password) {
		return events.APIGatewayProxyResponse{
			Body: "Invalid credential",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	token := types.CreateToken(user)


	return events.APIGatewayProxyResponse{
		Body: fmt.Sprintf("Success login with token %s", token),
		StatusCode: http.StatusOK,
	}, nil
}