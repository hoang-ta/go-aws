package main

import (
	"fmt"
	"lambda-func/app"

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

func main() {
	lambda.Start(app.NewApp().ApiHandler.ResisterUserHandler)
	// fmt.Println("Hello")
}