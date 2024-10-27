package main

import (
	"fmt"
)

type Event struct {
	UserName string `json:"username"`
}

func handleRquest(event Event) (string, error) {
	if (event.UserName == "") {
		return "", fmt.Errorf("username cannot be empty")
	}
	return fmt.Sprint("Called by %s", event.UserName), nil
}

func main() {
	// lambda.Start(handleRquest)
	fmt.Println("Hello")
}