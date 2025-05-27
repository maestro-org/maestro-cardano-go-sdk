package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/maestro-org/go-sdk/config"
)

type Client struct {
	apiKey     string
	network    string
	version    string
	HTTPClient *http.Client
	BaseUrl    string
}

func NewClient(apiKey string, network string) *Client {
	cfg := config.GetConfig()
	return &Client{
		apiKey:  apiKey,
		network: network,
		version: cfg.Client.Version,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		BaseUrl: fmt.Sprintf("https://%s.gomaestro-api.org/%s", network, cfg.Client.Version),
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// type successResponse struct {
// 	Code int         `json:"code"`
// 	Data interface{} `json:"data"`
// }

func (c *Client) sendRequest(req *http.Request, responseBody *string) error {
	if req == nil {
		return fmt.Errorf("empty request")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Add("api-key", c.apiKey)

	if c.HTTPClient == nil {
		return fmt.Errorf("missing http client")
	}
	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}
	if resp == nil {
		return fmt.Errorf("empty response")
	}

	// Try to unmarshall into errorResponse
	if resp.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", resp.StatusCode)
	}

	respBodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("failed to read body: %s", err)
	}
	defer resp.Body.Close() //nolint:errcheck

	*responseBody = string(respBodyBytes)

	return nil
}

func (c *Client) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseUrl+url, nil)
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Add("api-key", c.apiKey)
	return c.HTTPClient.Do(req)
}

func (c *Client) post(url string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.BaseUrl+url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Add("api-key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	return c.HTTPClient.Do(req)
}

func (c *Client) postBuffer(url string, buffer []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.BaseUrl+url, bytes.NewBuffer(buffer))
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, fmt.Errorf("empty request")
	}
	req.Header.Set("Accept", "application/cbor")
	req.Header.Add("api-key", c.apiKey)
	req.Header.Set("Content-Type", "application/cbor")
	return c.HTTPClient.Do(req)
}
