package client

import (
	"context"
	"time"
)

type ControllableClient struct {
	Config     *ControllableClientConfig
	HTTPClient *ControllableHTTPClient
}

type ControllableClientConfig struct {
	ServerEndpoint string
	Environment    string
	ClientTimeout  time.Duration
}

func NewControllableClient(config *ControllableClientConfig) *ControllableClient {
	return &ControllableClient{
		Config:     config,
		HTTPClient: NewControllableHTTPClient(config.ServerEndpoint, config.ClientTimeout),
	}
}

func (c *ControllableClient) CreatePropertyValue(ctx context.Context, propertyReferenceValuePairs *PropertyReferenceValuePairs) (*ExecutionResponse, error) {
	return nil, nil
}

func (c *ControllableClient) UpdatePropertyValue(ctx context.Context, propertyReferenceValuePairs *PropertyReferenceValuePairs) (*ExecutionResponse, error) {
	return nil, nil
}

func (c *ControllableClient) DeletePropertyValue(ctx context.Context, propertyReferenceValuePairs *PropertyReferenceValuePairs) (*ExecutionResponse, error) {
	return nil, nil
}

func (c *ControllableClient) ReadPropertyValue(ctx context.Context, readPropertyRequests *ReadPropertyRequests) (*ExecutionResponse, error) {
	return nil, nil
}
