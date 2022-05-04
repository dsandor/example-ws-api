package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	// Optional output with custom properties of the String, Number or Boolean type.
	authResponse.Context = map[string]interface{}{
		"stringKey":  "stringval",
		"numberKey":  123,
		"booleanKey": true,
	}
	return authResponse
}

func (d *dependencies) auth(ctx context.Context, websocketEvent events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	fmt.Println("Auth handler hit.")
	fmt.Printf("%+v\n", websocketEvent)

	policy := generatePolicy("user", "Allow", websocketEvent.MethodArn)

	fmt.Printf("policy: %+v\n", policy)

	return policy, nil
}
