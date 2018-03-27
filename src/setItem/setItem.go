package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/lestrrat-go/strftime"
	utils "github.com/seike460/utakata/src"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	f, _ := strftime.New("%Y%m%d%H%M%S")
	id := f.FormatString(time.Now())

	fmt.Println(request.Body)

	svc := dynamodb.New(session.New())
	// JsonをMapに変換
	byt := []byte(request.Body)
	var bodyMap map[string]interface{}
	if err := json.Unmarshal(byt, &bodyMap); err != nil {
		panic(err)
	}

	utils.SlackSend(bodyMap["body"].(string))

	fmt.Println("id ")
	fmt.Println(id)
	fmt.Println("bodyMap ")
	fmt.Println(bodyMap)

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
			"name": {
				S: aws.String(bodyMap["body"].(string)),
			},
		},
		TableName: aws.String("tasks"),
	}

	_, err := svc.PutItem(input)
	if err != nil {
		utils.AwsErrorPrint(err)
	}

	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
