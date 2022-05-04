#!/bin/sh
sam build
sam deploy --no-confirm-changeset --no-fail-on-empty-changeset --stack-name "example-ws-api" --capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM --s3-bucket example-ws-api-deploy --parameter-overrides ParameterKey=EnvironmentTagName,ParameterValue=dev
