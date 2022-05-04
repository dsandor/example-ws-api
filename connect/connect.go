package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func (d *dependencies) connect(ctx context.Context, websocketEvent events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Connect handler hit.")
	fmt.Printf("%+v\n", websocketEvent)

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         nil,
		IsBase64Encoded: false,
	}, nil
}
