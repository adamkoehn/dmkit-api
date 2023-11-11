package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	t.Run("it returns an npc that is randomly generated", func(t *testing.T) {
		response, err := handler(events.APIGatewayProxyRequest{})

		if err != nil {
			t.Errorf("An error was produced instead of a successful response %v", err)
		}

		if len(response.Body) < 10 {
			t.Errorf("Expected a response to have something in it but got %v", response.Body)
		}

		if response.StatusCode != 200 {
			t.Errorf("Expected status code 200, but got %v", response.StatusCode)
		}
	})
}
