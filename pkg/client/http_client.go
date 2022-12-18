package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ControllableHTTPClient struct {
	Endpoint          string
	HttpClientTimeout time.Duration
	Client            *http.Client
}

func NewControllableHTTPClient(endpoint string, httpClientTimeout time.Duration) *ControllableHTTPClient {
	return &ControllableHTTPClient{
		Endpoint:          endpoint,
		HttpClientTimeout: httpClientTimeout,
		Client: &http.Client{
			Timeout: httpClientTimeout,
		},
	}
}

func (c *ControllableHTTPClient) Execute(ctx context.Context, executionRequest *ExecutionRequest) (*ExecutionResponse, error) {
	requestBody, err := json.Marshal(executionRequest)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/properties/execute", c.Endpoint), bytes.NewReader(requestBody))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		responseBody, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		var errorResponse ErrorResponse

		err = json.Unmarshal(responseBody, &errorResponse)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorResponse.Message)
	}

	responseBody, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var executionResponse ExecutionResponse

	err = json.Unmarshal(responseBody, &executionResponse)

	if err != nil {
		return nil, err
	}

	return &executionResponse, nil
}
