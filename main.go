package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/netlify/open-api/go/plumbing"
	"github.com/netlify/open-api/go/plumbing/operations"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Finding deploy preview URL for commit:", request.QueryStringParameters["commit"])
	// Get the deploys
	// var client = NewHTTPClient(nil)
	var client = Default
	var authInfo = nil
	var deploys = Default.ListSiteDeploys(authInfo)
	fmt.Println("Deploys:", deploys)

	const deploy_preview_url = "https://netlify-function--agilepathway-co-uk.netlify.com"
	fmt.Println("Deploy preview url found:", deploy_preview_url)
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       deploy_preview_url,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
