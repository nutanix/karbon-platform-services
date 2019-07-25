package devtools

import (
	"crypto/tls"
	"net/http"
	devtools_client "xi-iot-cli/generated/devtools_swagger/client"
	"xi-iot-cli/generated/devtools_swagger/client/operations"
	"xi-iot-cli/generated/devtools_swagger/models"

	"xi-iot-cli/xi-iot/cloudmgmt"

	httptransport "github.com/go-openapi/runtime/client"
)

// Client ...
type Client struct {
	client    *devtools_client.Sherlock
	cmeClient *cloudmgmt.Client
}

// WithCMEClient modify client.cmeClient to point to CMEClient
func WithCMEClient(CMEClient *cloudmgmt.Client) func(*Client) {
	return func(client *Client) {
		client.cmeClient = CMEClient
	}
}

// New creates a new Client
func New(devtoolsURL string, options ...func(*Client)) *Client {
	clientConfig := &devtools_client.TransportConfig{
		Host:     devtoolsURL,
		BasePath: "/v1.0",
		Schemes:  []string{"https"},
	}
	devtoolsClient := devtools_client.NewHTTPClientWithConfig(nil, clientConfig)
	transport := httptransport.New(clientConfig.Host, clientConfig.BasePath, clientConfig.Schemes)
	transport.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	devtoolsClient.SetTransport(transport)
	c := &Client{client: devtoolsClient}
	for _, option := range options {
		option(c)
	}
	return c
}

func (c *Client) FetchLogs(endpoint string, time string) (*models.GetLogsResponse, error) {
	params := operations.NewGetStreamLogsParams()
	params.Endpoint = endpoint
	params.Latestts = time
	ok, err := c.client.Operations.GetStreamLogs(params)
	if err != nil {
		return nil, err
	}
	return ok.Payload, nil
}
