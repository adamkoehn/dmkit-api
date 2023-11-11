package main

import (
	"encoding/json"

	"github.com/adamkoehn/dmkit/generator"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	treasure := generator.GetTreasureHoard0To4()
	data, err := json.Marshal(treasure)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "{\"error\":\"could not json encode treasure\"}"}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(data),
	}, nil
}

func main() {
	lambda.Start(handler)
}
