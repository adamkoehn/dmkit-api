package main

import (
	"encoding/json"

	"github.com/adamkoehn/dmkit/generator"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	npc := generator.GenerateNpc()
	data, err := json.Marshal(npc)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "{\"error\":\"could not json encode npc\"}"}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(data),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
