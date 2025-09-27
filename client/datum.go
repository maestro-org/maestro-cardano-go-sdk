package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maestro-org/go-sdk/models"
)

func (c *Client) DatumFromHash(hash string) (*models.DatumFromHash, error) {
	url := fmt.Sprintf("/datums/%s", hash)
	resp, err := c.get(url)
	if err != nil {
		fmt.Println("Error getting datum from hash:", err)
		return nil, err
	}
	if resp == nil {
		fmt.Println("Empty response")
		return nil, fmt.Errorf("empty response")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected error: %d", resp.Body)
	}
	defer resp.Body.Close() //nolint:errcheck
	var datum models.DatumFromHash
	err = json.NewDecoder(resp.Body).Decode(&datum)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil, err
	}
	return &datum, nil
}
