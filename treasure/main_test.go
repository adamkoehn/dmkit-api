package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("it will generate a treasure hoard", func(t *testing.T) {
		response, err := handler(events.APIGatewayProxyRequest{})

		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		if response.StatusCode != 200 {
			t.Errorf("expected 200 but got %v", response.StatusCode)
		}

		if len(response.Body) < 10 {
			t.Errorf("expected some sort of json response but got %v", response.Body)
		}
	})
}
