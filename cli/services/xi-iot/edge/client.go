package edge

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"time"
	edge_client "xi-iot-cli/generated/edge_swagger/client"
	"xi-iot-cli/generated/edge_swagger/client/operations"

	"xi-iot-cli/generated/edge_swagger/models"

	"xi-iot-cli/xi-iot/cloudmgmt"
	"xi-iot-cli/xi-iot/errutils"

	httptransport "github.com/go-openapi/runtime/client"
)

// Client ...
type Client struct {
	client    *edge_client.Sherlock
	cmeClient *cloudmgmt.Client
}

// WithCMEClient modify client.cmeClient to point to CMEClient
func WithCMEClient(CMEClient *cloudmgmt.Client) func(*Client) {
	return func(client *Client) {
		client.cmeClient = CMEClient
	}
}

// New creates a new Client
func New(edgeURL string, options ...func(*Client)) *Client {
	clientConfig := &edge_client.TransportConfig{
		Host:     edgeURL,
		BasePath: "/v1",
		Schemes:  []string{"https"},
	}
	edgeClient := edge_client.NewHTTPClientWithConfig(nil, clientConfig)
	c := &Client{client: edgeClient}
	for _, option := range options {
		option(c)
	}
	certificate, err := c.cmeClient.CreateCertificates()
	if err != nil {
		errutils.Exitf("Failed to get HTTPS certificate from cloud server. %s", *err.Message)
	}
	clientCert := []byte(*certificate.Certificate)
	clientKey := []byte(*certificate.PrivateKey)
	caCert := []byte(*certificate.CACertificate)
	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(caCert)
	if !ok {
		errutils.Exitf("Failed to apped CA certificate")
	}
	cert, keyErr := tls.X509KeyPair(clientCert, clientKey)
	if keyErr != nil {
		errutils.Exitf(keyErr.Error())
	}

	transport := httptransport.New(clientConfig.Host, clientConfig.BasePath, clientConfig.Schemes)
	transport.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      caCertPool,
			ClientCAs:    caCertPool,
			Certificates: []tls.Certificate{cert},
			ClientAuth:   tls.RequireAndVerifyClientCert,
			ServerName:   "localhost",
		},
	}
	c.client.SetTransport(transport)

	return c
}

// GetInfo get info
func (c *Client) GetInfo() (*models.Info, error) {
	listParam := operations.NewInfoParams()
	ok, err := c.client.Operations.Info(listParam)
	if err != nil {
		return nil, err
	}
	return ok.Payload, nil
}

// GetDataPipeline gets datapipeline from edge given datapipeline id
func (c *Client) GetDataPipeline(ID string) (*models.Stream, error) {
	listParam := operations.NewStreamGetParams()
	listParam.ID = ID
	ok, err := c.client.Operations.StreamGet(listParam)
	if err != nil {
		return nil, err
	}
	return ok.Payload, nil
}

// DebugSend sends fake data msg to trigger datapipeline
func (c *Client) DebugSend(ID string, msg string) error {
	debugParam := operations.NewDebugSendPostParams()
	debugParam.ID = ID
	debugParam.Body = msg
	_, err := c.client.Operations.DebugSendPost(debugParam)
	if err != nil {
		return err
	}
	return nil
}

// DebugReceive receives data pipeline output within duration time
func (c *Client) DebugReceive(ID string, duration string) (*models.DebugReceiveResponse, error) {
	durationDecode, err := time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}
	if durationDecode < 5*time.Second || durationDecode > 1*time.Hour {
		return nil, fmt.Errorf("Duration should be within 5s - 1h")
	}
	// Set the http timeout to be duration plus 30 seconds.
	debugParam := operations.NewDebugReceiveGetParamsWithTimeout(durationDecode + 30*time.Second)
	debugParam.ID = ID
	debugParam.Duration = duration
	ok, err := c.client.Operations.DebugReceiveGet(debugParam)
	if err != nil {
		return nil, err
	}
	return ok.Payload, nil
}
