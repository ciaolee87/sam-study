package main

import (
	"github.com/FlowD2020/aws-sam-go-utils/src/aws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(aws.BizHandler(func(request *aws.AWSRequest) (events.APIGatewayProxyResponse, error) {
		return request.SendJSON(&aws.BizJSON{
			Code: 200,
			Msg:  "Hello World",
			Body: nil,
		})
	}))
}
