package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

// this will take in some payload and do some process w it
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("successfully called %s", event.Username), nil
}

func main() {
	lambda.Start(HandleRequest)
}
