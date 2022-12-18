package client

import (
	"context"
	"time"
)

type ControllableClient struct {
	config     *ControllableClientConfig
	httpClient *ControllableHTTPClient
}

type ControllableClientConfig struct {
	serverEndpoint string
	environment    string
	appKey         string
	clientTimeout  time.Duration
}

func NewControllableClient(config *ControllableClientConfig) *ControllableClient {
	return &ControllableClient{
		config:     config,
		httpClient: NewControllableHTTPClient(config.serverEndpoint, config.appKey, config.clientTimeout),
	}
}

func (c *ControllableClient) CreatePropertyValue(ctx context.Context, propertyReferenceValuePairs *PropertyReferenceValuePairs) (*ExecutionResponse, error) {
	executionRequest := &ExecutionRequest{
		Operation:   OperationCreatePropertyValue,
		Environment: c.config.environment,
	}

	requests := make([]*PropertyExecutionRequest, len(propertyReferenceValuePairs.Pairs))

	for i, pair := range propertyReferenceValuePairs.Pairs {
		requests[i] = &PropertyExecutionRequest{
			Property: pair.Reference,
			Value:    pair.Value,
		}
	}

	executionRequest.Requests = requests
	return c.httpClient.Execute(ctx, executionRequest)
}

func (c *ControllableClient) UpdatePropertyValue(ctx context.Context, propertyReferenceValuePairs *PropertyReferenceValuePairs) (*ExecutionResponse, error) {
	executionRequest := &ExecutionRequest{
		Operation:   OperationUpdatePropertyValue,
		Environment: c.config.environment,
	}

	requests := make([]*PropertyExecutionRequest, len(propertyReferenceValuePairs.Pairs))

	for i, pair := range propertyReferenceValuePairs.Pairs {
		requests[i] = &PropertyExecutionRequest{
			Property: pair.Reference,
			Value:    pair.Value,
		}
	}

	executionRequest.Requests = requests
	return c.httpClient.Execute(ctx, executionRequest)
}

func (c *ControllableClient) DeletePropertyValue(ctx context.Context, propertyReferenceValuePairs *PropertyReferenceValuePairs) (*ExecutionResponse, error) {
	return nil, nil
}

func (c *ControllableClient) ReadPropertyValue(ctx context.Context, readPropertyRequests *ReadPropertyRequests) (*ExecutionResponse, error) {
	return nil, nil
}
