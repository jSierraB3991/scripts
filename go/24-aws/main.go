package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Hello world in Deploy With AWS CLI 5")

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}
	return response, nil
}

func main() {
	lambda.Start(handler)
}
