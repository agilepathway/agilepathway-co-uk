package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/go-openapi/runtime"
	openapiClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing"
	"github.com/netlify/open-api/go/plumbing/operations"
)

const (
	netlifyAPIHost string = "api.netlify.com"
	netlifyAPIPath string = "/api/v1"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	netlifyAccessToken := (strings.Fields(request.Headers["authorization"]))[1]
	commit := request.QueryStringParameters["commit"]
	siteID := request.QueryStringParameters["siteid"]
	fmt.Println("Finding deploy preview URL for commit:", commit)

	rawDeploys, error := netlifyClient().Operations.ListSiteDeploys(listSiteDeploysParams(siteID), authInfo(netlifyAccessToken))
	deploys := rawDeploys.Payload

	deployID, error := deployIDForCommit(commit, deploys)

	if error != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       error.Error(),
		}, nil
	}

	deployPreviewURL := fmt.Sprintf("https://%s--agilepathway-co-uk.netlify.com", deployID)
	fmt.Printf("Deploy preview url for commit %s: %s", commit, deployPreviewURL)
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       deployPreviewURL,
	}, nil
}

func deployIDForCommit(commit string, deploys []*models.Deploy) (string, error) {
	for _, deploy := range deploys {
		if deploy.CommitRef == commit {
			return deploy.ID, nil
		}
	}
	return "", fmt.Errorf("No Netlify deployment found for commit: %s", commit)
}

func listSiteDeploysParams(siteID string) *operations.ListSiteDeploysParams {
	return operations.NewListSiteDeploysParams().WithSiteID(siteID)
}

func netlifyClient() *plumbing.Netlify {
	transport := openapiClient.NewWithClient(netlifyAPIHost, netlifyAPIPath, plumbing.DefaultSchemes, httpClient())
	client := plumbing.New(transport, strfmt.Default)
	return client
}

func authInfo(netlifyAccessToken string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("User-Agent", "User-Agent: NetlifyDeployPreviewFunction/0.0.0")
		r.SetHeaderParam("Authorization", "Bearer "+netlifyAccessToken)
		return nil
	})
}

func httpClient() *http.Client {
	return &http.Client{}
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
