package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"

	// "github.com/netlify/open-api/go/plumbing"
	// "github.com/netlify/open-api/go/plumbing/operations"
	

	// "github.com/go-openapi/runtime"
	// httptransport "github.com/go-openapi/runtime/client"

	// strfmt "github.com/go-openapi/strfmt"

)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Finding deploy preview URL for commit:", request.QueryStringParameters["commit"])
	lambdacontext.FromContext(ctx)
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       "Something went wrong :(",
		}, nil
	}

	cc := lc.ClientContext

	// Get the deploys

	client := cc.Client

	fmt.Println(client)

	// var authInfo = nil
	// var deploys = client.ListSiteDeploys(authInfo)
	// fmt.Println("Deploys:", deploys)

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
