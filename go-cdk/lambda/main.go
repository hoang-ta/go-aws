package main

import (
	"fmt"
	"lambda-func/app"
	"lambda-func/middleware"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	UserName string `json:"username"`
}

func handleRquest(event Event) (string, error) {
	if (event.UserName == "") {
		return "", fmt.Errorf("username cannot be empty")
	}
	return fmt.Sprintf("Called by %s", event.UserName), nil
}

func ProtectedHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body: "Protected route",
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	myapp := app.NewApp()
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)  {
		switch request.Path {
		case "/register":
			return myapp.ApiHandler.ResisterUserHandler(request)
		case "/login":
			return myapp.ApiHandler.LoginUser(request)
		case "/protected":
			return middleware.ValidateJWTMiddleware(ProtectedHandler)(request)
		default:
			return events.APIGatewayProxyResponse{
				Body: "not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})
	// fmt.Println("Hello")
}