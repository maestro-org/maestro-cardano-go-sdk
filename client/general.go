package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) ChainTip() (*models.ChainTip, error) {
	url := "/chain-tip"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var chainTip models.ChainTip
	err = json.NewDecoder(resp.Body).Decode(&chainTip)
	if err != nil {
		return nil, err
	}
	return &chainTip, nil
}

func (c *Client) EraSummaries() (*models.EraSummaries, error) {
	url := "/era-summaries"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var eraSummaries models.EraSummaries
	err = json.NewDecoder(resp.Body).Decode(&eraSummaries)
	if err != nil {
		return nil, err
	}
	return &eraSummaries, nil
}

func (c *Client) ProtocolParameters() (*models.ProtocolParameters, error) {
	url := "/protocol-parameters"
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var protocolParameters models.ProtocolParameters
	err = json.NewDecoder(resp.Body).Decode(&protocolParameters)
	if err != nil {
		return nil, err
	}
	return &protocolParameters, nil
}

func (c *Client) BlockChainStartTime() (models.BasicResponse, error) {
	url := "/system-start"
	resp, err := c.get(url)
	if err != nil {
		return models.BasicResponse{}, err
	}
	if resp == nil {
		return models.BasicResponse{}, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return models.BasicResponse{}, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var blockchainStartTime models.BasicResponse
	err = json.NewDecoder(resp.Body).Decode(&blockchainStartTime)
	if err != nil {
		return models.BasicResponse{}, err
	}
	return blockchainStartTime, nil
}
