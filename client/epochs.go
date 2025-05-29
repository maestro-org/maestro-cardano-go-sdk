package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) CurrentEpoch() (*models.EpochResp, error) {
	url := "/epochs/current"
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
	var currentEpoch models.EpochResp
	err = json.NewDecoder(resp.Body).Decode(&currentEpoch)
	if err != nil {
		return nil, err
	}
	return &currentEpoch, nil
}

func (c *Client) SpecificEpoch(epochNo int) (*models.EpochResp, error) {
	url := fmt.Sprintf("/epochs/%d/info", epochNo)
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
	var specificEpoch models.EpochResp
	err = json.NewDecoder(resp.Body).Decode(&specificEpoch)
	if err != nil {
		return nil, err
	}
	return &specificEpoch, nil
}
