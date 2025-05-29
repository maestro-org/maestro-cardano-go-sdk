package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) DatumFromHash(hash string) (*models.DatumFromHash, error) {
	url := fmt.Sprintf("/data/%s", hash)
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
	var datum models.DatumFromHash
	err = json.NewDecoder(resp.Body).Decode(&datum)
	if err != nil {
		return nil, err
	}
	return &datum, nil
}
