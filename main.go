package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/netlify/open-api/go/plumbing"
	"github.com/netlify/open-api/go/plumbing/operations"
	
	"github.com/go-openapi/runtime"
	openapiClient "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"

)

// Netlify specific constants
const (
	NetlifyAPIHost string = "api.netlify.com"

	// NetlifyAPIPath is path attached to baseURL for making Netlify API request
	NetlifyAPIPath string = "/api/v1"

)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Finding deploy preview URL for commit:", request.QueryStringParameters["commit"])

	var list_site_deploys_token = os.Getenv("LIST_SITE_DEPLOYS_TOKEN")

	authInfo := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("User-Agent", "agilepathway")
		r.SetHeaderParam("Authorization", "Bearer "+list_site_deploys_token)
		return nil
	})

	fmt.Println("authInfo:", authInfo)

	var client = getNetlifyClient()
	fmt.Println("Client:", client)


	// var deploys = client.ListSiteDeploys(authInfo)
	// fmt.Println("Deploys:", deploys)

	const deploy_preview_url = "https://netlify-function--agilepathway-co-uk.netlify.com"
	fmt.Println("Deploy preview url found:", deploy_preview_url)
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       deploy_preview_url,
	}, nil
}

func getNetlifyClient() (*plumbing.Netlify) {
	// Create OpenAPI transport
	transport := openapiClient.NewWithClient(NetlifyAPIHost, NetlifyAPIPath, plumbing.DefaultSchemes, p.getHTTPClient())
	transport.SetDebug(true)

	// Create Netlify client by adding the transport to it
	client := plumbing.New(transport, strfmt.Default)

	return client
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
