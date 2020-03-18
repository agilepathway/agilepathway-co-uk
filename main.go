package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

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
	// TODO: do I need these - I think they might be the defaults anyway?
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

	var netlify_client = getNetlifyClient()
	var list_site_deploys_params = getListSiteDeploysParams()

	var deploys, error = netlify_client.Operations.ListSiteDeploys(list_site_deploys_params, authInfo)
	fmt.Println("Deploys:", deploys)
	fmt.Println("Error:", error)

	const deploy_preview_url = "https://netlify-function--agilepathway-co-uk.netlify.com"
	fmt.Println("Deploy preview url found:", deploy_preview_url)
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       deploy_preview_url,
	}, nil
}

type ListSiteDeploysParams struct {
    site_id string
}

func getListSiteDeploysParams() (*ListSiteDeploysParams) {
	var site_id = os.Getenv("SITE_ID")
	fmt.Println("site id:", site_id)
	params := ListSiteDeploysParams{site_id: site_id}
	return params
}

func getNetlifyClient() (*plumbing.Netlify) {
	// Create OpenAPI transport
	transport := openapiClient.NewWithClient(NetlifyAPIHost, NetlifyAPIPath, plumbing.DefaultSchemes, getHTTPClient())
	transport.SetDebug(true)

	// Create Netlify client by adding the transport to it
	client := plumbing.New(transport, strfmt.Default)

	return client
}

func getHTTPClient() *http.Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   -1,
			DisableKeepAlives:     true}}

	return httpClient
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
