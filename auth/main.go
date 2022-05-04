package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type dependencies struct {
}

func main() {
	d := dependencies{}

	fmt.Println("Starting auth handler.")
	lambda.Start(d.auth)
}
