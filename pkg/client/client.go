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
	ServerEndpoint string
	Environment    string
	AppKey         string
	ClientTimeout  time.Duration
}

func NewControllableClient(config *ControllableClientConfig) *ControllableClient {
	return &ControllableClient{
		config:     config,
		httpClient: NewControllableHTTPClient(config.ServerEndpoint, config.AppKey, config.ClientTimeout),
	}
}

func (c *ControllableClient) CreatePropertyValue(ctx context.Context, propertyReferenceValuePairs *PropertyReferenceValuePairs) (*ExecutionResponse, error) {
	executionRequest := &ExecutionRequest{
		Operation:   OperationCreatePropertyValue,
		Environment: c.config.Environment,
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
		Environment: c.config.Environment,
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
	executionRequest := &ExecutionRequest{
		Operation:   OperationDeletePropertyValue,
		Environment: c.config.Environment,
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

func (c *ControllableClient) ReadPropertyValue(ctx context.Context, readPropertyRequests *ReadPropertyRequests) (*ExecutionResponse, error) {
	executionRequest := &ExecutionRequest{
		Operation:   OperationReadPropertyValue,
		Environment: c.config.Environment,
	}

	requests := make([]*PropertyExecutionRequest, len(readPropertyRequests.Requests))

	for i, readPropertyRequest := range readPropertyRequests.Requests {
		requests[i] = &PropertyExecutionRequest{
			Property: readPropertyRequest.Reference,
			Params:   readPropertyRequest.Params,
		}
	}

	executionRequest.Requests = requests
	return c.httpClient.Execute(ctx, executionRequest)
}
