package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

type DefaultResponse struct {
	StatusCode      int    `json:"statusCode"`
	Body            string `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded"`
	Headers         struct {
		ContentType string `json:"content-type"`
	} `json:"headers"`
}

func (d *dependencies) list(ctx context.Context, websocketEvent events.APIGatewayWebsocketProxyRequest) (DefaultResponse, error) {
	fmt.Println("Default route handler hit.")
	fmt.Printf("%+v\n", websocketEvent)

	response := &DefaultResponse{
		StatusCode:      200,
		Body:            "Hello, from default route.",
		IsBase64Encoded: false,
		Headers: struct {
			ContentType string `json:"content-type"`
		}(struct{ ContentType string }{
			ContentType: "application/json",
		})}

	return *response, nil
}
