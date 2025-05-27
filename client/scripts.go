package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) ScriptByHash(hash string) (*models.ScriptByHash, error) {
	url := fmt.Sprintf("/scripts/%s", hash)
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
	var scriptByHash models.ScriptByHash
	err = json.NewDecoder(resp.Body).Decode(&scriptByHash)
	if err != nil {
		return nil, err
	}
	return &scriptByHash, nil

}
