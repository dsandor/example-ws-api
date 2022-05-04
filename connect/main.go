package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type dependencies struct {
}

func main() {
	d := dependencies{}

	fmt.Println("Starting connect handler.")
	lambda.Start(d.connect)
}
