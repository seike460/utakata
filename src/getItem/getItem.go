package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	utils "github.com/seike460/utakata/src"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	svc := dynamodb.New(session.New())
	input := &dynamodb.ScanInput{
		TableName: aws.String("tasks"),
	}

	result, err := svc.Scan(input)
	if err != nil {
		utils.AwsErrorPrint(err)
	}

	//headers := []string{
	//	"Access-Control-Allow-Origin": "*", // Required for CORS support to work
	//}

	//jsonHeader, err := json.Marshal(result.Items)
	jsonString, err := json.Marshal(result.Items)

	fmt.Println(jsonString)

	return events.APIGatewayProxyResponse{
			Body: string(jsonString),
			//Headers:    string(jsonHeader),
			StatusCode: 200,
		},
		nil
}

func main() {
	lambda.Start(handleRequest)
}
